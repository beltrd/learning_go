package main

import "fmt"

func main() {
	fmt.Println("HEL:LOASD")
	// Arrays
	var fruitArray [2]string

	// Assign values
	fruitArray[0] = "Apple"
	fruitArray[1] = "Orange"

	// Decalre and assign
	fruitArr := [2]string{"Apple", "Naranja"}

	// Slice
	fruitSlice := []string{"Apple", "Pinnapple", "Orange", "Grape", "Berry"}

	fmt.Println(fruitArray)
	fmt.Println(fruitArray[1])
	fmt.Println(fruitArr[1])
	fmt.Println(fruitSlice)
	fmt.Println(len(fruitSlice))
	fmt.Println(fruitSlice[1:4])
}
