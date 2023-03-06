// Using Switch statements
// break; command is not needed. As soon as the case is evaluated
// it jumps to the end of the switch staement

package main

import (
	"fmt"
	"math/rand" // to generate random value
	"time"
)

func main() {
	rand.Seed(time.Now().Unix()) // generate randamized number based on current time in UNIX format
	// dow := rand.Intn(7) + 1      // random number between 1 and 7
	// fmt.Println("Day", dow)

	var result string
	// do not need parentheses for the expression
	// a statement can be added before the evaluation of the switch expression
	switch dow := rand.Intn(7) + 1; dow {
	case 1:
		result = "It is Sunday"
		// fallthrough		// fallthrough keyword execute the next case always
	case 2:
		result = "It is Monday"
		// fallthrough
	default:
		result = "It is some other day"
	}
	fmt.Println(result)

}
