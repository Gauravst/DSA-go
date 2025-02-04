// 👉 Problem: Given a sorted array, remove duplicates in-place and return the new length.
package main

import "fmt"

func removeDuplicate(arr []int) int {
	if len(arr) == 0 {
		return 0
	}

	write := 1
	for read := 1; read < len(arr); read++ {
		if arr[read] != arr[read-1] {
			arr[write] = arr[read]
			write++
		}
	}

	return write
}

func main() {
	arr := []int{1, 2, 2, 3, 4, 5}
	result := removeDuplicate(arr)
	fmt.Println(result)
}
