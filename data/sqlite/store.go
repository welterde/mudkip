package store

import "os"
import "mudkip/lib"

type Store struct {
	conn *Conn
	qry  *Stmt
}

func New() lib.DataStore {
	return new(Store)
}

func (this *Store) Open(params map[string]string) (err os.Error) {
	if this.conn != nil {
		return
	}

	this.conn, err = sqlite_Open(params["file"])
	return
}

func (this *Store) Close() {
	this.qry = nil

	if this.conn != nil {
		this.conn.Close()
		this.conn = nil
	}
}

func (this *Store) Initialize() (err os.Error) {
	if this.initialized() {
		return
	}

	if err = this.conn.Exec(`
		create table objects (
			id		INTEGER PRIMARY KEY AUTOINCREMENT,
			type	TINYINT,
			data	BLOB NOT NULL
		);`); err != nil {
		return
	}

	if err = this.conn.Exec(`
		create table users (
			id			INTEGER PRIMARY KEY AUTOINCREMENT,
			name		VARCHAR(120) NOT NULL,
			password	VARCHAR(41) NOT NULL,
			registered	INTEGER NOT NULL,
			zone        INTEGER NOT NULL
		);`); err != nil {
		return
	}

	return
}

func (this *Store) GetUser(id uint16) (usr *lib.UserInfo, err os.Error) {
	if this.qry, err = this.conn.Prepare("select * from users where id=?"); err != nil {
		return nil, err
	}

	this.qry.Exec(id)

	if !this.qry.Next() {
		return nil, lib.ErrUnknownUser
	}

	usr = new(lib.UserInfo)
	if err = this.qry.Scan(&usr.Id, &usr.Name, &usr.Password, &usr.Registered, &usr.Zone); err != nil {
		return
	}

	return usr, this.qry.Finalize()
}

func (this *Store) GetUserByName(name string) (usr *lib.UserInfo, err os.Error) {
	if this.qry, err = this.conn.Prepare("select * from users where name like ?"); err != nil {
		return nil, err
	}

	this.qry.Exec(name)

	if !this.qry.Next() {
		return nil, lib.ErrUnknownUser
	}

	usr = new(lib.UserInfo)
	if err = this.qry.Scan(&usr.Id, &usr.Name, &usr.Password, &usr.Registered, &usr.Zone); err != nil {
		return
	}

	return usr, this.qry.Finalize()
}

func (this *Store) SetUser(usr *lib.UserInfo) (err os.Error) {
	var exists bool
	if exists, err = this.userExists(usr.Id); err != nil {
		return
	}

	if exists { // update
		if this.qry, err = this.conn.Prepare(
			`update users set name=?, password=?, registered=?, zone=? where id=?`,
		); err != nil {
			return err
		}

		if err = this.qry.Exec(usr.Name, usr.Password, usr.Registered, usr.Zone, usr.Id); err != nil {
			return
		}

		this.qry.Next()
		this.qry.Finalize()
	} else { // insert
		if this.qry, err = this.conn.Prepare(
			`insert into users (name, password, registered, zone) values(?, ?, ?, ?)`,
		); err != nil {
			return
		}

		if err = this.qry.Exec(usr.Name, usr.Password, usr.Registered, usr.Zone); err != nil {
			return
		}

		this.qry.Next()
		this.qry.Finalize()

		var id int64
		if id, err = this.conn.LastInsertId(); err != nil {
			return
		}

		if id == 0 {
			return os.NewError("Insert of object failed")
		}

		usr.Id = uint16(id)
	}

	return
}

func (this *Store) GetUsers() (list []*lib.UserInfo, err os.Error) {
	var sz int
	var usr *lib.UserInfo

	list = make([]*lib.UserInfo, 0, 16)

	if this.qry, err = this.conn.Prepare("select * from users order by name desc"); err != nil {
		return
	}

	this.qry.Exec()

	for this.qry.Next() {
		usr = new(lib.UserInfo)
		if err = this.qry.Scan(&usr.Id, &usr.Name, &usr.Password, &usr.Registered, &usr.Zone); err != nil {
			return
		}

		if sz >= cap(list) {
			cp := make([]*lib.UserInfo, sz, sz+16)
			copy(cp, list)
			list = cp
		}

		list = list[0 : sz+1]
		list[sz] = usr
		sz++
	}

	return list, this.qry.Finalize()
}

