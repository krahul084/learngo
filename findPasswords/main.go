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
	c := make(chan string)
	defer close(c)
	for _, file := range files {
		go func(file string) {
			searchString(text, file, c)
		}(file)
	}
	for i := 0; i < len(files); i++ {
		<-c
	}
}

func searchString(t string, f string, c chan string) {
	b, err := ioutil.ReadFile(f)
	if err != nil {
		panic(err)
	}
	s := string(b)
	if strings.Contains(s, t) {
		c <- f
		fmt.Printf("%s --> %s\n", t, f)
	}
	c <- f
}
