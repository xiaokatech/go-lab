package main

import "fmt"

func main() {
	createUser("xiaoka", "", 10)
}

func createUser(firstName, lastName string, age int) {
	fmt.Printf("My name is %s %s, I'm %d years old\n", firstName, lastName, age)
	var returnOfSprintf string = fmt.Sprintf("My name is %s %s, I'm %d years old\n", firstName, lastName, age)
	fmt.Println(returnOfSprintf)
}
