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
	for _, value := range numbersToSum {
		sums = append(sums, Sum(value))
	}
	return
}

//SumAllTails summatory of all numbers except the first per each array
func SumAllTails(numbersToSum ...[]int) (sums []int) {
	for _, value := range numbersToSum {
		if len(value) == 0 {
			sums = append(sums, 0)
		} else {
			tail := value[1:]
			sums = append(sums, Sum(tail))
		}
	}
	return
}
