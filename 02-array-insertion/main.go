package main

import "fmt"

func addNewElement(arr *[]int, element int, position int) {
	if position < 0 || position > len(*arr) {
		fmt.Println("Invalid position")
		return
	}

	*arr = append((*arr)[:position], append([]int{element}, (*arr)[position:]...)...)
}

func main() {
	arr := []int{10, 20, 30, 40}
	element := 25
	position := 2

	fmt.Println("Before adding element")
	for index, value := range arr {
		fmt.Printf("index is %d and value is %d \n", index, value)
	}

	addNewElement(&arr, element, position)

	fmt.Println("\nAfter adding element")
	for index, value := range arr {
		fmt.Printf("index is %d and value is %d \n", index, value)
	}
}
