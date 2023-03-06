// Convert string input to another types

package main

import (
	"bufio" // for IO
	"fmt"
	"os"      // platform indepedent OS interface to IO
	"strconv" // convert strings to/from other types
	"strings" // functions for manipulating strings
)

func main() {

	// acquire and display user entred text string
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	input, _ := reader.ReadString('\n')
	fmt.Println("You entered:", input)

	// convert string into a number and validates whether it is a number
	fmt.Print("Enter a number: ")
	inputNum, _ := reader.ReadString('\n') // store user input in inputNum

	// TrimSpace: remove spaces before and after the string
	// convert the string to a float using strconv.ParseFloat
	// any errors are stored in err variable
	aFloat, err := strconv.ParseFloat(strings.TrimSpace(inputNum), 64)
	if err != nil { // If there is an error, show the error returned by ParseFloat
		fmt.Println("We encountered an Error: ", err)
	} else {
		fmt.Println("Value of Number You Entered: ", aFloat)
	}

}
