package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	expected := "Hello, world"
	value := Hello()

	if expected != value {
		t.Errorf("expected %q, actual value %q", expected, value)
	}
}
