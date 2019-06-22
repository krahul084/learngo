package main

import "fmt"

func main() {
	var (
		speed, tyreCount int
		heat             float64
		off              bool
		brand            string
	)
	_ = tyreCount
	fmt.Println(speed)
	fmt.Println(heat)
	fmt.Println(off)
	fmt.Printf("%q\n", brand)
	shortDecs()
	multiDecs()
}

func shortDecs() {
	safe := true
	place := "Burlington"
	fmt.Println(place, safe)
}

func multiDecs() {
	firstname, lastname := "Nicola", "Tesla"
	noFiles, off, owner := 5, true, "Rahul"
	fmt.Println(firstname, lastname)
	fmt.Println(noFiles, off, owner)
}
