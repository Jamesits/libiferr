// +build !libiferr.logrus

package exception

import "log"

func SoftFailWithReason(t string, e error) {
	if e != nil {
		log.Printf("%s: %s", t, e)
	}
}

func HardFailWithReason(t string, e error) {
	if e != nil {
		log.Fatalf("%s: %s", t, e)
	}
}
