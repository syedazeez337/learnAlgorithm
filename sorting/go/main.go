package main

import (
	"fmt"
	"sorting/sorts"
)

func main() {
	arr := []int{5,3,8,6,2}

	fmt.Println("Before sorting", arr)
	res := sorts.InsertionSort(arr)
	fmt.Println("After sorting", res)
}