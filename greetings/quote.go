package greetings

import (
	"rsc.io/quote"
)

// Hello returns a greeting for the named person.
func GetQuote() string {
	// Return a greeting that embeds the name in a message.
	return quote.Glass()
}
