package main

import (
	"testing"
)

// TestHello TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestHello(t *testing.T) {
	got := Hello()
	want := "Hello, World!"
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
