package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type getResponse struct {
	Origin  string            `json:"origin"`
	URL     string            `json:"url"`
	Headers map[string]string `json:"headers"`
}

func main() {
	client := http.DefaultClient
	req, err := http.NewRequest("GET", "https://httpbin.org/get", nil)
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
	respContent := getResponse{}
	json.Unmarshal(content, &respContent)
	fmt.Println(respContent.toString())
}

func (r getResponse) toString() string {
	s := fmt.Sprintf("GET RESPONSE\nRequest URL: %s\nOrigin IP: %s", r.URL, r.Origin)
	for k, v := range r.Headers {
		s += fmt.Sprintf("Header: %s = %s\n", k, v)
	}
	return s
}
