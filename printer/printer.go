package printer

import (
	"fmt"
	"runtime"
)

// Hello is an exported function
func Hello() {
	fmt.Println("hello package")
	fmt.Println(runtime.Version())
}
