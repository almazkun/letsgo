package main

import (
	"fmt"
	"log"
	
	"akun.dev/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	names := []string{"Yelena", "Anna", "Madina"}
	messages, err := greetings.Hellos(names)
	
	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Println(messages)
}
