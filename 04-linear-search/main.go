package main

import "fmt"

func main() {
	arr := [5]int{1, 5, 2, 6, 3}
	target := 6

	for i := 0; i < len(arr); i++ {
		if arr[i] == target {
			fmt.Println("index - ", i)
			return
		}
	}

	fmt.Println("Not found")

}
