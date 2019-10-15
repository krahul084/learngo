package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Your name")
		return
	}
	inpName := os.Args[1]
	moods := [4]string{"happy", "sad", "terrible", "excited"}
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(len(moods))
	fmt.Printf("%s feels %s\n", inpName, moods[n])
}
