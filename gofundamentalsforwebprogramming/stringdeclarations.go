package main

import (
	"fmt"
)

func main() {
	s := "hello world"
	c := []byte(s)
	c[0] = 'G'
	s2 := string(c)
	fmt.Println(s2)
}
