package main

import (
	"fmt"
	"os"
)

func main() {
	name := os.Args[1]
	age := os.Args[2]
	fmt.Println("Hello", name, "seems your age is", age)
	fmt.Println("number of arguments passed -", len(os.Args)-1)
}
