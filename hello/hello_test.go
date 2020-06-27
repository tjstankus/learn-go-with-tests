package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	expected := "Hello, TJ"
	value := Hello("TJ")

	if expected != value {
		t.Errorf("expected %q, actual value %q", expected, value)
	}
}
