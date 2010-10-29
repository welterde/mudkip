package store

import "os"
import "mudkip/lib"

func (this *Store) GetUser(id int64) (usr *lib.UserInfo, err os.Error) {
	if this.qry, err = this.conn.Prepare("select * from users where id=?"); err != nil {
		return nil, err
	}

	this.qry.Exec(id)

	if !this.qry.Next() {
		return nil, lib.ErrUnknownUser
	}

	usr = new(lib.UserInfo)
	if err = this.qry.Scan(&usr.Id, &usr.Name, &usr.Password, &usr.Registered, &usr.Zone, &usr.Character); err != nil {
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
	if err = this.qry.Scan(&usr.Id, &usr.Name, &usr.Password, &usr.Registered, &usr.Zone, &usr.Character); err != nil {
		return
	}

	return usr, this.qry.Finalize()
}

func (this *Store) SetUser(usr *lib.UserInfo) (err os.Error) {
	rwl.Lock()
	defer rwl.Unlock()

	var exists bool
	if exists, err = this.itemExists(usr.Id, "users"); err != nil {
		return
	}

	if exists {
		if this.qry, err = this.conn.Prepare(
			`update users set name=?, password=?, registered=?, zone=?, characterid=? where id=?`,
		); err != nil {
			return err
		}

		if err = this.qry.Exec(usr.Name, usr.Password, usr.Registered, usr.Zone, usr.Id, usr.Character); err != nil {
			return
		}

		this.qry.Next()
		this.qry.Finalize()
	} else {
		if this.qry, err = this.conn.Prepare(
			`insert into users (name, password, registered, zone, characterid) values(?,?,?,?,?)`,
		); err != nil {
			return
		}

		if err = this.qry.Exec(usr.Name, usr.Password, usr.Registered, usr.Zone, usr.Character); err != nil {
			return
		}

		this.qry.Next()
		this.qry.Finalize()

		if usr.Id, err = this.conn.LastInsertId(); err != nil {
			return
		}

		if usr.Id == 0 {
			return os.NewError("Insert of user failed")
		}
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
