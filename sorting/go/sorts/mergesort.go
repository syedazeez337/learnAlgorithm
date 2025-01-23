package sorts

// Merge two sorted slices into one sorted slice
func merge(left, right []int) []int {
	result := []int{}
	i, j := 0, 0

	// Compare elements from both slices and append the smaller one
	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	// Append remaining elements from left or right
	result = append(result, left[i:]...)
	result = append(result, right[j:]...)

	return result
}

// MergeSort recursively divides and sorts the array
func MergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr // Base case: array with 1 or no element is already sorted
	}

	// Divide the array into two halves
	mid := len(arr) / 2
	left := MergeSort(arr[:mid])
	right := MergeSort(arr[mid:])

	// Merge the two sorted halves
	return merge(left, right)
}
