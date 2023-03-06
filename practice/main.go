// Loops with for statement
// There is no 'whhile' statement in GO
// 'While' loop is achieved by extended version of 'for' syntax
// 'break' and 'continue' statements have their usual meaning as in C, Java

package main

import (
	"fmt"
)

func main() {
	colors := []string{"Red", "Green", "Blue"}
	fmt.Println(colors)

	fmt.Println("\nThree-part for loop:")
	// for loop with three-part declaration
	for i := 0; i < len(colors); i++ {
		fmt.Println(colors[i])
	}

	fmt.Println("\nfor loop using range keyword:")
	// use range keyword for more concise writing
	// works for arrays
	for i := range colors { // set the index position to i each time loop goes through colors array
		fmt.Println(colors[i])
	}

	fmt.Println("\nLoop using foreach loop:")
	// foreach implementation returns two values as a comma delimited list
	// firt returned variable is the index and the second is the value
	// if only the value is needed ignore the first by _
	for _, clr := range colors {
		fmt.Println(clr) // returned 'value' from the foreach loop
	}

	fmt.Println("\nWhile Loop implementation by for loop:")
	// to implement 'while loop' declare a boolean condition
	// and use as the expression in the for loop
	value := 1
	for value < 10 {
		fmt.Println("Value: ", value)
		value++
	}

	fmt.Println("\nUse of 'GoTo' labels in 'for' loop:")
	// breakout from the loop if the sum > 200 using goto label
	sum := 1
	for sum < 1000 {
		sum += sum
		if sum > 200 { // if the sum > 200, break and goto the lable 'theEnd'
			goto theEnd
		}
		fmt.Println("Sum: ", sum)
	}

	// defining a 'goto' label in the format
	// label:
theEnd:
	fmt.Println("End of the program")

}
