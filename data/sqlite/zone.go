package store

import "os"
import "mudkip/lib"

func (this *Store) GetZone(int64) (*lib.Zone, os.Error) {
	return nil, nil
}

func (this *Store) SetZone(*lib.Zone) os.Error {
	return nil
}

func (this *Store) ZoneExists(id int64) (bool, os.Error) {
	var err os.Error

	if this.qry, err = this.conn.Prepare("select count(*) from zones where id=?"); err != nil {
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
