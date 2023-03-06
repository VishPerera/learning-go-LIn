// Define and Call Functions

package main

import (
	"fmt"
)

func main() {
	doSomething()
	fmt.Println("\nAdding predefined number of values: 2 values")
	sum := addValues(5, 7)
	fmt.Println("Sum: ", sum)

	fmt.Println("\nAdding arbitrary number of values")
	multiSum := addMany(5, 8, 6, 4)
	fmt.Println("Multi Sum: ", multiSum)

	fmt.Println("\nReturn multiple function values")
	multiSum2, numberCount := addMany2(5, 8, 6, 4)
	fmt.Println("Multi Sum: ", multiSum2)
	fmt.Println("Number of numbers: ", numberCount)
}

// if the function does not return anything no need to put type between () and {}
func doSomething() {
	fmt.Println("Doing Something")
}

func addValues(value1, value2 int) int { // if all the values are of same type, use type once
	return value1 + value2
}

// define arbitary number of arguments by '...' followed by type. Should be same type
func addMany(values ...int) int {
	total := 0
	for _, v := range values {
		total += v
	}
	return total
}

// return multiple values
// wrap comma delimited return types list in parentheses
// ex: (int, int, float)
func addMany2(values ...int) (int, int) {
	total := 0
	for _, v := range values {
		total += v
	}
	return total, len(values)
}
