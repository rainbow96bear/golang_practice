package main

import "testing"

func TestAdd(t *testing.T) {
	result := Add(4, 7)
	expected := 11
	if result != expected {
		t.Errorf("Expected %d but got %d", expected, result)
	}
}