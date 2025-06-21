package sort

func PerformQuickSort(input []int) []int {
	if len(input) <= 1 {
		return input
	}

	pivot := input[0]
	var high []int
	var low []int

	for _, val := range input[1:] {
		if val <= pivot {
			low = append(low, val)
		} else {
			high = append(high, val)
		}
	}

	return append(append(PerformQuickSort(low), pivot), PerformQuickSort(high)...)
}
