// Read and Write local text files

package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	content := "Hello from Go"                 // string to be written to the file
	file, err := os.Create("./fromString.txt") // create a file in the same directory
	checkErrors(err)
	length, err := io.WriteString(file, content)
	checkErrors(err)
	fmt.Printf("Wrote a file with %v characters\n", length)
	defer file.Close()                 // Defer waits until everything is finished to close the file
	defer readFile("./fromString.txt") // Wait until everything is done to read the file

}

// Check for errors. If error exists, show the error and exit the program
func checkErrors(err error) {
	if err != nil {
		panic(err)
	}
}

// Read from file
// data is read from the file in bytes
// Therefore, read data should be made to a string to display: string(data)
func readFile(fileName string) {
	data, err := ioutil.ReadFile(fileName)
	checkErrors(err)
	fmt.Println("Text read from file: ", string(data))
}
