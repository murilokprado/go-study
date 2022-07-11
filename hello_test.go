package main

import "testing"

func TestHello(t *testing.T) {
	result := Hello("Murilo")
	expected := "Hello, Murilo!!!"

	if result != expected {
		t.Errorf("result '%s', expected '%s'", result, expected)
	}

}
