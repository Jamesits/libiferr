package lifecycle

import "sync"

type sleepLock struct {
	wg *sync.WaitGroup
}

// NewSleepLock creates a sleepLock object.
func NewSleepLock() *sleepLock {
	return &sleepLock{
		wg: &sync.WaitGroup{},
	}
}

// LockLocal effectively pauses the current goroutine.
func (sl *sleepLock) LockLocal() {
	sl.wg.Add(1)
	sl.wg.Wait()
}

// Unlock releases the goroutine from pausing.
func (sl *sleepLock) Unlock() {
	sl.wg.Done()
}
