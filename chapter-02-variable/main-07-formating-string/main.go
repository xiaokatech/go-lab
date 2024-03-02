package main

import "fmt"

func main() {
	fmt.Printf("I am %v years old\n", 10)
	// I am 10 years old
	fmt.Printf("I am %v years old\n", "way too many")
	// I am way to many years old

	fmt.Printf("I am %s years old\n", "way too many")
	// I am way to many years old

	fmt.Printf("I am %d years old\n", 10)
	// I am 10 years old

	fmt.Printf("I am %f years old\n", 10.523)
	// I am 10 years old
	fmt.Printf("I am %.2f years old\n", 10.523)
	// I am 10 years old

	const name = "Saul Goodman"
	const openRate = 30.5

	msg := fmt.Sprintf("Hi %s, your open rate is %.1f percent", name, openRate)

	fmt.Println(msg)
}
