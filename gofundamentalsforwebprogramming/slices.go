// slices in golang is very useful

package main

import (
	"fmt"
)

var globalarray [10]int

func main() {
	globalarray[0] = 2
	globalarray[1] = 3
	globalarray[2] = 4
	globalarray[3] = 5
	globalarray[4] = 6
	globalarray[5] = 7
	globalarray[6] = 8
	globalarray[7] = 9
	globalarray[8] = 10
	globalarray[9] = 11

	fmt.Println(globalarray)

	slice := globalarray[1:3]
	// should print 3,4
	fmt.Println(slice)

	slicedouble := globalarray[:]
	// should displaye the same result as that of global array
	fmt.Println(slicedouble)

	fmt.Println("lenght  and capcity of slice : ", len(slice), cap(slice))

	slicedouble = append(slicedouble, 1, 3, 4)
	fmt.Println(slicedouble, len(slicedouble), cap(slicedouble))

}
