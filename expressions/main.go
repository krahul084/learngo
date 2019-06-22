// This is the program which creates an executable program
package main

import (
	"fmt"
	"runtime"
)

/*
This is the main entry point for the executable command
Executable programs are called as commands
*/
func main() {
	fmt.Println(runtime.NumCPU() + 1)
}
