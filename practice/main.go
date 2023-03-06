// Conditional Logic
// if statement does not need parentheses () for the conditonal statement
// Also have the option to initialize variables in the codiditional statement

package main

import (
	"fmt"
)

func main() {
	theAnswer := 45
	var result string

	// No need of parentheses for the conditional statement
	if theAnswer < 0 {
		result = "Less Than Zero"
	} else if theAnswer == 0 {
		result = "Equal to Zero"
	} else {
		result = "Greater Than Zero"
	}
	fmt.Println(result)

	// Variables can be initialized in the conditional statement
	if x := -36; x < 0 {
		result = "Less Than Zero"
	} else if x == 0 {
		result = "Equal to Zero"
	} else {
		result = "Greater Than Zero"
	}
	fmt.Println(result)

}
