package iteration

var amount int = 5

func Repeat(text string) string {
	result := ""
	for repeated := 0; repeated < amount; repeated++ {
		result += text
	}
	return result
}
