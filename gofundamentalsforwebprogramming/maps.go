//maps provide an efficient way like a dictionary

package main

import (
	"fmt"
)

func main() {
	mapv := make(map[string]string)
	mapv["dhiraj"] = "dhiraj"
	mapv["roshan"] = "roshan"
	mapv["golang"] = "golang"

	fmt.Println(mapv)

	// use the range loop to iterate over map
	for k, v := range mapv {
		fmt.Println("key: ", k)
		fmt.Println("value: ", v)
	}

	// delete a ket
	delete(mapv, "dhiraj")
	for k, v := range mapv {
		fmt.Println("key: ", k)
		fmt.Println("value: ", v)
	}

}
