package main

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

func scan(filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	for s.Scan() {
		line := s.Text()
		words := strings.Fields(line)

		for i, w := range words {
			if w == "the" {
				fmt.Printf("%s is the %dth word\n", w, i)
			}
		}
	}
	return s.Err()
}
