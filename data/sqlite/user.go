package store

import "os"
import "mudkip/lib"

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
	rwl.Lock()
	defer rwl.Unlock()

	var exists bool
	if exists, err = this.userExists(usr.Id); err != nil {
		return
	}

	if exists {
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
	} else {
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

func (this *Store) userExists(id int64) (bool, os.Error) {
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
