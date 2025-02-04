package main

import "fmt"

func deleteElement(arr *[]int, pos int) {
	if pos < 0 || pos >= len(*arr) {
		fmt.Println("Invalid Position")
		return
	}

	*arr = append((*arr)[:pos], (*arr)[pos+1:]...)
}

func main() {
	arr := []int{10, 20, 25, 30, 40}
	pos := 2

	fmt.Println("Before deletion element")
	for index, value := range arr {
		fmt.Printf("index is %d and value is %d \n", index, value)
	}

	deleteElement(&arr, pos)

	fmt.Println("\nAfter deletion element")
	for index, value := range arr {
		fmt.Printf("index is %d and value is %d \n", index, value)
	}

}
