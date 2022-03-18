package panicked

type f func()

// Catch stops a panic from stopping the world.
func Catch(f f) (err interface{}) {
	defer func() {
		err = recover()
		return
	}()

	f()
	return nil
}
