// Grouping realted values using Structs
// Structs are data structures
// similar to C's structs
// Struct is a custom data type
package main

import (
	"fmt"
)

func main() {
	poodle := Dog{"Poodle", 10}
	fmt.Println(poodle)
	fmt.Printf("%+v\n", poodle)                                        //	%+v gives more information for debugging
	fmt.Printf("Breed: %v\nWeight: %v\n", poodle.Breed, poodle.Weight) // accessing values using . operator
	poodle.Weight = 8                                                  // changing individual values
	fmt.Printf("Breed: %v\nWeight: %v\n", poodle.Breed, poodle.Weight)
}

// defining a structure
type Dog struct {
	Breed  string
	Weight int
}
