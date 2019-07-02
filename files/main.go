package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	var files []string

	root := os.Args[2]
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if fi, err := os.Stat(path); err == nil {
			if fi.Mode().IsRegular() {
				files = append(files, path)
			}
		}

		return nil
	})
	if err != nil {
		panic(err)
	}
	text := os.Args[1]
	for _, file := range files {
		st, op := searchString(text, file)
		if st == true {
			fmt.Printf("%s--> %s\n", text, op)
		}

	}

}

func searchString(t string, f string) (bool, string) {
	b, err := ioutil.ReadFile(f)
	if err != nil {
		panic(err)
	}
	s := string(b)
	if strings.Contains(s, t) {
		return true, f
	}
	return false, f
}
