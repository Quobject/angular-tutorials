package main

import "fmt"

//Animal ...
type Animal interface {
	Speak() string
}

//Dog ...
type Dog struct {
}

//Speak ...
func (d Dog) Speak() string {
	return "Woof!"
}

//Cat ...
type Cat struct {
}

//Speak ...
func (d Cat) Speak() string {
	return "Meow!"
}

//Cow ...
type Cow struct {
}

//Speak ...
func (d Cow) Speak() string {
	return "Moo!"
}

func main() {
	poodle := Animal(Dog{})
	fmt.Println(poodle)

	animals := []Animal{Dog{}, Cat{}, Cow{}}
	for _, animal := range animals {
		fmt.Println(animal.Speak())
	}
}
