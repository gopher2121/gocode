// sprinkle will add some words so that it is easy to find the available domains

package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

const (
	otherWords = "*"
)

var transforms = []string{
	otherWords,
	otherWords,
	otherWords + "app",
	otherWords + "site",
	otherWords + "time",
	"lets" + otherWords,
	"get" + otherWords,
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	s := bufio.NewScanner(os.Stdin)

	for s.Scan() {
		t := transforms[rand.Intn(len(transforms))]
		fmt.Println(strings.Replace(t, otherWords, s.Text(), -1))
	}

}
