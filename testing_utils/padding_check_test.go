package testing_utils

import "testing"

type Test1 struct {
	A uint16
	B uint32
	C uint64
}

type Test2 struct {
	A uint32
	B uint32
	_ [8]uint8
}

type Test3 struct {
	A uint32
}

func TestAssertSameSize(t *testing.T) {
	AssertSameSize[Test1, Test2](t)
	AssertDifferentSize[Test1, Test3](t)
}
