package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	const Usage = `
	Usage: <script> <Enter mins>
    This Script converts minutes to seconds
	`
	if len(os.Args) < 2 {
		fmt.Println(strings.TrimSpace(Usage))
		return
	}
	mins, _ := strconv.Atoi(os.Args[1])
	secs := mins * 60
	fmt.Printf("%d Mins in seconds is %d\n", mins, secs)
}
