package lib

import "os"

type Item interface{}

type InventorySlot struct {
	Id    int64
	Item  Item
	Count int
}

type Inventory struct {
	Id    int64
	slots []*InventorySlot
}

func NewInventory(size int) *Inventory {
	v := new(Inventory)
	v.slots = make([]*InventorySlot, 0, size)
	return v
}

func (this *Inventory) Clear() {
	this.slots = make([]*InventorySlot, 0, cap(this.slots))
}

func (this *Inventory) Add(item Item) (err os.Error) {
	sz := len(this.slots)

	if sz >= cap(this.slots) {
		return os.NewError(ErrInventoryFull)
	}

	this.slots = this.slots[0 : sz+1]
	this.slots[sz] = new(InventorySlot)
	this.slots[sz].Item = item
	this.slots[sz].Count = 1
	return
}
