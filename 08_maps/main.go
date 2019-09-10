package main

import "fmt"

func main() {
	emails := make(map[string]string)

	// Assign key value
	emails["Diego"] = "diego@gmail.com"
	emails["Karan"] = "karan@gmail.com"
	emails["Alexa"] = "alexa@gmail.com"
	emails["Mike"] = "mike@gmail.com"
	fmt.Println(emails)
	fmt.Println(emails["Alexa"])

	// Delete from maps
	delete(emails, "Karan")
	fmt.Println(emails)

	// Decalre map and add key value
	names := map[string]string{"Zinc": "Fook", "Diego": "Beltran", "Alexa": "Martinez", "Zac": "Fack"}
	fmt.Println(names)
}
