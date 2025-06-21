package search

func BinarySearch(array []int, search int) int {
	low := 0
	high := len(array) - 1

	for low <= high {
		mid := low + (high-low)/2

		if array[mid] == search {
			return mid
		} else if array[mid] > search {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	
	return -1
}
