package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	fmt.Println("Welcome to the calculator App")

	reader := bufio.NewReader(os.Stdin) // establish a new read from stdin
	fmt.Printf("Enter a number for Value 1: ")
	inNum1, _ := reader.ReadString('\n') // read until a newline is found
	// convert read string to a float
	value1, err := strconv.ParseFloat(strings.TrimSpace(inNum1), 64)
	if err != nil {
		panic("You did not Enter a number. Exiting Program")
	} else {
		fmt.Println("You Entered: ", value1)
	}

	fmt.Printf("Enter a number for Value 2: ")
	inNum2, _ := reader.ReadString('\n') // read until a newline is found
	value2, err := strconv.ParseFloat(strings.TrimSpace(inNum2), 64)
	if err != nil {
		panic("You did not Enter a number. Exiting Program")
	} else {
		fmt.Println("You Entered: ", value2)
	}

	// sum the values and format the output to 2 decimal places
	sum := value1 + value2
	fmt.Printf("The sum of the values %v and %v is: %.2f", value1, value2, sum)

}
