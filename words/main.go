package main

import (
	"fmt"
	"os"
	"strings"
)

const corpus = "" + "a cat jump from wall again and again on the wall"

func main() {
	words := strings.Fields(corpus)
	query := os.Args[1:]

queries:
	for _, q := range query {
		for i, w := range words {
			if q == w {
				fmt.Printf("#%-2d: %q\n", i, w)
				break queries
			}
		}
	}
}
