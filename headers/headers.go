package main

import (
	"bytes"
	"encoding/base64"
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

	buffer := &bytes.Buffer{}
	enc := base64.NewEncoder(base64.URLEncoding, buffer)
	enc.Write([]byte("user:passw0rd"))
	enc.Close()
	encodedCreds, err := buffer.ReadString('\n')
	if err != nil && err.Error() != "EOF" {
		log.Fatalln("Not able to read the creds")
	}
	req.Header.Add("Authorization", fmt.Sprintf("Basic %s", encodedCreds))
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
