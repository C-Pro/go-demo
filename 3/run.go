package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func twoResults(x, y int) (a, b int) {
	a = x + 1
	b = y + 2
	return
}

const filename = "nonexistent.file"

func main() {
	x, y := twoResults(1, 2)
	fmt.Printf("x=%d y=%d\n", x, y)

	buf, err := ioutil.ReadFile(filename)

	if err != nil {
		log.Fatalf("Error reading file %s: %v", filename, err)
	}
	fmt.Printf("Read %d bytes", len(buf))

}
