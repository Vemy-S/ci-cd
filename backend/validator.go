package main

import "strings"

func ValidateEmail(email string) bool {
	if !strings.Contains(email, "@") {
		return false
	}
	if !strings.Contains(email, ".") {
		return false
	}
	if len(email) < 5 {
		return false
	}
	return true
}
