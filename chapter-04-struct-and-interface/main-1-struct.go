package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

// Define method for struct
func (p Person) Greet() string {
	return fmt.Sprintf("Hello, my name is %s and I'm %d years old", p.Name, p.Age)
}

// define method with input of pointer (Can change struct)
func (p *Person) Birthday() {
	p.Age++
}

func main_1() {
	// init method 1
	p1 := Person{"Alice", 25}

	// init method 2
	p2 := Person{
		Name: "Bob",
		Age:  30,
	}

	fmt.Println(p1)

	// Read field
	fmt.Println(p1.Name)

	// Modify field
	p2.Age = 31

	p := Person{"Charlie", 40}
	fmt.Println(p.Greet())
	p.Birthday() // modify struct
	fmt.Println(p.Greet())
}
