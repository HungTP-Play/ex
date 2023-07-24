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

func TestEx1Recursive(t *testing.T) {
	tests := []struct {
		s, t string
		want int
	}{
		// Base cases
		{"", "", 1},
		{"", "a", 0},
		{"a", "a", 1},
		{"aa", "a", 2},
		{"abc", "ab", 1},
		{"abc", "bc", 1},
		{"abc", "ac", 1},
		{"abc", "abc", 1},
		{"babgbag", "bag", 5},

		// Cases with repeated characters
		{"rabbbit", "rabbit", 3},
		{"babab", "ab", 3},
		{"aaaaa", "a", 5},
		{"aaaaa", "aa", 10},
		{"aaaaa", "aaa", 10},
		{"aaaaa", "aaaa", 5},
		{"aaaaa", "aaaaa", 1},
		{"abcabc", "abc", 4},
		{"abcabc", "ac", 3},
	}
	for _, tt := range tests {
		got := Ex1Recursive(tt.s, tt.t)
		if got != tt.want {
			t.Errorf("Ex1(%q, %q) = %d, want %d", tt.s, tt.t, got, tt.want)
		}
	}
}

func TestEx1Memoization(t *testing.T) {
	tests := []struct {
		s, t string
		want int
	}{
		// Base cases
		{"", "", 1},
		{"", "a", 0},
		{"a", "a", 1},
		{"aa", "a", 2},
		{"abc", "ab", 1},
		{"abc", "bc", 1},
		{"abc", "ac", 1},
		{"abc", "abc", 1},
		{"babgbag", "bag", 5},

		// Cases with repeated characters
		{"rabbbit", "rabbit", 3},
		{"babab", "ab", 3},
		{"aaaaa", "a", 5},
		{"aaaaa", "aa", 10},
		{"aaaaa", "aaa", 10},
		{"aaaaa", "aaaa", 5},
		{"aaaaa", "aaaaa", 1},
		{"abcabc", "abc", 4},
		{"abcabc", "ac", 3},
	}
	for _, tt := range tests {
		got := Ex1Memoization(tt.s, tt.t)
		if got != tt.want {
			t.Errorf("Ex1(%q, %q) = %d, want %d", tt.s, tt.t, got, tt.want)
		}
	}
}

func TestEx1DP(t *testing.T) {
	tests := []struct {
		s, t string
		want int
	}{
		// Base cases
		{"", "", 1},
		{"", "a", 0},
		{"a", "a", 1},
		{"aa", "a", 2},
		{"abc", "ab", 1},
		{"abc", "bc", 1},
		{"abc", "ac", 1},
		{"abc", "abc", 1},
		{"babgbag", "bag", 5},

		// Cases with repeated characters
		{"rabbbit", "rabbit", 3},
		{"babab", "ab", 3},
		{"aaaaa", "a", 5},
		{"aaaaa", "aa", 10},
		{"aaaaa", "aaa", 10},
		{"aaaaa", "aaaa", 5},
		{"aaaaa", "aaaaa", 1},
		{"abcabc", "abc", 4},
		{"abcabc", "ac", 3},
	}
	for _, tt := range tests {
		got := Ex1Dp(tt.s, tt.t)
		if got != tt.want {
			t.Errorf("Ex1(%q, %q) = %d, want %d", tt.s, tt.t, got, tt.want)
		}
	}
}

func BenchmarkEx1Recursive(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Ex1Recursive("babgbag", "bag")
	}
}

func BenchmarkEx1Memoization(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Ex1Memoization("babgbag", "bag")
	}
}

func BenchmarkEx1DP(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Ex1Dp("babgbag", "bag")
	}
}
