// arrays in golang

package main

import (
	"fmt"
)

var array [5]int // array type includes the size aswell

func main() {
	array[0] = 2
	array[1] = 3
	fmt.Println(array[0])

	singleDarray := [5]int{2, 3, 4, 5, 6}
	fmt.Println(singleDarray)

	doubleDarray := [2][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}}
	fmt.Println(doubleDarray)

	// passing the singleDarray to the function total
	total := total(singleDarray)
	fmt.Println(total)
}

// pass the array... it will be copied as a value
func total(d [5]int) int {
	return d[0] + d[1] + d[2] + d[3] + d[4]

}
