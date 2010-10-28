package store

import "os"
import "sync"
import "mudkip/lib"

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
