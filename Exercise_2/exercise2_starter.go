package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	counts := make(map[string]int)

	for _, p := range os.Args[1:] {
		if err := scan(p, counts); err != nil {
			log.Fatal(err)
		}
	}

	for w, c := range counts {
		fmt.Printf("%s %d\n", w, c)
	}
}

func scan(p string, counts map[string]int) error {
	// TODO: Implement this function
}
