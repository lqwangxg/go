package greetings

import (
	"errors"
	"fmt"
)

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name") // error.Error("empty name")
	}
	// Return a greeting that embeds the name in a message.
	message := fmt.Sprintf("Hello, %v. Welcome!", name)
	return message, nil
}

//Hi somebody. How do you do
func Hi(name string) string {
	var message string
	message = fmt.Sprintf("Hi, %v. How do you do!", name)
	return message

}
