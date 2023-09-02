package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type countsErr struct {
	counts map[string]int
	err    error
}

func main() {
	ccounts := make(chan countsErr)

	for _, p := range os.Args[1:] {
		go workerScan(p, ccounts)
	}

	totalCounts := make(map[string]int)
	for range os.Args[1:] {
		ce := <-ccounts
		if ce.err != nil {
			log.Fatal(ce.err)
		}
		for w, c := range ce.counts {
			totalCounts[w] += c
		}
	}

	for w, c := range totalCounts {
		fmt.Printf("%s %d\n", w, c)
	}
}

func workerScan(p string, ccounts chan countsErr) {
	cs := make(map[string]int)
	err := scan(p, cs)
	ccounts <- countsErr{cs, err}
}

func scan(p string, counts map[string]int) error {
	f, err := os.Open(p)
	if err != nil {
		return err
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	for s.Scan() {
		line := s.Text()
		words := strings.Fields(line)

		for _, w := range words {
			counts[w]++
		}
	}
	return s.Err()
}
