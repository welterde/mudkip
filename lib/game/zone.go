package lib

import "io"
import "os"

type Zone struct {
	id          uint16
	name        string
	description string
}

func NewZone(name, desc string) *Zone {
	v := new(Zone)
	v.name = name
	v.description = desc
	return v
}

func (this *Zone) Type() uint8                       { return OTZone }
func (this *Zone) Id() uint16                        { return this.id }
func (this *Zone) SetId(id uint16)                   { this.id = id }
func (this *Zone) Name() string                      { return this.name }
func (this *Zone) Description() string               { return this.description }
func (this *Zone) Pack(w io.Writer) (err os.Error)   { return }
func (this *Zone) Unpack(w io.Reader) (err os.Error) { return }
