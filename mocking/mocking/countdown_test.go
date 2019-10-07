package mocking

import (
	"bytes"
	"testing"
)

func TestCountdown(t *testing.T) {
	buffer := &bytes.Buffer{}
	sleeperSpy := &SpySleeper{}

	Countdown(buffer, sleeperSpy)

	got := buffer.String()
	want :=
		`3
2
1
Go!`

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}

	if sleeperSpy.Calls != 4 {
		t.Errorf("not enough calls to sleeper, want 4 got %d", sleeperSpy.Calls)
	}
}
