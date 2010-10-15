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
	EInvalidCredentials
	EDuplicateUser
	EUserLoggedIn
	EUserNotLoggedIn
	ENoWorldInfo
)

var (
	ErrUnknownError       = os.NewError("Unknown error")
	ErrUnknownMessage     = os.NewError("Unknown message")
	ErrInvalidUsername    = os.NewError("Username can not exceed 50 bytes")
	ErrInvalidPassword    = os.NewError("Password can not exceed 50 bytes")
	ErrUsernameExists     = os.NewError("Specified username already exists")
	ErrUnknownObject      = os.NewError("Unknown object")
	ErrTypeMismatch       = os.NewError("Stored object and requested object differ in type")
	ErrUnknownUser        = os.NewError("Unknown user")
	ErrInvalidCredentials = os.NewError("Invalid user credentials supplied")
	ErrDuplicateUser      = os.NewError("User already exists")
	ErrUserLoggedIn       = os.NewError("User already logged in")
	ErrUserNotLoggedIn    = os.NewError("User not logged in")
	ErrNoWorldInfo        = os.NewError("No world info defined in datastore")
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
	case ErrInvalidCredentials:
		return EInvalidCredentials
	case ErrDuplicateUser:
		return EDuplicateUser
	case ErrUserLoggedIn:
		return EUserLoggedIn
	case ErrUserNotLoggedIn:
		return EUserNotLoggedIn
	case ErrNoWorldInfo:
		return ENoWorldInfo
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
	case EInvalidCredentials:
		return ErrInvalidCredentials
	case EDuplicateUser:
		return ErrDuplicateUser
	case EUserLoggedIn:
		return ErrUserLoggedIn
	case EUserNotLoggedIn:
		return ErrUserNotLoggedIn
	case ENoWorldInfo:
		return ErrNoWorldInfo
	}
	return ErrUnknownError
}
