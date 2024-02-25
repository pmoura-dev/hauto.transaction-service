package dataaccess

import (
	"errors"
)

var ErrDeviceNotFound = errors.New("device not found")

var ErrDeviceStateNotFound = errors.New("device state not found")
