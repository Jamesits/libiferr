package lifecycle

import (
	"os"
	"os/signal"
)

type GracefulShutdownCallbackFunction func() (exitCode int)

// WaitForKeyboardInterruptAsync starts a new goroutine waiting for Ctrl-C events. If an event is received, it runs the
// hook function and quits.
func WaitForKeyboardInterruptAsync(hook GracefulShutdownCallbackFunction) {
	var signalChannel = make(chan os.Signal, 1)
	signal.Notify(signalChannel, os.Interrupt)

	go func() {
		for range signalChannel {
			os.Exit(hook())
		}
	}()
}
