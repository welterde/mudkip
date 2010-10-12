package store

import "os"
import "bytes"
import "mudkip/lib"
import "sqlite"

type Store struct {
	conn *sqlite.Conn
	qry  *sqlite.Stmt
}

func New() lib.DataStore {
	s := new(Store)
	return s
}

func (this *Store) Open(params map[string]string) (err os.Error) {
	if this.conn != nil {
		return
	}

	if this.conn, err = sqlite.Open(params["file"]); err != nil {
		return
	}

	return this.init()
}

func (this *Store) Close() {
	this.qry = nil

	if this.conn != nil {
		this.conn.Close()
		this.conn = nil
	}
}

func (this *Store) GetObject(objtype uint8, id uint16) (obj lib.Object, err os.Error) {
	return
}

func (this *Store) SetObject(lib.Object) (err os.Error) {
	return
}

func (this *Store) init() (err os.Error) {
	var d []byte

	buf := bytes.NewBuffer(d)
	buf.WriteString("create table if not exists world (i INTEGER, s VARCHAR(20));")
	buf.WriteString("create table if not exists players (i INTEGER, s VARCHAR(20));")

	if this.qry, err = this.conn.Prepare(buf.String()); err != nil {
		return
	}

	if err = this.qry.Exec(); err != nil {
		return
	}

	for this.qry.Next() {
	}

	this.qry.Finalize()
	return
}
