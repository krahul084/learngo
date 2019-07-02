package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	client := http.DefaultClient
	req, err := http.NewRequest("GET", "https://httpbin.org/basic-auth/user/passw0rd", nil)
	if err != nil {
		log.Fatalln("Unable to create a request")
	}

	req.SetBasicAuth("user", "passw0rd")
	resp, err := client.Do(req)

	if err != nil {
		log.Fatalln("Unable to get the response")
	}
	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln("Unable to read content")
	}
	fmt.Println(string(content))
	fmt.Println(resp.StatusCode)
}
