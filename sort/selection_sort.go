package sort

func PerformSelectionSort(input []int) {
	for i := 0; i < len(input)-1; i++ {
		currentMin := input[i]
		minIndex := i
		for j := i + 1; j < len(input); j++ {
			if input[j] < currentMin {
				minIndex = j
				currentMin = input[j]
			}
		}

		input[i], input[minIndex] = input[minIndex], input[i]

		PrintArray(input)
	}
}
