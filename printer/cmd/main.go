package main

import (
	"fmt"

	"github.com/krahul084/learngo/printer"
)

func main() {
	planet := "Venus"
	distance := 225
	orbital := 224.7
	hasLife := false
	printer.Hello()
	fmt.Printf("Does %s have life with its orbital %f million kms away from sun?\n", planet, orbital)
	fmt.Printf("Yes %[2]v has life: %[1]v\n", hasLife, planet)
	fmt.Printf("%v is %v million kms away from us\n", planet, distance)
	fmt.Printf("Type of hasLife is %T\n", hasLife)
}
