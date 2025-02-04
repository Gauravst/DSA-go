// 👉 Problem: Find the maximum sum of any contiguous subarray of size k.
package main

import "fmt"

func maxSum(arr []int, k int) int {

	if len(arr) < k {
		return -1
	}

	hightSum := 0
	windowSum := 0

	// first window sum
	for i := 0; i < k; i++ {
		windowSum += arr[i]
	}

	for i := k; i < len(arr); i++ {
		windowSum += arr[i] - arr[i-k]
		if windowSum > hightSum {
			hightSum = windowSum
		}

	}

	return hightSum
}

func main() {
	arr := []int{1, 2, 3, 4, 5}
	k := 3

	result := maxSum(arr, k)
	fmt.Println(result)
}
