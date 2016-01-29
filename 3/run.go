//This example shows functions returning multiple values and error handling
package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

//This function returns two values at a time
//We could just pass them by reference
//but I wanted to show off this first :)
func twoResults(x, y int) (a, b int) {
	a = x + 1
	b = y + 2
	return
}

//Constant type is inferred from value
const filename = "nonexistent.file"

func main() {
	//multiple assignment syntax
	x, y := twoResults(1, 2)
	fmt.Printf("x=%d y=%d\n", x, y)

	buf, err := ioutil.ReadFile(filename)

	if err != nil {
		//Notice pretty datetime before log message
		log.Fatalf("Error reading file %s: %v", filename, err)
	}
	//you should fix error to run down to this line
	fmt.Printf("Read %d bytes", len(buf))
}
