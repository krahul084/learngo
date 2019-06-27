package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct {
}

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error opening google", err)
		os.Exit(1)
	}
	lw := logWriter{}
	io.Copy(lw, resp.Body)
}

func (lw logWriter) Write(bs []byte) (n int, err error) {
	fmt.Println(string(bs))
	return len(bs), nil
}
