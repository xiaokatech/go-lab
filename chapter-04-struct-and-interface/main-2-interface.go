package main

import "fmt"

type Greeter interface {
	Greet() string
}

type Employee struct {
	Name    string
	Company string
}

func (e Employee) Greet() string {
	return fmt.Sprintf("Hi, I'm %s from %s", e.Name, e.Company)
}

type Dog struct {
	Name string
}

func (d Dog) Greet() string {
	return fmt.Sprintf("Woof! I'm %s", d.Name)
}

func greetAll(greeters []Greeter) {
	for _, g := range greeters {
		fmt.Println(g.Greet())
	}
}

func main() {
	people := []Greeter{
		Employee{"Alice", "Tech Co"},
		Dog{"Buddy"},
	}

	greetAll(people)
}
