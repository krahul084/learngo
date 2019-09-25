package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	n := t.Hour()
	switch {
	case n >= 4 && n < 12:
		fmt.Println("Good Morning")
	case n >= 12 && n < 18:
		fmt.Println("Good Afternoon")
	case n >= 18 && n < 21:
		fmt.Println("Good Evening")
	case n >= 21 && n <= 23:
		fmt.Println("Good Night")
	default:
		fmt.Println("Its midnight")
	}
}
