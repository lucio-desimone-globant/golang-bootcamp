package db

import (
	"errors"
)

const recordNotFound = "Record not found"

//ErrFind represent any unsuccessfull findRecord in database
func ErrFind() error {
	return errors.New(recordNotFound)
}
