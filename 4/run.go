//This example shows arrays, slices, maps, defer, interface{} and reflection
package main

import (
	"fmt"
	"reflect"
)

var (
	myarray [3]int
	myslice []string
	mymap   map[string]string
)

//Empty interface can hold any type
//and can be used to create somewhat "generic" functions
func printit(a interface{}) {
	defer fmt.Println("===END===")
	fmt.Println("==BEGIN==")
	val := reflect.ValueOf(a)

	switch reflect.TypeOf(a).Kind() {
	case reflect.Array:
		for i := 0; i < val.Len(); i++ {
			fmt.Printf("a[%d] = %d\n", i, val.Index(i).Int())
		}
	case reflect.Slice:
		for i := 0; i < val.Len(); i++ {
			fmt.Printf("a[%d] = %s\n", i, val.Index(i).String())
		}
	case reflect.Map:
		for _, v := range val.MapKeys() {
			fmt.Printf("a[%s] = %s\n", v.String(), val.MapIndex(v).String())
		}
	default:
		fmt.Println("Unexpected type!")
		return
	}
}

func main() {
	printit(myarray)
	myslice = make([]string, 5)
	for i := 0; i < len(myslice); i++ {
		switch i % 2 {
		case 0:
			myslice[i] = "Even"
		case 1:
			myslice[i] = "Odd"
		}
	}
	printit(myslice)
	myslice = append(myslice, "Odd again!")
	printit(myslice)

	mymap = make(map[string]string, 2)
	mymap["test"] = "value"
	mymap["foo"] = "bar"
	printit(mymap)
	printit("ERROR HERE!")
}