func (this *Store) GetObject(id uint16, objtype uint8) (lib.Object, os.Error) {
	var blob []byte
	var dbtype uint8
	var err os.Error

	if this.qry, err = this.conn.Prepare("select type, data from objects where id=?"); err != nil {
		return nil, err
	}

	this.qry.Exec(id)

	if !this.qry.Next() {
		return nil, lib.ErrUnknownObject
	}

	this.qry.Scan(&dbtype, &blob)
	this.qry.Finalize()

	if dbtype != objtype {
		return nil, lib.ErrTypeMismatch
	}

	return lib.Deserialize(id, objtype, blob)
}

func (this *Store) GetObjectsByType(objtype uint8) ([]lib.Object, os.Error) {
	var err os.Error
	var sz int
	var data []byte
	var id int64

	list := make([]lib.Object, 0, 16)

	if this.qry, err = this.conn.Prepare("select id, data from objects where type=?"); err != nil {
		return nil, err
	}

	this.qry.Exec(objtype)

	for this.qry.Next() {
		if err = this.qry.Scan(&id, &data); err != nil {
			return nil, err
		}

		if sz >= cap(list) {
			cp := make([]lib.Object, sz, sz+16)
			copy(cp, list)
			list = cp
		}

		list = list[0 : sz+1]

		if list[sz], err = lib.Deserialize(uint16(id), objtype, data); err != nil {
			return nil, err
		}

		sz++
	}

	return list, this.qry.Finalize()
}

func (this *Store) SetObject(obj lib.Object) (err os.Error) {
	var blob []byte
	var exists bool

	if blob, err = lib.Serialize(obj); err != nil {
		return
	}

	if exists, err = this.objectExists(obj.Id()); err != nil {
		return
	}

	if exists { // update
		if this.qry, err = this.conn.Prepare(
			`update objects set type=?, data=? where id=?`,
		); err != nil {
			return err
		}

		if err = this.qry.Exec(obj.Type(), blob, obj.Id()); err != nil {
			return
		}

		this.qry.Next()
		this.qry.Finalize()
	} else { // insert
		if this.qry, err = this.conn.Prepare(
			`insert into objects (type, data) values(?, ?)`,
		); err != nil {
			return
		}

		if err = this.qry.Exec(obj.Type(), blob); err != nil {
			return
		}

		this.qry.Next()
		this.qry.Finalize()

		var id int64
		if id, err = this.conn.LastInsertId(); err != nil {
			return
		}

		if id == 0 {
			return os.NewError("Insert of object failed")
		}

		obj.SetId(uint16(id))
	}

	return
}

func (this *Store) objectExists(id uint16) (bool, os.Error) {
	var err os.Error

	if this.qry, err = this.conn.Prepare("select count(*) from objects where id=?"); err != nil {
		return false, err
	}

	if err = this.qry.Exec(id); err != nil {
		return false, err
	}

	var count int

	this.qry.Next()
	this.qry.Scan(&count)
	this.qry.Finalize()

	return count == 1, nil
}

func (this *Store) userExists(id uint16) (bool, os.Error) {
	var err os.Error

	if this.qry, err = this.conn.Prepare("select count(*) from users where id=?"); err != nil {
		return false, err
	}

	if err = this.qry.Exec(id); err != nil {
		return false, err
	}

	var count int

	this.qry.Next()
	this.qry.Scan(&count)
	this.qry.Finalize()

	return count == 1, nil
}


func (this *Store) initialized() bool {
	var err os.Error
	var name string
	var ok bool

	if this.qry, err = this.conn.Prepare("select tbl_name from sqlite_master;"); err != nil {
		return false
	}

	defer this.qry.Finalize()

	required := make(map[string]bool)
	required["objects"] = false
	required["users"] = false

	for this.qry.Next() {
		if err = this.qry.Scan(&name); err != nil {
			continue
		}

		if _, ok = required[name]; ok {
			required[name] = true
		}
	}

	for _, ok = range required {
		if !ok {
			return false
		}
	}

	return true
}
