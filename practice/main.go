// Use of Math package is illustrated here

package main

import (
	"fmt"
	"math"
)

func main() {

	// Sum of Integers
	i1, i2, i3 := 12, 45, 68
	intSum := i1 + i2 + i3
	fmt.Println("Integer Sum:", intSum)

	// Sum of floats
	f1, f2, f3 := 23.5, 65.1, 76.3
	floatSum := f1 + f2 + f3
	fmt.Println("Float Sum:", floatSum) // there is a precision issue with this summation

	/*
		floatSum = math.Round(floatSum) // round to nearest integer
		fmt.Println("The Sum is now:", floatSum)
	*/

	// round to 2 decimal places
	floatSum = math.Round(floatSum*100) / 100
	fmt.Println("The Sum is now:", floatSum)

	// type conversion for summing integers and floats
	sum := float64(i1) + f2
	fmt.Printf("The Sum is %v and type is %T\n", sum, sum)

	// rounding numbers with Printf formatting
	// rounded to 2 decimal points with %.2f
	circleRadius := 15.5
	circumference := 2 * math.Pi * circleRadius
	fmt.Printf("The circumference is %.2f ", circumference)

}
