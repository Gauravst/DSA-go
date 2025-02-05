package main

import "fmt"

func selectionSort(arr *[]int) {
	if len(*arr) <= 1 {
		return
	}

	for i := 0; i < len(*arr)-1; i++ {
		min := i
		for j := i + 1; j < len(*arr); j++ {
			if (*arr)[min] > (*arr)[j] {
				min = j
			}
		}

		(*arr)[i], (*arr)[min] = (*arr)[min], (*arr)[i]
	}
}

func main() {
	arr := []int{2, 55, 1, 25, 0, 566, 3, 45, 999}

	fmt.Println("Before selection sort")
	for _, value := range arr {
		fmt.Printf("%d ", value)
	}

	selectionSort(&arr)

	fmt.Println("\nAfter selection sort")
	for _, value := range arr {
		fmt.Printf("%d ", value)
	}
}
