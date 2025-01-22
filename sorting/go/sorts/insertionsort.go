package sorts

func InsertionSort(arr []int) []int {
	n := len(arr)

	for i := 1; i < n; i++ {
		key := arr[i]
		j := i - 1

		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}

		arr[j+1] = key

		// fmt.Printf("i = %v -> array %v\n", i, arr)
	}

	return arr
}
