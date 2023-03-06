// Advanced Calculator App

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin) // establish a new read from stdin

func main() {

	// Prompt user for inputs
	fmt.Println("Welcome to the calculator App")
	num1 := numRead("Enter Value 1: ")
	num2 := numRead("Enter Value 2: ")
	operator := selectOperator("Select one from (+ - / *): ")

	var result float64

	// Switch to correct math operation
	switch operator {
	case "+":
		result = addValues(num1, num2)
	case "-":
		result = subtractValues(num1, num2)
	case "/":
		result = divideValues(num1, num2)
	case "*":
		result = multiplyValues(num1, num2)
	default:
		panic("Invalid Operation")
	}

	//result = math.Round(result*100) / 100
	fmt.Println("The Result Is: ", result)

}

// Read user input number from the keyboard
// input parameters: string to indicate which operand
// return: number value
// validation: validate for number. If not, panic and exit program
func numRead(msg string) float64 {
	fmt.Printf("%v", msg)
	inNum1, _ := reader.ReadString('\n') // read until a newline is found
	// convert read string to a float
	value, err := strconv.ParseFloat(strings.TrimSpace(inNum1), 64)
	if err != nil {
		message := fmt.Sprintf("\"%v\" must be a number", msg)
		panic(message)
	}
	return value
}

// Read a mathematical operator from the keyboard
// input parameters: string to input an operator
// return: operator as a byte
// validation: validate for a character. If not, panic and exit program
func selectOperator(msg string) string {
	fmt.Printf("%v", msg)
	operator, _ := reader.ReadString('\n') // read until a newline is found
	// convert read string to a character
	return strings.TrimSpace(operator)
}

// List of math operations
func addValues(num1, num2 float64) float64 {
	return num1 + num2
}
func subtractValues(num1, num2 float64) float64 {
	return num1 - num2
}
func divideValues(num1, num2 float64) float64 {
	return num1 / num2
}
func multiplyValues(num1, num2 float64) float64 {
	return num1 * num2
}
