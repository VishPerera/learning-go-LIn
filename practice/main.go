// Getting user input from the Console

package main

import (
	"bufio" // for Input Output purposes
	"fmt"
	"os" // platform-independent interface to operating system for IO
)

func main() {

	// variable to store the user input
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Text: ") // requst user to input text

	// pasue the application and let user type in
	// variable input stores whatever user types in
	// to ignore a variable name it with _
	input, _ := reader.ReadString('\n')
	fmt.Println("You entered: ", input)
}
