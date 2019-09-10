package main

import "fmt"

func greeting(name string) string {
	return "Hello " + name
}

func getSum(num1, num2 int) int {
	return num1 + num2
}

func main() {
	fmt.Println("hello World")
	fmt.Println(greeting("Diego"))
	fmt.Println(getSum(3, 5))
}
