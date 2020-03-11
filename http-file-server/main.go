package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// handle ctrl+c (signal interrupt) close in the terminal
	func() {
		c := make(chan os.Signal, 2)
		signal.Notify(c, os.Interrupt, syscall.SIGTERM)
		go func() {
			<-c
			fmt.Println("\rKeyboard interrupt received, exiting.")
			os.Exit(0)
		}()
	}()

	// handle http requests
	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("./"))))

	// handle http service
	fmt.Println("Serving HTTP on 0.0.0.0 port 8000 (http://0.0.0.0:8000/) ...")
	fmt.Println("Serving HTTP on localhost port 8000 (http://localhost:8000/) ...")
	log.Fatal(http.ListenAndServe(":8000", nil))

}
