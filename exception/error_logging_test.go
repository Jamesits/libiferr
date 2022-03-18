// +build !libiferr.logrus

package exception

import (
	"errors"
	"testing"
)

func TestFailLoggers(t *testing.T) {
	err1 := errors.New("test error")
	var err2 error = nil

	SoftFailWithReason("reason", err1)
	SoftFailWithReason("reason", err2)

	// You can't test fatal calls -- https://stackoverflow.com/a/46841524
	//HardFailWithReason("reason", err1)
	//HardFailWithReason("reason", err2)
}
