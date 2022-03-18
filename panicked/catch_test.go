package panicked

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCatchPanicked(t *testing.T) {
	defer func() {
		err := recover()
		assert.Nil(t, err)
	}()

	testString := "test panicked.Catch"
	err := Catch(func() {
		panic(testString)
	})

	assert.Equal(t, testString, err)
}

func TestCatchNotPanicked(t *testing.T) {
	defer func() {
		err := recover()
		assert.Nil(t, err)
	}()

	err := Catch(func() {

	})

	assert.Equal(t, nil, err)
}

func TestCatchErrorPanickedError(t *testing.T) {
	testErr := errors.New("test panicked.Catch")
	err := CatchError(func() {
		panic(testErr)
	})

	assert.Equal(t, testErr, err)
}

func TestCatchErrorPanickedString(t *testing.T) {
	testString := "test panicked.Catch"
	err := CatchError(func() {
		panic(testString)
	})

	assert.Equal(t, errors.New(testString), err)
}

func TestCatchErrorPanickedUnknownType(t *testing.T) {
	err := CatchError(func() {
		panic(1)
	})

	assert.Equal(t, errors.New("int: 1"), err)
}

func TestCatchErrorNotPanicked(t *testing.T) {
	err := CatchError(func() {

	})

	assert.Nil(t, err)
}
