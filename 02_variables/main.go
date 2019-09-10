package main

import "fmt"

func main() {
	// Variables
	// string
	// bool
	// int
	// uint only positive numbers
	// float32 float 64
	// complex64

	// Using var
	// var name string = "Diego"

	var lastName = "Beltran"
	var age int = 23
	const isCool = true
	// shorthand only works in funcs
	// name := "Diego"
	var size float32 = 2.3
	// email := "diego@gmail.com"
	name, email := "diego", "diego@gmail.com"

	fmt.Println(name, lastName, age, isCool, size)
	fmt.Println(name, email)
	fmt.Printf("%T\n", name)
	fmt.Printf("%T\n", size)
	fmt.Printf("%T\n", email)
}
