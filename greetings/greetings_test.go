package greetings

import (
	"regexp"
	"testing"
)

// TestHelloEmpty calls Hello with an empty string
// and checks for the error
func TestHelloEmpty(t *testing.T) {
	name := ""
	message, err := Hello(name)
	if message != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, expected "", error`, message, err)
	}
}

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestHelloName(t *testing.T) {
	name := "Aron"
	want := regexp.MustCompile(`\b` + name + `\b`)
	messsage, err := Hello("Aron")

	// %q is single quoted character literal matching Go syntax
	if !want.MatchString(messsage) || err != nil {
		t.Fatalf(`Hello("Aron") = %q, %v, want match for %#q, nil`, messsage, err, want)
	}
}

// TestHellosNames calls greetings.Hellos with a list of names
// and checks we have a greeting for each name
func TestHellosNames(t *testing.T) {
	names := []string{
		"foo",
		"bar",
	}

	messages, err := Hellos(names)
	if messages == nil || err != nil {
		t.Fatalf(`Hellos(%v) = %q, %v error`, names, messages, err)
	}

	for _, name := range names {
		want := regexp.MustCompile(`\b` + name + `\b`)
		if !want.MatchString(messages[name]) {
			t.Fatalf(`Hello(%q) = %q, expected %#q`, name, messages[name], want)
		}
	}
}
