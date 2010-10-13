package lib

import "bufio"
import "os"

type Character struct {
	id          uint16
	name        string
	description string
}

func NewCharacter() *Character {
	v := new(Character)
	return v
}

func (this *Character) Type() uint8             { return OTCharacter }
func (this *Character) Id() uint16              { return this.id }
func (this *Character) SetId(id uint16)         { this.id = id }
func (this *Character) Name() string            { return this.name }
func (this *Character) SetName(v string)        { this.name = v }
func (this *Character) Description() string     { return this.description }
func (this *Character) SetDescription(v string) { this.description = v }

func (this *Character) Pack(w *bufio.Writer) (err os.Error) {
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

func (this *Character) Unpack(r *bufio.Reader) (err os.Error) {
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
