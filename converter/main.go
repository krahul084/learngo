package main

import (
	"fmt"
	"os"
	"strconv"
	"unicode/utf8"
)

func main() {
	arg := os.Args[1]
	heatCelcius, _ := strconv.ParseFloat(arg, 64)
	heatFarenheit := (heatCelcius*(9/5.) + 32)
	name := "Hello"
	fmt.Printf("Celcius value %g in Farenheit is %g\n", heatCelcius, heatFarenheit)
	fmt.Printf("Length of name is %d\n", utf8.RuneCountInString(name))
}
