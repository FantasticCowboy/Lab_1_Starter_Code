package main

// Helpful!
//	os.Open(name string) (*File, error): Open a file
//	bufio.NewScanner(r io.Reader): Scanner from a file
//	string.Fields(s string) []string: Split a string into words

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	for _, p := range os.Args[1:] {
		if err := scan(p); err != nil {
			log.Fatal(err)
		}
	}
}

func scan(p string) error {
	f, err := os.Open(p)
	if err != nil {
		panic("Oh no file doesn't exist!")
	}
	buf := bufio.NewScanner(f)
	for buf.Scan() {
		words := strings.Fields(buf.Text())
		for i, word := range words {
			fmt.Printf("Found %s at %d\n", word, i)
		}
	}
	return nil
}
