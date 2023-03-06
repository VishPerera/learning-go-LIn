// Managing ordered values in Slices
// Slice is an abstraction layer that sits on top of an Array
// Slices can be re-sized and sorted whereas arrays cannot
package main

import (
	"fmt"
	"sort"
)

func main() {
	// Declaration of an Array with size 3
	var colors = [3]string{"Red", "Green", "Blue"}
	fmt.Println(colors)

	// Declaration of a Slice. Syntatical difference: no size within []
	var colors2 = []string{"Red", "Green", "Blue"}
	fmt.Println(colors)

	colors2 = append(colors2, "Purple") // Add another entry to the array
	fmt.Println(colors2)

	// Removing or selecting a range with append function
	colors2 = append(colors2[1:len(colors2)]) // Removes the first item
	fmt.Println(colors2)

	colors2 = append(colors2[:len(colors2)-1]) // Removes the last item
	fmt.Println(colors2)

	// Initialize a Slice with type and size with make()
	// make() takes 3 arguments. 1) type, 2) size, 3) optional cap limit

	numbers := make([]int, 5)
	numbers[0] = 154
	numbers[1] = 346
	numbers[2] = 12
	numbers[3] = 48
	numbers[4] = 7
	fmt.Println(numbers)

	numbers = append(numbers, 200)
	fmt.Println(numbers)

	// A slice can be sorted by using the sort package
	sort.Ints(numbers)
	fmt.Println(numbers)

}
