package lib

import "io"
import "os"

type World struct {
	id          uint16
	name        string
	description string
	Zones       map[uint16]*Zone
	Players     map[uint16]*Player
}

func NewWorld(id uint16, name, desc string) *World {
	v := new(World)
	v.id = id
	v.name = name
	v.description = desc
	v.Zones = make(map[uint16]*Zone)
	v.Players = make(map[uint16]*Player)
	return v
}

func (this *World) Type() uint8                       { return OTWorld }
func (this *World) Id() uint16                        { return this.id }
func (this *World) Name() string                      { return this.name }
func (this *World) Description() string               { return this.description }
func (this *World) Pack(w io.Writer) (err os.Error)   { return }
func (this *World) Unpack(w io.Reader) (err os.Error) { return }
