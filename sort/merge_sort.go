package sort

func MergeSort(input []int) []int {
	if len(input) <= 1 {
		return input
	}

	mid := len(input) / 2
	left := MergeSort(input[:mid])
	right := MergeSort(input[mid:])

	return merge(left, right)

}

func merge(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))
	var i, j int
	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	result = append(result, left[i:]...)
	result = append(result, right[j:]...)
	return result
}
