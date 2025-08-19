// Package hello provides simple greetings and math utilities.
package hello

import "strings"

// Greet returns a greeting message for the given name.
// If name is empty or whitespace, it uses "World".
func Greet(name string) string {
	trimmed := strings.TrimSpace(name)
	if trimmed == "" {
		trimmed = "World"
	}
	return "Hello, " + trimmed + "!"
}

// Sum returns the sum of the provided integers.
func Sum(nums ...int) int {
	s := 0
	for _, n := range nums {
		s += n
	}
	return s
}
