package channel

import "testing"

func TestDone(t *testing.T) {
	done1 := make(chan struct{})
	done2 := make(chan struct{})
	close(done1)

	if !Doneˀ(done1) {
		t.Fail()
	}

	if Doneˀ(done2) {
		t.Fail()
	}

	close(done2)

	if !Doneˀ(done2) {
		t.Fail()
	}
}
