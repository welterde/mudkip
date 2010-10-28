package builder

// A group can be anything from a clan, guild or simple fellowship of friends.
type Group struct {
	Name        string
	Description string
	Members     []int
}

func NewGroup() *Group {
	v := new(Group)
	v.Members = make([]int, 0, 32)
	return v
}

// Add a new member
func (this *Group) AddMember(id int) {
	sz := len(this.Members)

	if sz >= cap(this.Members) {
		cp := make([]int, sz, sz+32)
		copy(cp, this.Members)
		this.Members = cp
	}

	this.Members = this.Members[0 : sz+1]
	this.Members[sz] = id
}
