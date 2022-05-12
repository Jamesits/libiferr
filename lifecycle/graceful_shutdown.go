package lifecycle

import (
	"context"
	"os"
	"os/signal"
)

type GracefulShutdownCallbackFunction func() (exitCode int)

// WaitForKeyboardInterruptAsync starts a new goroutine waiting for Ctrl-C events. If an event is received, it runs the
// hook function and quits.
// It returns a stop function which can be used to disable the event hook. 
func WaitForKeyboardInterruptAsync(hook GracefulShutdownCallbackFunction) func() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)

	go func() {
		<-ctx.Done()
		stop()
		GracefulShutdownCallbackFunction()
	}()

	return stop
}
