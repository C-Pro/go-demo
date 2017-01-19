//This example shows how to create packages
//export and use their contents
//and do some unit testing
package main

//Go promotes open source
//Package names are urls, by which you can download the package
//To download and install all referenced packages run "go get"
import (
	"fmt"

	"github.com/c-pro/go-demo/2/mypack"
)

func main() {
	fmt.Printf("Hello, %s!\n", mypack.GetName())
}
