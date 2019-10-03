package main

//Sum make a summatory of all numbers of an array
func Sum(numbers []int) int {
	numberResult := 0

	for _, number := range numbers {
		numberResult += number
	}

	return numberResult

}
