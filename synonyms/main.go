package main

import (
	"bufio"
	"fmt"
	"github.com/gocode/thesaurus"

	"log"
	"os"
)

func main() {

	apiKey := "9857ad7002ee71145d6852cab3229fff" //os.Getenv("BHI_APIKEY")

	thesaurus := &thesaurus.BigHuge{APIKey: apiKey}
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		word := s.Text()
		syns, err := thesaurus.Synonyms(word)
		if err != nil {
			log.Fatalln("failed when looking for synonyms")
		}
		if len(syns) == 0 {
			log.Fatalln("couldnot find any synonyms")
		}
		for _, syn := range syns {
			fmt.Println(syn)
		}
	}

}
