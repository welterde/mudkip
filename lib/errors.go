package lib

import "os"

const (
	_ uint8 = iota
	EUnknownError
	EUnknownMessage
	EInvalidUsername
	EInvalidPassword
	EUsernameExists
	EUnknownObject
	ETypeMismatch
	EUnknownUser
)

var (
	ErrUnknownError    = os.NewError("Unknown error")
	ErrUnknownMessage  = os.NewError("Unknown message")
	ErrInvalidUsername = os.NewError("Username can not exceed 50 bytes")
	ErrInvalidPassword = os.NewError("Password can not exceed 50 bytes")
	ErrUsernameExists  = os.NewError("Specified username already exists")
	ErrUnknownObject   = os.NewError("Unknown object")
	ErrTypeMismatch    = os.NewError("Stored object and requested object differ in type")
	ErrUnknownUser     = os.NewError("Unknown user")
)

func errToInt(err os.Error) uint8 {
	switch err {
	case ErrUnknownMessage:
		return EUnknownMessage
	case ErrInvalidUsername:
		return EInvalidUsername
	case ErrInvalidPassword:
		return EInvalidPassword
	case ErrUsernameExists:
		return EUsernameExists
	case ErrUnknownObject:
		return EUnknownObject
	case ErrTypeMismatch:
		return ETypeMismatch
	case ErrUnknownUser:
		return EUnknownUser
	}
	return EUnknownError
}

func intToErr(errno uint8) os.Error {
	switch errno {
	case EUnknownMessage:
		return ErrUnknownMessage
	case EInvalidUsername:
		return ErrInvalidUsername
	case EInvalidPassword:
		return ErrInvalidPassword
	case EUsernameExists:
		return ErrUsernameExists
	case EUnknownObject:
		return ErrUnknownObject
	case ETypeMismatch:
		return ErrTypeMismatch
	case EUnknownUser:
		return ErrUnknownUser
	}
	return ErrUnknownError
}
