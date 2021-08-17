package main

import (
	"fmt"
	
	"akun.dev/greetings"
)

func main() {
	message := greetings.Hello("Yelena")
	fmt.Println(message)
}
