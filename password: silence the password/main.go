// silence the user inputed password
package main

import (
	"fmt"
	"log"
	"syscall"

	"golang.org/x/crypto/ssh/terminal"
)

func main() {
	//get the password from standard input
	fmt.Print("your password: ")
	newBytePass, err := terminal.ReadPassword(int(syscall.Stdin))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println()

	// print out the password
	fmt.Println(string(newBytePass))
}
