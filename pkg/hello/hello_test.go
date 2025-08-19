package hello

import "testing"

func TestGreet(t *testing.T) {
	tests := []struct {
		name string
		in   string
		out  string
	}{
		{"empty uses World", "", "Hello, World!"},
		{"spaces uses World", "   \t\n", "Hello, World!"},
		{"name works", "Alice", "Hello, Alice!"},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := Greet(tc.in)
			if got != tc.out {
				t.Fatalf("Greet(%q) = %q; want %q", tc.in, got, tc.out)
			}
		})
	}
}

func TestSum(t *testing.T) {
	tests := []struct {
		name string
		in   []int
		out  int
	}{
		{"none", nil, 0},
		{"single", []int{5}, 5},
		{"several", []int{1, 2, 3}, 6},
		{"negatives", []int{-2, 5, -3}, 0},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := Sum(tc.in...)
			if got != tc.out {
				t.Fatalf("Sum%v = %d; want %d", tc.in, got, tc.out)
			}
		})
	}
}
