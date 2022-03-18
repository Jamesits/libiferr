package panicked

import (
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
	assert.Equal(t, err, testString)
}

func TestCatchNotPanicked(t *testing.T) {
	defer func() {
		err := recover()
		assert.Nil(t, err)
	}()

	err := Catch(func() {

	})
	assert.Equal(t, err, nil)
}
