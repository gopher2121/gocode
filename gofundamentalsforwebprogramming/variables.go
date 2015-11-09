// how to declare variables in golang

package main

import (
	"fmt"
)

// defining one variable
var l int

// defining multiple variables
var a, b, c int

// initializing one variable
var d int = 5

// initializing multiple variables
var m, n, v, k int = 2, 3, 4, 5

// one can omit the type if they like
var h = 4

// one can also group the variables
var (
	q int = 2
)

// one can even omit the var keyword , but it can be only inside a function
func main() {
	j := 2
	fmt.Println("l a b c d m n v k h j q\n", l, a, b, c, d, m, n, k, h, j, q)
}
