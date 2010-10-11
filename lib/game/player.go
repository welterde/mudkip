package lib

import "io"
import "os"

type Player struct {
	id          uint16
	name        string
	description string
}

func NewPlayer(id uint16, name, desc string) *Player {
	v := new(Player)
	v.id = id
	v.name = name
	v.description = desc
	return v
}

func (this *Player) Type() uint8                       { return OTPlayer }
func (this *Player) Id() uint16                        { return this.id }
func (this *Player) Name() string                      { return this.name }
func (this *Player) Description() string               { return this.description }
func (this *Player) Pack(w io.Writer) (err os.Error)   { return }
func (this *Player) Unpack(w io.Reader) (err os.Error) { return }
