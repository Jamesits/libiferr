package testing_utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"unsafe"
)

// AssertSameSize tests if 2 data types are of the same size in memory.
// It is designed to check for things like C unions, where one struct need to be manually padded to the size of another
// struct.
func AssertSameSize[Expected any, Actual any](t *testing.T) {
	assert.EqualValues(t, unsafe.Sizeof(*new(Expected)), unsafe.Sizeof(*new(Actual)))
}

func AssertDifferentSize[Expected any, Actual any](t *testing.T) {
	assert.NotEqualValues(t, unsafe.Sizeof(*new(Expected)), unsafe.Sizeof(*new(Actual)))
}
