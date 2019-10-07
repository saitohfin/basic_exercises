package mocking

//SpySleeper class to mock an Sleeper
type SpySleeper struct {
	Calls int
}

//Sleep increments the counter of this call was used
func (s *SpySleeper) Sleep(){
	s.Calls++
}