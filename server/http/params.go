package main

import "strconv"

type ParamList []string

func (this ParamList) S(i int, defval string) string {
	if -1 < i && i < len(this) {
		return this[i]
	}
	return defval
}

func (this ParamList) B(i int, defval bool) bool {
	if -1 < i && i < len(this) {
		if v, err := strconv.Atob(this[i]); err == nil {
			return v
		}
	}
	return defval
}

func (this ParamList) I(i, defval int) int {
	if -1 < i && i < len(this) {
		if v, err := strconv.Atoi(this[i]); err == nil {
			return v
		}
	}
	return defval
}

func (this ParamList) I8(i int, defval int8) int8 {
	return int8(this.I(i, int(defval)))
}

func (this ParamList) I16(i int, defval int16) int16 {
	return int16(this.I(i, int(defval)))
}

func (this ParamList) I32(i int, defval int32) int32 {
	return int32(this.I(i, int(defval)))
}

func (this ParamList) I64(i int, defval int64) int64 {
	if -1 < i && i < len(this) {
		if v, err := strconv.Atoi64(this[i]); err == nil {
			return v
		}
	}
	return defval
}

func (this ParamList) U(i int, defval uint) uint {
	if -1 < i && i < len(this) {
		if v, err := strconv.Atoui(this[i]); err == nil {
			return v
		}
	}
	return defval
}

func (this ParamList) U8(i int, defval uint8) uint8 {
	return uint8(this.U(i, uint(defval)))
}

func (this ParamList) U16(i int, defval uint16) uint16 {
	return uint16(this.U(i, uint(defval)))
}

func (this ParamList) U32(i int, defval uint32) uint32 {
	return uint32(this.U(i, uint(defval)))
}

func (this ParamList) U64(i int, defval uint64) uint64 {
	if -1 < i && i < len(this) {
		if v, err := strconv.Atoui64(this[i]); err == nil {
			return v
		}
	}
	return defval
}
