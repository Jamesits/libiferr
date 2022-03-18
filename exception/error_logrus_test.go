// +build libiferr.logrus

package exception

import (
	"errors"
	"github.com/sirupsen/logrus"
	"testing"
)

func TestFailLoggers(t *testing.T) {
	err1 := errors.New("test error")
	var err2 error = nil
	logger := logrus.WithField("source", "TestFailLoggers()")

	SoftFailWithContext(logger, "reason", err1)
	SoftFailWithContext(logger, "reason", err2)
	SoftFailWithReason("reason", err1)
	SoftFailWithReason("reason", err2)

	// You can't test fatal calls -- https://stackoverflow.com/a/46841524

	//HardFailWithContext(logger, "reason", err1)
	//HardFailWithContext(logger, "reason", err2)
	//HardFailWithReason("reason", err1)
	//HardFailWithReason("reason", err2)
}
