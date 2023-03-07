// Parse JSON formatted Text

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
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

	content := string(bytes) // forming a string from bytes
	// fmt.Print(content)

	// Output tour names found in the downloaded file 'content'
	tours := toursFromJson(content)
	for _, tour := range tours {
		fmt.Println(tour.Name)
	}
}

// Function to extract tour information in the downloaded file from the web
func toursFromJson(content string) []Tour {
	tours := make([]Tour, 0, 20)

	// decode the information in the downloaded file: json.NewDecoder(readerObject)
	// readerObject: new reader from strings.NewReader(content)
	decoder := json.NewDecoder(strings.NewReader(content))

	// to check errors look a the Token of the decoder
	_, err := decoder.Token()
	if err != nil {
		panic(err)
	}

	var tour Tour
	// read the next available object JSON content
	// if finds content return True, else False
	for decoder.More() {
		err := decoder.Decode(&tour)
		if err != nil {
			panic(err)
		}

		tours = append(tours, tour) // append the tours slice
	}
	return tours
}

type Tour struct {
	Name, Price string
}
