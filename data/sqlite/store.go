package store

import "os"
import "sync"
import "mudkip/lib"

type ItemType uint8

// Item types - used as enumerations in inventory slots to distinguish different
// item sources: armor, weapons, consumbales, etc.
const (
	Armor ItemType = iota
	Weapon
	Consumbale
)

// SQlite has some issues with write operations from multiple clients.
// We therefor use this read/write lock in the routines that perform write
// operations. Read operations should not be a problem. It is not part of the
// Store type, because we can have many instance of Store in a single server
// session.
var rwl = new(sync.RWMutex)

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

func (this *Store) Initialize(world *lib.World) (err os.Error) {
	rwl.Lock()

	if err = this.conn.ExecRange(`
		begin transaction;

		drop table if exists users;
		create table users (
			id          integer primary key autoincrement,
			name        varchar(120) not null,
			password    varchar(50) not null,
			registered  integer not null,
			zoneid      integer not null,
			characterid integer not null
		);

		drop table if exists armors;
		create table armors (
			id          integer primary key autoincrement,
			name        text not null,
			description text not null,
			type        tinyint not null,
			statsid     integer not null
		);

		drop table if exists characters;
		create table characters (
			id          integer primary key autoincrement,
			name        text not null,
			description text not null,
			title       text not null,
			level       integer not null,
			bankroll    integer not null,
			standing    tinyint not null,
			groupid     integer not null,
			classid     integer not null,
			raceid      integer not null,
			zoneid      integer not null,
			statsid     integer not null,
			inventoryid integer not null
		);

		drop table if exists classes;
		create table classes (
			id          integer primary key autoincrement,
			name        text not null,
			description text not null,
			statsid     integer not null
		);

		drop table if exists consumables;
		create table consumables (
			id          integer primary key autoincrement,
			name        text not null,
			description text not null,
			liquid      boolean not null,
			statsid     integer not null
		);

		drop table if exists currency;
		create table currency (
			id          integer primary key autoincrement,
			name        text not null,
			value       integer not null
		);

		drop table if exists groups;
		create table groups (
			id          integer primary key autoincrement,
			name        text not null,
			description text not null
		);

		drop table if exists inventory;
		create table inventory (
			id          integer primary key autoincrement,
			size        integer not null
		);

		drop table if exists inventoryslots;
		create table inventoryslots (
			id          integer primary key autoincrement,
			inventoryid integer not null,
			itemid      integer not null,
			itemtype    tinyint not null,
			count       tinyint not null
		);

		drop table if exists portals;
		create table portals (
			id          integer primary key autoincrement,
			zoneid      integer not null,
			direction   tinyint not null
		);

		drop table if exists races;
		create table races (
			id          integer primary key autoincrement,
			name        text not null,
			description text not null,
			statsid     integer not null
		);

		drop table if exists stats;
		create table stats (
			id   integer primary key autoincrement,
			hp   tinyint not null,
			mp   tinyint not null,
			ap   tinyint not null,
			def  tinyint not null,
			agi  tinyint not null,
			str  tinyint not null,
			wis  tinyint not null,
			luc  tinyint not null,
			chr  tinyint not null,
			per  tinyint not null
		);

		drop table if exists weapons;
		create table weapons (
			id          integer primary key autoincrement,
			name        text not null,
			description text not null,
			type        tinyint not null,
			dmglo       smallint not null,
			dmghi       smallint not null,
			statsid     integer not null
		);

		drop table if exists worlds;
		create table worlds (
			id            integer primary key autoincrement,
			created       integer NOT NULL,
			name          text not null,
			description   text not null,
			logo          text not null,
			motd          text,
			defaultzoneid integer not null,
			allowregister boolean not null,
			levelcap      smallint not null
		);

		drop table if exists zones;
		create table zones (
			id            integer primary key autoincrement,
			name          text not null,
			description   text not null,
			lighting      text not null,
			smell         text not null,
			sound         text not null
		);

		end transaction;`,
	); err != nil {
		rwl.Unlock()
		return
	}

	rwl.Unlock()
	return //this.SetWorld(world)
}
