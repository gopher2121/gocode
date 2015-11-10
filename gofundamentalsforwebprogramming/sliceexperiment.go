package main

import (
	"fmt"
)

var array = [10]byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j'}

func main() {
	var aslice, bslice []byte

	aslice = array[3:7] // defg
	fmt.Println(string(aslice))
	fmt.Println("lenght and capacity", len(aslice), cap(aslice))

	bslice = aslice[0:5] // defg
	fmt.Println(string(bslice))
	//bslice 5 10
	fmt.Println(len(aslice), len(bslice), cap(aslice), cap(bslice))
}
