// Reference values with pointers

package main

import (
	"fmt"
)

func main() {
	anInt := 45
	var p = &anInt                  // stores the address of anInt variable in p
	fmt.Println("Address of p:", p) // Address stored in p
	fmt.Println("Value of p:", *p)  // Value stored at the address reffered by p

	value1 := 42.6
	pointer1 := &value1
	fmt.Println("Value 1", *pointer1)

	// Original variable value can be changed by
	// manipulating the dereferenced pointer
	*pointer1 = *pointer1 / 24
	fmt.Println("Pointer 1", *pointer1)
	fmt.Println("Value 1", value1)

}
