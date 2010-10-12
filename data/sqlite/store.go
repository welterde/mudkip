package store

import "os"
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
	return
}

func (this *Store) GetObject(objtype uint8, id uint16) (obj lib.Object, err os.Error) {
	return
}

func (this *Store) SetObject(lib.Object) (err os.Error) {
	return
}
