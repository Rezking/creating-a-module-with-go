package hello

import (
	"regexp"
	"testing"
)

// TestHelloName calls hello.Hello with a name
func TestHelloName(t *testing.T) {
	name := "Ore"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello("Ore")
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("Ore") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}

// Test Hello with an empty string
func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
	}
}
