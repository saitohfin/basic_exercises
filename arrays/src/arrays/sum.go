package main

//Sum make a summatory of all numbers of an array
func Sum(numbers []int) int {
	numberResult := 0

	for _, number := range numbers {
		numberResult += number
	}

	return numberResult
}

//SumAll do a summatory per each array of numbers
func SumAll(numbersToSum ...[]int) (sums []int) {
	amountArrays := len(numbersToSum)
	sums = make([]int, amountArrays)
	for position, value := range numbersToSum {
		sums[position] = Sum(value)
	}
	return
}

func SumAllTails(numbersToSum ...[]int) (sums []int) {
	amountArrays := len(numbersToSum)
	sums = make([]int, amountArrays)
	for position, value := range numbersToSum {
		if len(value) == 0 {
			sums[position] = 0
		} else {
			tail := value[1:]
			sums[position] = Sum(tail)
		}
	}
	return
}
