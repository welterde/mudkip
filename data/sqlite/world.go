package store

import "os"
import "time"
import "mudkip/lib"

func (this *Store) Initialize(world *lib.World) (err os.Error) {
	rwl.Lock()
	defer rwl.Unlock()

	if err = this.conn.ExecRange(`
		begin transaction;

		drop table if exists world;
		create table world (
			id            integer primary key autoincrement,
			created       integer NOT NULL,
			name          text not null,
			description   text not null,
			logo          text not null,
			motd          text,
			defaultzone   integer not null,
			allowregister boolean not null,
			levelcap      integer not null
		);

		drop table if exists users;
		create table users (
			id          integer primary key autoincrement,
			name        varchar(120) not null,
			password    varchar(50) not null,
			registered  integer not null,
			zone        integer not null,
			character   integer not null
		);

		end transaction;`,
	); err != nil {
		return
	}

	return this.SetWorld(world)
}

func (this *Store) GetWorld() (info *lib.World, err os.Error) {
	if this.qry, err = this.conn.Prepare("select * from world"); err != nil {
		return nil, err
	}

	this.qry.Exec()

	if !this.qry.Next() {
		return nil, lib.ErrNoWorldInfo
	}

	world := lib.NewWorld()
	if err = this.qry.Scan(
		&world.Id, &world.Created, &world.Name, &world.Description, &world.Logo,
		&world.Motd, &world.DefaultZone, &world.AllowRegister,
	); err != nil {
		return
	}

	return world, this.qry.Finalize()
}

func (this *Store) SetWorld(world *lib.World) (err os.Error) {
	rwl.Lock()
	defer rwl.Unlock()

	var exists bool
	if exists, err = this.worldExists(world.Id); err != nil {
		return
	}

	if exists {
		if this.qry, err = this.conn.Prepare(
			`update world set name=?, description=?, logo=?, motd=?, defaultzone=?, allowregister=? where id=?`,
		); err != nil {
			return err
		}

		if err = this.qry.Exec(
			world.Name, world.Description, world.Logo, world.Motd,
			world.DefaultZone, world.AllowRegister, world.Id,
		); err != nil {
			return
		}

		this.qry.Next()
		this.qry.Finalize()
	} else {
		if this.qry, err = this.conn.Prepare(
			`insert into world (created, name, description, logo, motd, defaultzone, allowregister) values(?, ?, ?, ?, ?, ?, ?)`,
		); err != nil {
			return
		}

		world.Created = time.Seconds()

		if err = this.qry.Exec(
			world.Created, world.Name, world.Description, world.Logo,
			world.Motd, world.DefaultZone, world.AllowRegister,
		); err != nil {
			return
		}

		this.qry.Next()
		this.qry.Finalize()

		if world.Id, err = this.conn.LastInsertId(); err != nil {
			return
		}

		if world.Id == 0 {
			return os.NewError("Insert of world failed")
		}
	}

	return
}

func (this *Store) worldExists(id int64) (bool, os.Error) {
	var err os.Error

	if this.qry, err = this.conn.Prepare("select count(*) from world where id=?"); err != nil {
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
