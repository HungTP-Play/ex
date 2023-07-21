package ex

import "testing"

func Equal(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, x := range a {
		if x != b[i] {
			return false
		}
	}
	return true
}

func TestEx1(t *testing.T) {
	tests := []struct {
		s, t string
		want int
	}{
		{"abc", "ab", 1},
		{"babgbag", "bag", 5},
	}
	for _, tt := range tests {
		got := Ex1(tt.s, tt.t)
		if got != tt.want {
			t.Errorf("Ex1(%q, %q) = %d, want %d", tt.s, tt.t, got, tt.want)
		}
	}
}

func TestEx1_2(t *testing.T) {
	tests := []struct {
		s, t string
		want int
	}{
		{"abc", "ab", 1},
		{"babgbag", "bag", 5},
	}
	for _, tt := range tests {
		got := Ex1_2(tt.s, tt.t)
		if got != tt.want {
			t.Errorf("Ex1(%q, %q) = %d, want %d", tt.s, tt.t, got, tt.want)
		}
	}
}

