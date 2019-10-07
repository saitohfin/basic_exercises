package mocking

import "time"

//DefaultSleeper is the default implementation for the Sleeper interface
type DefaultSleeper struct {
}

//Sleep one second
func (s *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}
