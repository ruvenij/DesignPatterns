package sort

func PerformInsertionSort(input []int) {
	for i := 1; i < len(input); i++ {
		current := input[i]
		insertIndex := i

		for j := i - 1; j >= 0; j-- {
			if input[j] > current {
				input[j+1] = input[j]
				insertIndex = j
			}
		}

		input[insertIndex] = current

		PrintArray(input)
	}
}
