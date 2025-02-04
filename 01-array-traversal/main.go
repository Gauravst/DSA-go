package main

import "fmt"

func main() {
	arr := [5]int{1, 4, 6, 4, 9}

	// using range
	fmt.Println("Using Normal Range")
	for index, value := range arr {
		fmt.Printf("index = %d --- value = %d \n", index, value)
	}

	// using for loop
	fmt.Println("\nUsing Normal For Loop")
	for i := 0; i < len(arr); i++ {
		fmt.Printf("index = %d --- value = %d \n", i, arr[i])
	}

}
