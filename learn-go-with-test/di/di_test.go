package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Ian")

	got := buffer.String()
	want := "Hello, Ian"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
