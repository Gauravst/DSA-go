package main

import (
	"fmt"
)

func bubbleSort(arr *[]int) {
	if len(*arr) <= 1 {
		return
	}

	for i := 0; i < len(*arr)-1; i++ {
		swapped := false
		for j := 0; j < len(*arr)-i-1; j++ {
			if (*arr)[j] > (*arr)[j+1] {
				(*arr)[j], (*arr)[j+1] = (*arr)[j+1], (*arr)[j]
				swapped = true
			}
		}

		if !swapped {
			break
		}
	}
}

func main() {
	arr := []int{1, 4, 2, 6, 1}

	fmt.Println("Before bubble sort")
	for _, value := range arr {
		fmt.Printf("%d ", value)
	}

	bubbleSort(&arr)

	fmt.Println("\nAfter bubble sort")
	for _, value := range arr {
		fmt.Printf("%d ", value)
	}

}
