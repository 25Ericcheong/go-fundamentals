package sum

func Sum(numbers []int) int {
	sum := 0
	for i := 0; i < len(numbers); i++ {
		sum += numbers[i]
	}
	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	lengthOfArgs := len(numbersToSum)
	sums := make([]int, lengthOfArgs)

	for i, numbers := range numbersToSum {
		sums[i] = Sum(numbers)
	}

	// alternative - instead of predefining slice length
	// var sums []int
	// for i, numbers := range numbersToSum {
	// 	sums = append(sums, Sum(numbers))
	// }

	return sums
}
