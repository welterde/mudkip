package lib

import "bufio"
import "os"
import "utf8"

type Zone struct {
	id          uint16
	name        string
	description string
}

func NewZone() *Zone {
	v := new(Zone)
	return v
}

func (this *Zone) Type() uint8     { return OTZone }
func (this *Zone) Id() uint16      { return this.id }
func (this *Zone) SetId(id uint16) { this.id = id }
func (this *Zone) Name() string    { return this.name }
func (this *Zone) SetName(v string) {
	if str := utf8.NewString(v); str.RuneCount() > LimitName {
		this.name = str.Slice(0, LimitName+1)
	} else {
		this.name = v
	}
}

func (this *Zone) Description() string { return this.description }
func (this *Zone) SetDescription(v string) {
	if str := utf8.NewString(v); str.RuneCount() > LimitDescription {
		this.description = str.Slice(0, LimitDescription+1)
	} else {
		this.description = v
	}
}

func (this *Zone) Pack(w *bufio.Writer) (err os.Error) {
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

func (this *Zone) Unpack(r *bufio.Reader) (err os.Error) {
	var data []byte

	if data, err = r.ReadBytes(0x00); err != nil {
		return
	}

	if len(data) > 0 {
		data = data[0 : len(data)-1]
		this.name = string(data)
	}

	if data, err = r.ReadBytes(0x00); err != nil {
		return
	}

	if len(data) > 0 {
		data = data[0 : len(data)-1]
		this.description = string(data)
	}

	return
}
