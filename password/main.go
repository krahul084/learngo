package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	usage = `
	Usage: go run <file.go> <username> <password>
	This script validates the username and password
	`
	errUser = "Invalid Username-"
	errPass = "Invalid Password-"
	credsOK = "Welcome! Credentials validated successfully"
	user    = "rkonidala"
	pass    = "hello123"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println(strings.TrimSpace(usage))
	}
	args := os.Args
	inputUser, inputPass := args[1], args[2]

	if user != inputUser {
		log.Fatalln(errUser, inputUser)
	} else if pass != inputPass {
		log.Fatalln(errPass, inputPass)
	} else {
		fmt.Println(credsOK)
	}
}
