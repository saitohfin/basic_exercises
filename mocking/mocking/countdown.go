package mocking

import (
	"fmt"
	"io"
)

const finalWord = "Go!"
const countdownStart = 3

//Countdown print a countdown in console
func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		sleeper.Sleep()
		fmt.Fprintln(out, i)
	}
	sleeper.Sleep()
	fmt.Fprint(out, finalWord)
}
