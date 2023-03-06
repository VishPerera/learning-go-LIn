// Storing unordered values in maps
// map = unordered collection of key value pairs. a.k.a. hash table

package main

import (
	"fmt"
	"sort"
)

func main() {

	// Declaration of a map
	states := make(map[string]string)
	fmt.Println(states)

	// adding key value pairs to the map
	states["WA"] = "Washington"
	states["OR"] = "Origon"
	states["CA"] = "Califonia"
	fmt.Println(states)

	// Retrieve value from a map
	califonia := states["CA"]
	fmt.Println(califonia)

	delete(states, "OR")      // Deleting an item in a map
	states["NY"] = "New York" // Add item to map
	fmt.Println(states)

	// looping through a map to output values
	for k, v := range states {
		fmt.Printf("%v: %v\n", k, v)
	}

	// Sorting values in a map alphabatically
	// 1. extract the keys in the map as a slice
	// 2. sort the slice
	// 3. output value for each element of sorted slice

	keys := make([]string, len(states)) // initialize a slice
	i := 0
	for k := range states { // extract keys from map to keys slice
		keys[i] = k
		i++
	}
	fmt.Println(keys)
	sort.Strings(keys) // sort the slice
	fmt.Println(keys)

	for i := range keys { // iterative loop within indexed numbers
		fmt.Println(states[keys[i]])
	}

}
