package main

import (
	"fmt"
	"log"

	"clalarco.io/greetings"
	"clalarco.io/names"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	greet, _ := greet("Mike")
	fmt.Println(greet)

	namesToGreet, _ := names.GetNames(4)

	greets, _ := greetAll(namesToGreet)
	printGreets(greets)

	// Get a greeting message and print it.
	fmt.Println(greetings.GetQuote())

}

func greet(name string) (string, error) {
	messages, err := greetAll([]string{name})
	return messages[name], err
}

func greetAll(names []string) (map[string]string, error) {
	// A map to associate names with messages.
	messages := make(map[string]string)
	for _, name := range names {
		message, err := greetings.Hello(name)
		if err != nil {
			return nil, err
		}
		messages[name] = message
	}
	return messages, nil
}

func printGreets(m map[string]string) {
	for _, message := range m {
		fmt.Println(message)
	}
}
