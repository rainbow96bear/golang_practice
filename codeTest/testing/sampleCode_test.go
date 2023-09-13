package main

import "testing"

func TestAdd1(t *testing.T) {
	result := Add(4, 7)
	expected1 := 11
	expected2 := 13
	if result != expected1 {
		t.Errorf("Expected %d but got %d", expected1, result)
	}
	if result != expected2 {
		t.Errorf("Expected %d but got %d", expected2, result)
	}
}
func TestAdd2(t *testing.T) {
	result := Add(4, 7)
	expected := 11
	
	if result != expected {
		t.Errorf("Expected %d but got %d", expected, result)
	}
}