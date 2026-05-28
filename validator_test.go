package main

import "testing"

func TestValidateEmail(t *testing.T) {
	expected := true
	result := ValidateEmail("jeremyjvf16@gmail.com")

	if result != expected {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestValidateEmail_Invalid(t *testing.T) {
	expected := false
	result := ValidateEmail("jeremygmailcom")

	if result != expected {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}
