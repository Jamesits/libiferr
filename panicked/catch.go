package panicked

import (
	"errors"
	"fmt"
	"runtime/debug"
)

type f func()

// Catch stops a panic from stopping the world.
func Catch(f f) (err interface{}, stack []byte) {
	defer func() {
		err = recover()
		stack = debug.Stack()
		return
	}()

	f()
	return nil, nil
}

// CatchError stops a panic from stopping the world and ensures the return value is a standard error.
func CatchError(f f) (err error, stack []byte) {
	var rawError interface{}
	rawError, stack = Catch(f)
	switch rawError.(type) {
	case nil:
		err = nil

	case error:
		err = rawError.(error)

	case string:
		err = errors.New(rawError.(string))

	default:
		err = errors.New(fmt.Sprintf("%T: %v", rawError, rawError))
	}

	return
}
