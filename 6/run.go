//This example shows usage of goroutines and channels
//It allocates slice of data_size int64 values
//sets them all to 1
//and runs num_goroutines concurrent goroutines to sum all
//values
package main

import (
	"fmt"
	"time"
)

//Sums part of slice
//Slice is passed buy reference
func sum(a *[]int64, f, t int, c chan int64) {
	var s int64
	for i := f; i < t; i++ {
		s += (*a)[i]
	}
	//push sum into our channel
	c <- s
}

func main() {

	//alter this value to set number of concurrent goroutines
	numGoroutines := 10
	//alter this to change amount of data processed
	//be careful 1000 * 1000 * 100 of int64 is already 762Mb
	dataSize := 1000 * 1000 * 100
	chunkSize := dataSize / numGoroutines

	a := make([]int64, dataSize)
	for i := 0; i < len(a); i++ {
		a[i] = 1
	}
	c := make(chan int64)

	start := time.Now()

	for i := 0; i < numGoroutines; i++ {
		if i < numGoroutines-1 {
			go sum(&a, i*chunkSize, i*chunkSize+chunkSize, c)
		} else {
			//Last goroutine takes chunkSize + dataSize%numGoroutines
			go sum(&a, i*chunkSize, dataSize, c)
		}

	}

	var s int64
	for i := 0; i < numGoroutines; i++ {
		//wait for all goroutines to put values in the channel
		s += <-c
	}
	fmt.Printf("Sum is %d\n", s)
	fmt.Printf("Time elapsed: %f sec\n", time.Now().Sub(start).Seconds())
}
