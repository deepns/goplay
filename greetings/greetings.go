package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// init() is a special function like main()
// init() is called before the globals are initialized
func init() {
	rand.Seed(time.Now().UnixNano())
}

// Returns a random format string from a collection
func randomFormat() string {
	// A slice of strings
	// Note that even the last entry in the slice ends with a ','
	formatMessages := []string{
		"Hi, %v, Good morning",
		"Howdy, %v",
		"How are you, %v",
	}

	return formatMessages[rand.Intn(len(formatMessages))]
}

// Hello returns a greeting message with the given name
func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("Empty name")
	}
	return fmt.Sprintf(randomFormat(), name), nil
}

// Hellos takes a list of names and returns a greeting message for each of them
func Hellos(names []string) (map[string]string, error) {
	if len(names) == 0 {
		return nil, errors.New("Names are empty")
	}

	// make a map[key-type]value-type
	messages := make(map[string]string)

	for index, name := range names {
		if name == "" {
			return nil, fmt.Errorf("names[%v] is empty", index)
		}

		message := fmt.Sprintf(randomFormat(), name)
		messages[name] = message
	}

	return messages, nil
}
