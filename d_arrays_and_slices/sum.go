package arrays_and_slices

func Sum(numbers []int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}

	return total
}

func SumAll(arrOfNums ...[]int) []int {
	var sums []int
	for _, numbers := range arrOfNums {
		sums = append(sums, Sum(numbers))
	}

	return sums
}

func SumAllTails(arrOfNums ...[]int) []int {
	var sumsOfTails []int
	for _, nums := range arrOfNums {
		sumsOfTails = append(sumsOfTails, Sum(nums[1:]))
	}

	return sumsOfTails
}
