package main

import "testing"

func TestHello(t *testing.T) {

	checkMessageRight := func(t *testing.T, result, expected string) {
		t.Helper()
		if result != expected {
			t.Errorf("result '%s', expected '%s'", result, expected)
		}
	}

	t.Run("say hello to people", func(t *testing.T) {
		result := Hello("Murilo")
		expected := "Hello, Murilo!!!"

		checkMessageRight(t, result, expected)
	})

	t.Run("say 'Hello, World!!! when string is empty", func(t *testing.T) {
		result := Hello("")
		expected := "Hello, World!!!"

		checkMessageRight(t, result, expected)
	})
}
