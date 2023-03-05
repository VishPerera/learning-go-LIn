package main

// Using a first capital letter in variable naming
// makes the variable public. simple first letter
// allows access within the declared file
import (
	"fmt"
)

const aConstant int = 67

func main() {
	// Declaring a variable
	var aString string = "This is Go"

	// Print result to terminal using Println
	fmt.Println(aString)

	// To use output formating use Printf
	fmt.Printf("The type of the variable is: %T\n", aString)

	var anInteger int = 42
	fmt.Println(anInteger)

	// Variables have default values: int default value = 0
	var defaultInt int
	fmt.Println(defaultInt)

	// Implicit data type infrred by the complier
	var anotherString = "This is another string"
	fmt.Println(anotherString)
	fmt.Printf("The type of the variable is: %T\n", anotherString)

	// If a decimal point is added the type becomes float64
	var anotherInt = 65 // implicit assignment of an integer
	fmt.Println(anotherInt)
	fmt.Printf("The type of the variable is: %T\n", anotherInt)

	// shortcut to declare implicit variables
	// := is only valid for variables inside a function
	myString := "This is myString"
	fmt.Println(myString)
	fmt.Printf("The type of the variable is: %T\n", myString)

	// Constant variable
	fmt.Println(aConstant)
	fmt.Printf("The type of the variable is: %T\n", aConstant)

}
