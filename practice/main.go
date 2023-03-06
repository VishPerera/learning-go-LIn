// Date and Time implementation

package main

import (
	"fmt"
	"time" // module package for date/time/timezone
)

func main() {

	n := time.Now() // Date/time/time zone right now
	fmt.Println("The time of program run:", n)

	t := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC) // user defined
	fmt.Println("Go launched on:", t)
	fmt.Println(t.Format(time.ANSIC)) // Format time objet as ANSIC standard

	// parsing time
	parsedTime, _ := time.Parse(time.ANSIC, "Tue Nov 10 23:00:00 2009")
	fmt.Printf("The type of parsedTime is %T\n", parsedTime)

}
