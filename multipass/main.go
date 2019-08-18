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
	errPass = "Invalid Password for "
	credsOK = "Welcome! Credentials validated successfully"
)

func main() {
	users := []string{"rkonidala", "krahul"}
	passwds := []string{"hello123", "hello"}
	if len(os.Args) != 3 {
		fmt.Println(strings.TrimSpace(usage))
	}

	args := os.Args
	inputUser, inputPass := args[1], args[2]

	userExists, index := sliceContains(users, inputUser)
	if !userExists {
		log.Fatalln(errUser, inputUser)
	} else if passwds[index] != inputPass {
		log.Fatalln(errPass, inputUser)
	} else {
		fmt.Println(credsOK)
	}
}

func sliceContains(s []string, t string) (bool, int) {
	for i, v := range s {
		if v == t {
			return true, i
		}
	}
	return false, 0
}
