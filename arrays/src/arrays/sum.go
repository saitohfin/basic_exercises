package main

//Sum make a summatory of all numbers of an array
func Sum(numbers []int) int {
	numberResult := 0

	for _, number := range numbers {
		numberResult += number
	}

	return numberResult
}

func SumAll(numbersToSum ...[]int) (sums []int) {
	amountArrays := len(numbersToSum)
	sums = make([]int, amountArrays)
	for position, value := range numbersToSum {
		sums[position] = Sum(value)
	}
	return
}
