package main

import "fmt"

//Dog ...
type Dog struct {
	Breed  string
	Weight int
	Sound  string
}

//Speak ...
func (d Dog) Speak() {
	fmt.Println(d.Sound)
}

//SpeakThreeTimes ...
func (d *Dog) SpeakThreeTimes() {
	d.Sound = fmt.Sprintf("%v %v %v", d.Sound, d.Sound, d.Sound)
	fmt.Println(d.Sound)
}

func main() {
	poodle := Dog{"Poodel", 37, "Woof"}
	poodle.Speak()
	fmt.Println(poodle)

	poodle.Sound = "Arf"
	poodle.SpeakThreeTimes()
	poodle.SpeakThreeTimes()

}
