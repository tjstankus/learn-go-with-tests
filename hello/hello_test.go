package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	assertExpectedValue := func(t *testing.T, expected, value string) {
		t.Helper()
		if expected != value {
			t.Errorf("expected %q, actual value %q", expected, value)
		}
	}

	t.Run("say hello with name", func(t *testing.T) {
		expected := "Hello, TJ"
		value := Hello("TJ", "English")
		assertExpectedValue(t, expected, value)
	})

	t.Run("say hello with default name value", func(t *testing.T) {
		expected := "Hello, world"
		value := Hello("", "English")
		assertExpectedValue(t, expected, value)
	})

	t.Run("say hello in Spanish", func(t *testing.T) {
		expected := "Hola, TJ"
		value := Hello("TJ", "Spanish")
		assertExpectedValue(t, expected, value)
	})
}
