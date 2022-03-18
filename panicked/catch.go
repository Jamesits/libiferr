package panicked

import (
	"errors"
	"fmt"
)

type f func()

// Catch stops a panic from stopping the world.
func Catch(f f) (err interface{}) {
	defer func() {
		err = recover()
		return
	}()

	f()
	return nil
}

// CatchError stops a panic from stopping the world and ensures the return value is a standard error.
func CatchError(f f) (err error) {
	rawError := Catch(f)
	switch rawError.(type) {
	case nil:
		return nil

	case error:
		return rawError.(error)

	case string:
		return errors.New(rawError.(string))

	default:
		return errors.New(fmt.Sprintf("%T: %v", rawError, rawError))
	}
}
