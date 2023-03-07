// Read a text file from the web

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "http://services.explorecalifornia.org/json/tours.php"

func main() {

	resp, err := http.Get(url) // download the JSON content into my machine
	if err != nil {
		panic(err)
	}

	fmt.Printf("Response Type %T\n", resp)

	// The Body field is public and can be read
	defer resp.Body.Close() // close the body after everything is done (in the following lines)

	bytes, err := ioutil.ReadAll(resp.Body) // A byte array is received when read
	if err != nil {
		panic(err)
	}

	content := string(bytes)
	fmt.Print(content)
}
