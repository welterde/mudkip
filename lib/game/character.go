package lib

import "io"
import "os"

type Character struct {
	id          uint16
	name        string
	description string
}

func NewCharacter(name, desc string) *Character {
	v := new(Character)
	v.name = name
	v.description = desc
	return v
}

func (this *Character) Type() uint8                       { return OTCharacter }
func (this *Character) Id() uint16                        { return this.id }
func (this *Character) SetId(id uint16)                   { this.id = id }
func (this *Character) Name() string                      { return this.name }
func (this *Character) Description() string               { return this.description }
func (this *Character) Pack(w io.Writer) (err os.Error)   { return }
func (this *Character) Unpack(w io.Reader) (err os.Error) { return }
