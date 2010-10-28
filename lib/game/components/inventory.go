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
	Slots []*InventorySlot
	Size  int
}

func NewInventory(size int) *Inventory {
	v := new(Inventory)
	v.Slots = make([]*InventorySlot, 0, size)
	v.Size = size
	v.Id = 1
	return v
}

func (this *Inventory) Clear() {
	this.Slots = make([]*InventorySlot, 0, this.Size)
}

func (this *Inventory) Resize(size int) {
	if size == this.Size {
		return
	}

	if size > this.Size {
		cp := make([]*InventorySlot, this.Size, size)
		copy(cp, this.Slots)
		this.Slots = cp
	} else {
		cp := make([]*InventorySlot, size, size)
		copy(cp, this.Slots[0:size])
		this.Slots = cp
	}

	this.Size = size
}

func (this *Inventory) Add(item Item) (err os.Error) {
	sz := len(this.Slots)

	if sz >= cap(this.Slots) {
		return os.NewError(ErrInventoryFull)
	}

	slot := new(InventorySlot)
	slot.Id = int64(sz + 1)
	slot.Item = item
	slot.Count = 1

	this.Slots = this.Slots[0 : sz+1]
	this.Slots[sz] = slot
	return
}
