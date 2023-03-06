// Define fuctions as methods of custom types
// we can attach functions to a custom type. These functions are called methods
// method is a member of type. In other OO languages like Java, these
// fuctions are called member of a class

package main

import (
	"fmt"
)

func main() {
	poodle := Dog{"Poodle", 10, "Woof!"}
	fmt.Println(poodle)
	fmt.Printf("%+v\n", poodle)
	fmt.Printf("Breed: %v\nWeight: %v\n", poodle.Breed, poodle.Weight)

	// accessing 'Speak()' member of type 'Dog'
	poodle.Speak()

	// Change fields of exported object and call again
	poodle.Sound = "Arf!"
	poodle.Speak()

	// calling the func SpeakThreeTimes() creates a copy of the dog object
	// therefore the orignal dog struct variables are not changed
	// therefore the dog barks only 3 times.
	poodle.SpeakThreeTimes()
	poodle.SpeakThreeTimes()
}

// Dog is a struct
type Dog struct {
	Breed  string
	Weight int
	Sound  string
}

// Reciever -> (d Dog). passing a reference to a Dog object
// d -> identifier, Dog -> type
// Speak is how the dog barks
func (d Dog) Speak() {
	// reference the dog object with the identifier 'd'
	fmt.Println(d.Sound)
}

// SpeakThreeTimes is how loud the dog barks
func (d Dog) SpeakThreeTimes() {
	// fmt.Sprintf -> formats acording to the format specifiers and return the string
	d.Sound = fmt.Sprintf("%v %v %v", d.Sound, d.Sound, d.Sound)
	fmt.Println(d.Sound)
}
