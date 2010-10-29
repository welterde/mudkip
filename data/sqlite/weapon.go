package store

import "os"
import "mudkip/lib"

func (this *Store) GetWeapon(int64) (*lib.Weapon, os.Error) {
	return nil, nil
}

func (this *Store) SetWeapon(*lib.Weapon) os.Error {
	return nil
}

func (this *Store) WeaponExists(id int64) (bool, os.Error) {
	var err os.Error

	if this.qry, err = this.conn.Prepare("select count(*) from weapons where id=?"); err != nil {
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