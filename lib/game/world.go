package lib

import "io"
import "os"

type World struct {
	id          uint16
	name        string
	description string
	Zones       map[uint16]*Zone
}

func NewWorld(name, desc string) *World {
	v := new(World)
	v.name = name
	v.description = desc
	v.Zones = make(map[uint16]*Zone)
	return v
}

func (this *World) Type() uint8                       { return OTWorld }
func (this *World) Id() uint16                        { return this.id }
func (this *World) SetId(id uint16)                   { this.id = id }
func (this *World) Name() string                      { return this.name }
func (this *World) Description() string               { return this.description }
func (this *World) Pack(w io.Writer) (err os.Error)   { return }
func (this *World) Unpack(w io.Reader) (err os.Error) { return }
