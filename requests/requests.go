package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	client := http.DefaultClient
	req, err := http.NewRequest("GET", "https://httpbin.org/get", nil)
	if err != nil {
		log.Fatalln("Unable to create the request")
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln("Unable to get the page")
	}
	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Unale to read the content")
	}
	fmt.Println(string(content))
}
