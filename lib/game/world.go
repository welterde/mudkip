package lib

import "bufio"
import "os"

type World struct {
	id          uint16
	name        string
	description string
	Zones       map[uint16]*Zone
}

func NewWorld() *World {
	v := new(World)
	v.Zones = make(map[uint16]*Zone)
	return v
}

func (this *World) Type() uint8                       { return OTWorld }
func (this *World) Id() uint16                        { return this.id }
func (this *World) SetId(id uint16)                   { this.id = id }
func (this *World) Name() string                      { return this.name }
func (this *World) SetName(v string)                   { this.name = v }
func (this *World) Description() string               { return this.description }
func (this *World) SetDescription(v string)              { this.description = v }

func (this *World) Pack(w *bufio.Writer) (err os.Error) {
	if _, err = w.WriteString(this.name); err == nil {
		if err = w.WriteByte(0x00); err != nil {
			return
		}
	}

	if _, err = w.WriteString(this.description); err == nil {
		if err = w.WriteByte(0x00); err != nil {
			return
		}
	}

	return
}

func (this *World) Unpack(r *bufio.Reader) (err os.Error) {
	var data []byte

	if data, err = r.ReadBytes(0x00); err != nil {
		return
	}

	if len(data) > 0 {
		data = data[0:len(data)-1]
		this.name = string(data)
	}

	if data, err = r.ReadBytes(0x00); err != nil {
		return
	}

	if len(data) > 0 {
		data = data[0:len(data)-1]
		this.description = string(data)
	}

	return
}
