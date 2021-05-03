package main

import (
	"fmt"
	"log"

	greetings "github.com/Moranilt/golang-greet"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	names := []string{"Danil", "Sam", "Jon"}

	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Println(message)
	for _, message := range messages {
		fmt.Println(message)
	}
}
