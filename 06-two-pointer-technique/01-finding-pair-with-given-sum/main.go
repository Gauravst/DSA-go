// 👉 Problem: Given a sorted array, find if there exists a pair with a given sum.
package main

import "fmt"

func findPair(arr []int, target int) bool {
	start := 0
	end := len(arr) - 1

	for start < end {
		sum := arr[start] + arr[end]

		if sum == target {
			return true
		}

		if sum < target {
			start++
		} else {
			end--
		}
	}

	return false
}

func main() {

	arr := []int{1, 2, 3, 4, 5}
	target := 6

	result := findPair(arr, target)
	if result {
		fmt.Println("Pair found")
		return
	}

	fmt.Println("Pair not found")

}
