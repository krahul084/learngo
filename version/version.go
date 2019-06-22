package version

import (
	"fmt"
	"runtime"
)

// Printversion prints the version value
func Printversion() {
	fmt.Println(runtime.Version())
}
