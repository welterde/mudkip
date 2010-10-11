package lib

import "os"

const (
	_ uint8 = iota
	EUnknownError
	EUnknownMessage
	EInvalidUsername
	EInvalidPassword
)

var (
	ErrUnknownError    = os.NewError("Unknown error")
	ErrUnknownMessage  = os.NewError("Unknown message")
	ErrInvalidUsername = os.NewError("Username can not exceed 50 bytes")
	ErrInvalidPassword = os.NewError("Password can not exceed 50 bytes")
)

func ErrToInt(err os.Error) uint8 {
	switch err {
	case ErrUnknownMessage:
		return EUnknownMessage
	case ErrInvalidUsername:
		return EInvalidUsername
	case ErrInvalidPassword:
		return EInvalidPassword
	}
	return EUnknownError
}

func IntToErr(errno uint8) os.Error {
	switch errno {
	case EUnknownMessage:
		return ErrUnknownMessage
	case EInvalidUsername:
		return ErrInvalidUsername
	case EInvalidPassword:
		return ErrInvalidPassword
	}
	return ErrUnknownError
}
