package main

import (
	"fmt"

	greetings "github.com/Moranilt/golang-greet"
)

func main() {
	message := greetings.Hello("Danil")

	fmt.Println(message)
}
