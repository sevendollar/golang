package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	jef, err := newUser("jef", "p")
	jef.Role = "admin"
	if err != nil {
		log.Fatal(err)
	}
	manager := newJWTManager("secret", time.Minute*30)
	token, err := manager.Generate(*jef)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(token)

	c, err := manager.Verify(token)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(c)
}
