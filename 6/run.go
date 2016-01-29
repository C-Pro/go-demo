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
	num_goroutines := 1
	//alter this to change amount of data processed
	//be careful 1000 * 1000 * 100 of int64 is already 762Mb
	data_size := 1000 * 1000 * 100
	chunk_size := data_size / num_goroutines

	a := make([]int64, data_size)
	for i := 0; i < len(a); i++ {
		a[i] = 1
	}
	c := make(chan int64)

	start := time.Now()

	for i := 0; i < num_goroutines; i++ {
		if i < num_goroutines-1 {
			go sum(&a, i*chunk_size, i*chunk_size+chunk_size, c)
		} else {
			//Last goroutine takes chunk_size + datasize%num_goroutines
			go sum(&a, i*chunk_size, data_size, c)
		}

	}

	var s int64
	for i := 0; i < num_goroutines; i++ {
		//wait for all goroutines to put values in the channel
		s += <-c
	}
	fmt.Printf("Sum is %d\n", s)
	fmt.Printf("Time elapsed: %f sec\n", time.Now().Sub(start).Seconds())
}
