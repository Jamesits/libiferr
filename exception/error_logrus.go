// +build libiferr.logrus

package exception

import log "github.com/sirupsen/logrus"

func SoftFailWithContext(context *log.Entry, t string, e error) {
	if e != nil {
		if context != nil {
			context.Errorf("%s: %s", t, e)
		} else {
			SoftFailWithReason(t, e)
		}
	}
}

func SoftFailWithReason(t string, e error) {
	if e != nil {
		log.Errorf("%s: %s", t, e)
	}
}

func HardFailWithContext(context *log.Entry, t string, e error) {
	if e != nil {
		if context != nil {
			context.Fatalf("%s: %s", t, e)
		} else {
			HardFailWithReason(t, e)
		}
	}
}

func HardFailWithReason(t string, e error) {
	if e != nil {
		log.Fatalf("%s: %s", t, e)
	}
}
