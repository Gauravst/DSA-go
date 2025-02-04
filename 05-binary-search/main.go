package main

import "fmt"

func binarySearch(arr []int, target int, start int, end int) int {

	if start > end {
		return -1
	}

	mid := start + (end-start)/2

	if arr[mid] == target {
		return mid
	}

	if arr[mid] > target {
		return binarySearch(arr, target, start, mid)
	} else {
		return binarySearch(arr, target, mid+1, end)
	}
}

func main() {
	arr := []int{1, 2, 3, 4, 5}
	target := 5

	start := 0
	end := len(arr) - 1

	result := binarySearch(arr, target, start, end)
	if result != -1 {
		fmt.Printf("Element found on index - %d\n", result)
		return
	}

	fmt.Println("Element not found")
}
