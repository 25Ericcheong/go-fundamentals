package arrays_and_slices

func Sum(numbers [5]int) int {
	total := 0
	for i := 0; i < len(numbers); i++ {
		total += numbers[i]
	}

	return total
}
