package lifecycle

import "sync"

type sleepLock struct {
	wg *sync.WaitGroup
}

// NewSleepLock creates a sleepLock object. A sleepLock works like a `for();` except it does not waste a CPU core.
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

// UnlockFromRemote releases the goroutine from pausing.
func (sl *sleepLock) UnlockFromRemote() {
	sl.wg.Done()
}
