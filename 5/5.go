package main

import (
	"fmt"
	"math"
)

//Struct example
type Box struct {
	Width, Height, Depth int
}

//Struct embedding example
type StuffedBox struct {
	Box
	Contents string
}

//Struct method and pointer example
func (b *Box) Volume() float64 {
	//Notice type conversion  to float64 here
	//Go does not have implicit type conversions!
	return float64(b.Width*b.Height*b.Depth) / math.Pow(1000, 3)
}

func main() {

	//multiline var
	var (
		box Box
		//Pointer to struct
		sbox *StuffedBox
	)

	box.Width = 1000
	box.Height = 1000
	box.Depth = 1000
	fmt.Printf("Box volume is %f cubic meters\n", box.Volume())

	//create Stuffed box struct and assign
	//its address to sbox pointer
	sbox = &StuffedBox{
		Box: Box{
			Width:  1000,
			Height: 1000,
			Depth:  2000},
		Contents: "Teddy bears"}
	//Stuffed box embeds Box's fields and methods

	//We can access embedded struct fields directly, or via field name:
	fmt.Printf("%d = %d\n", (*sbox).Box.Width, (*sbox).Width)

	//struct and pointer to struct field and method access is done with dot
	fmt.Printf("StuffedBox volume is %f cubic meters\n", sbox.Volume())

}
