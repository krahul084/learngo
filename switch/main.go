package main

import "fmt"

func main() {
	switch 1 {
	case 1, 5, 6:
		fmt.Println("Prints 1,5,6")
	case 2, 3, 4:
		fmt.Println("Prints 2,3,4")
	default:
		fmt.Println("Prints nothing")
	}
}
