package main

import (
	"fmt"
	"strconv"
)

// Define person struct
type Person struct {
	firstName string
	lastName  string
	city      string
	gender    string
	age       int
	alive     bool
}

// Greeting method for Person (value reciever)
func (person Person) greet() string {
	return "Hello my name is " + person.firstName + " " + person.lastName + " and I am " + strconv.Itoa(person.age)
}

// hasBirthday method (pointer reciever)
func (person *Person) hasBirthday() {
	person.age++
}

// getMarried (pointer reciever)
func (person *Person) getMarried(spouseLastName string) {
	if person.gender == "m" {
		return
	} else {
		person.lastName = spouseLastName
	}
}

func main() {
	// Init person using struct
	person1 := Person{firstName: "Diego", lastName: "Beltran",
		city: "Winnipeg", gender: "m", age: 23, alive: true}

	// Alternative
	person2 := Person{"Alexa", "Martinez", "Winnipeg", "f", 23, true}

	fmt.Println(person1)
	fmt.Println(person2)
	fmt.Println(person1.firstName, person2.firstName)
	fmt.Println(person1.greet())
	fmt.Println(person2.greet())
	for i := 0; i <= 10; i++ {
		person1.hasBirthday()
	}
	fmt.Println(person1.greet())
	person1.getMarried(person2.lastName)
	person2.getMarried(person1.lastName)
	fmt.Println("I got Married ", person2)
}
