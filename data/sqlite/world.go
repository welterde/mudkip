package store

import "os"
import "time"
import "mudkip/lib"

func (this *Store) GetWorld() (world *lib.World, err os.Error) {
	if this.qry, err = this.conn.Prepare("select * from worlds"); err != nil {
		return nil, err
	}

	this.qry.Exec()

	if !this.qry.Next() {
		return nil, lib.ErrNoWorldInfo
	}

	world = lib.NewWorld()
	if err = this.qry.Scan(
		&world.Id, &world.Created, &world.Name, &world.Description, &world.Logo,
		&world.Motd, &world.DefaultZone, &world.AllowRegister, &world.LevelCap,
	); err != nil {
		return
	}

	return world, this.qry.Finalize()
}

func (this *Store) SetWorld(world *lib.World) (err os.Error) {
	rwl.Lock()
	defer rwl.Unlock()

	var exists bool
	if exists, err = this.itemExists(world.Id, "worlds"); err != nil {
		return
	}

	if exists {
		if this.qry, err = this.conn.Prepare(
			`update worlds set name=?, description=?, logo=?, motd=?, defaultzone=?, allowregister=?, levelcap=? where id=?`,
		); err != nil {
			return err
		}

		if err = this.qry.Exec(
			world.Name, world.Description, world.Logo, world.Motd,
			world.DefaultZone, world.AllowRegister, world.LevelCap, world.Id,
		); err != nil {
			return
		}

		this.qry.Next()
		this.qry.Finalize()
	} else {
		if this.qry, err = this.conn.Prepare(
			`insert into worlds (created, name, description, logo, motd, defaultzone, allowregister, levelcap) values(?,?,?,?,?,?,?,?)`,
		); err != nil {
			return
		}

		world.Created = time.Seconds()

		if err = this.qry.Exec(
			world.Created, world.Name, world.Description, world.Logo,
			world.Motd, world.DefaultZone, world.AllowRegister, world.LevelCap,
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
