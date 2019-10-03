package main

//Sum make a summatory of all numbers of an array
func Sum(numbers [5]int) int {
	number := 0
	amountNumbers := len(numbers)

	for position := 0; position < amountNumbers; position++ {
		number += numbers[position]
	}

	return number

}
