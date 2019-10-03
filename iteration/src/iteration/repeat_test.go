package iteration

import "testing"

func TestRepeatDefaultTimes(t *testing.T) {
	repeated := Repeat("a")
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func TestRepeatAmountTimesReturnsTheTextRepeatedTheIndicatedTimes(t *testing.T) {
	repeated := RepeatAmount("ba", 8)
	expected := "babababababababa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}
