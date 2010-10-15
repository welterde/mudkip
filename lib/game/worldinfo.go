package lib

type WorldInfo struct {
	Id            uint16
	Name          string
	Logo          string
	Description   string
	Created       int64
	Motd          string
	DefaultZone   uint16
	AllowRegister bool
}
