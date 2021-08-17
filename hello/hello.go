package main

import (
	"fmt"
	"log"
	
	"akun.dev/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	message, err := greetings.Hello("Yelena")

	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Println(message)
}
