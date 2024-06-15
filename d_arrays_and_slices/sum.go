package arrays_and_slices

func Sum(numbers []int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}

	return total
}

func SumAll(arrOfNums ...[]int) []int {
	numOfSlices := len(arrOfNums)
	sums := make([]int, numOfSlices)

	for i, numbers := range arrOfNums {
		sums[i] = Sum(numbers)
	}

	return sums
}
