package iteration

import "strings"

var defaultTimes int = 5

//Repeat default times the text passed as parameter
func Repeat(text string) string {
	return RepeatAmount(text, defaultTimes)
}

//RepeatAmount amount times the text passed as parameter
func RepeatAmount(text string, times int) string {
	return strings.Repeat(text, times)
}
