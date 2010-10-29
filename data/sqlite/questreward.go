package store

import "os"
import "mudkip/lib"

func (this *Store) GetQuestReward(int64) (*lib.QuestReward, os.Error) {
	return nil, nil
}

func (this *Store) SetQuestReward(*lib.QuestReward) os.Error {
	return nil
}

func (this *Store) QuestRewardExists(id int64) (bool, os.Error) {
	var err os.Error

	if this.qry, err = this.conn.Prepare("select count(*) from questrewards where id=?"); err != nil {
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
