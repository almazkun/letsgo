package greetings

import (
	"testing"
	"regexp"
)

func TestHelloName(t *testing.T) {
	name := "Yelena"
	want := regexp.MustCompile(`\b`+name+`\b`)
	msg, err := Hello("Yelena")
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("Yelena") = %q, %v, want match for %#q, nill`, msg, err, want)
	}
}

func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
	}
}