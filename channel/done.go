package channel

// Done observes a "done" channel and return if it is actually done, without blocking.
func Done(done chan struct{}) bool {
	select {
	case <-done:
		return true
	default:
		return false
	}
}

// random wild aliases

//goland:noinspection ALL
var Doneˀ = Done

//goland:noinspection ALL
var DoneHuh = Done

//goland:noinspection ALL
var Done了没 = Done

//goland:noinspection ALL
var Doneしましたと言うは教えて来れますか = Done

//goland:noinspection ALL
var Done했다고말해줄수있습니까 = Done
