// Storing ordered values in Arrays

package main

import (
	"fmt"
)

func main() {
	// Defining an array.
	// name [lenght]type
	var colors [3]string // array named colors of size 3 of string type
	colors[0] = "Red"
	colors[1] = "Green"
	colors[2] = "Blue"
	fmt.Println(colors)
	fmt.Println(colors[0])

	// declare and initialize an array in single line
	var number = [5]int{2, 6, 7, 1, 8}
	fmt.Println(number)

	// length of an array can be obtained by len() function
	fmt.Println("Number of Colors:", len(colors))
	fmt.Println("Number of Number:", len(number))
}
