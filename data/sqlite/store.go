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

func (this *Store) Initialized() bool {
	var err os.Error
	var name string

	if this.qry, err = this.conn.Prepare("select tbl_name from sqlite_master;"); err != nil {
		return false
	}

	defer this.qry.Finalize()

	for this.qry.Next() {
		if err = this.qry.Scan(&name); err != nil {
			return false
		}

		switch name {
		case "objects":
			return true
		}
	}

	return false
}

func (this *Store) Initialize() (err os.Error) {
	// We store objects in a compressed bit stream format. Only the ID and the
	// object Type are independant columns, so we can select on them.
	// Rowid column is implicit.

	if err = this.conn.Exec(`
	create table objects (
		type	TINYINT,
		data	BLOB NOT NULL
	);`); err != nil {
		return
	}

	return
}

func (this *Store) GetObject(id uint16, objtype uint8) (lib.Object, os.Error) {
	var blob []byte
	var dbtype uint8
	var err os.Error

	if this.qry, err = this.conn.Prepare("select type, data from objects where rowid=?"); err != nil {
		return nil, err
	}

	this.qry.Exec(id)

	if !this.qry.Next() {
		return nil, lib.ErrUnknownObject
	}

	if err = this.qry.Scan(&dbtype, &blob); err != nil {
		return nil, err
	}

	if err = this.qry.Finalize(); err != nil {
		return nil, err
	}

	if dbtype != objtype {
		return nil, lib.ErrTypeMismatch
	}

	// Use the builtin lib.Deserialize function to actually create the
	// object with it's fields values. This way we do not have to care about 
	// what fields are actually in a given type.
	return lib.Deserialize(id, objtype, blob)
}

func (this *Store) SetObject(obj lib.Object) (err os.Error) {
	var blob []byte
	var exists bool

	// Use the builtin lib.Serialize function to actually create the bit stream.
	// This way we do not have to care about what fields are actually in a given type.
	if blob, err = lib.Serialize(obj); err != nil {
		return
	}

	if exists, err = this.ObjectExists(obj.Id()); err != nil {
		return
	}

	if exists { // update
		if this.qry, err = this.conn.Prepare(
			`update objects set type=?, data=? where rowid=?`,
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

		if err = this.qry.Finalize(); err != nil {
			return
		}

		var id int64
		if id, err = this.conn.LastInsertId(); err != nil {
			return
		}

		if id == 0 {
			return os.NewError("Insert of object failed")
		}

		// Make sure we set the object's ID to the rowid in the database.
		obj.SetId(uint16(id))
	}

	return
}

func (this *Store) ObjectExists(id uint16) (bool, os.Error) {
	var err os.Error

	if this.qry, err = this.conn.Prepare("select count(*) from objects where rowid=?"); err != nil {
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

func (this *Store) LastInsertId() (int64, os.Error) {
	var err os.Error
	var id int64

	if this.qry, err = this.conn.Prepare(`select last_insert_rowid()`); err != nil {
		return 0, err
	}

	defer this.qry.Finalize()

	if err = this.qry.Exec(); err != nil {
		return 0, err
	}

	this.qry.Next()
	this.qry.Scan(&id)
	return id, nil
}
