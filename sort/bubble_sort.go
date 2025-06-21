package sort

import "fmt"

func PerformBubbleSort(input []int) {
	isSwapped := false

	for j := 0; j < len(input)-1; j++ {
		isSwapped = false
		for i := 0; i < len(input)-1-j; i++ {
			if input[i] > input[i+1] {
				isSwapped = true

				temp := input[i]
				input[i] = input[i+1]
				input[i+1] = temp
			}
		}

		if !isSwapped {
			break
		}

		PrintArray(input)
	}
}

func PrintArray(input []int) {
	for _, val := range input {
		fmt.Print(val, " ")
	}

	fmt.Println()
}
