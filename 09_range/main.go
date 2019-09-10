package main

import "fmt"

func main() {
	fmt.Println("Hello World!")
	ids := []int{12, 312, 432, 13, 56, 86, 10}

	for i, id := range ids {
		fmt.Printf("%d - ID: %d\n", i, id)
	}

	// use an underscore if you arent using the variable
	for _, id := range ids {
		fmt.Printf("ID: %d\n", id)
	}

	// Adding all the ids together
	sum := 0
	for _, id := range ids {
		sum += id
	}
	fmt.Printf("Sum of IDs: %d\n", sum)

	// range with Maps
	names := map[string]string{"Zinc": "Fook", "Diego": "Beltran", "Alexa": "Martinez", "Zac": "Fack"}
	for k, v := range names {
		fmt.Printf("%s: %s\n", k, v)
	}

	for k := range names {
		fmt.Printf("First Name: %s\n", k)
	}

	for _, v := range names {
		fmt.Printf("Last Name: %s\n", v)
	}
}
