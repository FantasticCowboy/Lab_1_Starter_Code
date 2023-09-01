package main

// Helpful!
//	os.Open(name string) (*File, error): Open a file
//	bufio.NewScanner(r io.Reader): Scanner from a file
//	string.Fields(s string) []string: Split a string into words

import (
	"log"
	"os"
)

func main() {
	for _, p := range os.Args[1:] {
		if err := scan(p); err != nil {
			log.Fatal(err)
		}
	}
}

func scan(p string) error {
	// TODO: Implement
	return nil
}
