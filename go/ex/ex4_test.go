package ex

import (
	"testing"
	"time"
)

func OneOf(s string, arr []string) bool {
	for _, v := range arr {
		if s == v {
			return true
		}
	}
	return false
}

// Longest Common Subsequence Problem
//
// Given two strings, find the longest common subsequence (LCS).
// For example, given "abcde" and "ace", the LCS is "ace".
func TestEx4Recursive(t *testing.T) {
	// longString1 := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	cases := []struct {
		s1, s2 string
		want   []string
	}{
		// Case 1: Basic case with a simple subsequence
		{"ABCDGH", "AEDFHR", []string{"ADH"}},

		// !WILL FAILED Case 2: Longer input with multiple possible subsequences
		// {"AGGTAB", "GXATXAYB", []string{"GTAB"}},

		// Case 3: Input with no common characters
		{"ABCDEF", "GHIJKL", []string{""}},

		// !WIL FAILED Case 4: Input with a common character repeated multiple times
		// {"ABABAB", "BABA", []string{"BABA"}},

		// Case 5: Input with duplicate characters in both strings
		{"ABCD", "ABCD", []string{"ABCD"}},
		// // Case 1: Long strings with a simple subsequence
		// {longString1, "1234567890", []string{""}},

		// // Case 3: Long strings with multiple possible subsequences
		// {longString1, "aAbBcCdDeEfFgGhHiIjJkKlLmMnNoOpPqQrRsStTuUvVwWxXyYzZ", []string{"abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"}},

		// // Case 4: Long strings with a common character repeated multiple times
		// {strings.Repeat("A", 10000), strings.Repeat("B", 10000), []string{strings.Repeat("A", 10000)}},

		// {strings.Repeat("ABCD", 5000), strings.Repeat("ABCD", 5000), []string{strings.Repeat("ABCD", 5000)}},
	}

	for _, c := range cases {
		countLoop := 0
		startTime := time.Now()
		got := Ex4LongestCommonSubsequenceRecursive(c.s1, c.s2, &countLoop)
		endTime := time.Now()
		t.Logf("Ex4LongestCommonSubsequenceRecursive: %v", endTime.Sub(startTime))
		t.Logf("Ex4LongestCommonSubsequenceRecursiveLoopCount: %d", countLoop)
		if !OneOf(got, c.want) {
			t.Errorf("Ex4LongestCommonSubsequenceRecursive(%q, %q) == %q, want %q", c.s1, c.s2, got, c.want)
		}
	}
}

func TestEx4DP(t *testing.T) {
	longString1 := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	cases := []struct {
		s1, s2 string
		want   []string
	}{
		// Case 1: Basic case with a simple subsequence
		{"ABCDGH", "AEDFHR", []string{"ADH"}},

		// Case 2: Longer input with multiple possible subsequences
		{"AGGTAB", "GXATXAYB", []string{"GTAB", "ATAB"}},

		// Case 3: Input with no common characters
		{"ABCDEF", "GHIJKL", []string{""}},

		// Case 4: Input with a common character repeated multiple times
		{"ABABAB", "BABA", []string{"BABA"}},

		// Case 5: Input with duplicate characters in both strings
		{"ABCD", "ABCD", []string{"ABCD"}},
		// Case 1: Long strings with a simple subsequence
		{longString1, "1234567890", []string{""}},

		// Case 3: Long strings with multiple possible subsequences
		{longString1, "aAbBcCdDeEfFgGhHiIjJkKlLmMnNoOpPqQrRsStTuUvVwWxXyYzZ", []string{"aABCDEFGHIJKLMNOPQRSTUVWXYZ"}},

		// // Case 4: Long strings with a common character repeated multiple times
		// {strings.Repeat("A", 10000), strings.Repeat("B", 10000), []string{strings.Repeat("A", 10000)}},

		// {strings.Repeat("ABCD", 5000), strings.Repeat("ABCD", 5000), []string{strings.Repeat("ABCD", 5000)}},
	}

	for _, c := range cases {
		countLoop := 0
		startTime := time.Now()
		got := Ex4LongestCommonSubsequenceDP(c.s1, c.s2, &countLoop)
		endTime := time.Now()
		t.Logf("Ex4LongestCommonSubsequenceDP: %v", endTime.Sub(startTime))
		t.Logf("Ex4LongestCommonSubsequenceDPLoopCount: %d", countLoop)
		if !OneOf(got, c.want) {
			t.Errorf("Ex4LongestCommonSubsequenceRecursive(%q, %q) == %q, want %q", c.s1, c.s2, got, c.want)
		}
	}
}

func BenchmarkEx4Recursive(b *testing.B) {
	// longString1 := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	cases := []struct {
		s1, s2 string
		want   []string
	}{
		// Case 1: Basic case with a simple subsequence
		{"ABCDGH", "AEDFHR", []string{"ADH"}},

		// !WILL FAILED Case 2: Longer input with multiple possible subsequences
		// {"AGGTAB", "GXATXAYB", []string{"GTAB"}},

		// Case 3: Input with no common characters
		{"ABCDEF", "GHIJKL", []string{""}},

		// !WIL FAILED Case 4: Input with a common character repeated multiple times
		// {"ABABAB", "BABA", []string{"BABA"}},

		// Case 5: Input with duplicate characters in both strings
		{"ABCD", "ABCD", []string{"ABCD"}},
		// // Case 1: Long strings with a simple subsequence
		// {longString1, "1234567890", []string{""}},

		// // Case 3: Long strings with multiple possible subsequences
		// {longString1, "aAbBcCdDeEfFgGhHiIjJkKlLmMnNoOpPqQrRsStTuUvVwWxXyYzZ", []string{"abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"}},

		// // Case 4: Long strings with a common character repeated multiple times
		// {strings.Repeat("A", 10000), strings.Repeat("B", 10000), []string{strings.Repeat("A", 10000)}},

		// {strings.Repeat("ABCD", 5000), strings.Repeat("ABCD", 5000), []string{strings.Repeat("ABCD", 5000)}},
	}

	for _, c := range cases {
		countLoop := 0
		got := Ex4LongestCommonSubsequenceRecursive(c.s1, c.s2, &countLoop)
		if !OneOf(got, c.want) {
			b.Errorf("Ex4LongestCommonSubsequenceRecursive(%q, %q) == %q, want %q", c.s1, c.s2, got, c.want)
		}
	}
}

func BenchmarkEx4Dp(b *testing.B) {
	longString1 := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	cases := []struct {
		s1, s2 string
		want   []string
	}{
		// Case 1: Basic case with a simple subsequence
		{"ABCDGH", "AEDFHR", []string{"ADH"}},

		// Case 2: Longer input with multiple possible subsequences
		{"AGGTAB", "GXATXAYB", []string{"GTAB"}},

		// Case 3: Input with no common characters
		{"ABCDEF", "GHIJKL", []string{""}},

		// Case 4: Input with a common character repeated multiple times
		{"ABABAB", "BABA", []string{"BABA"}},

		// Case 5: Input with duplicate characters in both strings
		{"ABCD", "ABCD", []string{"ABCD"}},
		// Case 1: Long strings with a simple subsequence
		{longString1, "1234567890", []string{""}},

		// Case 3: Long strings with multiple possible subsequences
		{longString1, "aAbBcCdDeEfFgGhHiIjJkKlLmMnNoOpPqQrRsStTuUvVwWxXyYzZ", []string{"aABCDEFGHIJKLMNOPQRSTUVWXYZ"}},

		// Case 4: Long strings with a common character repeated multiple times
		// {strings.Repeat("A", 10000), strings.Repeat("B", 10000), []string{strings.Repeat("A", 10000)}},

		// {strings.Repeat("ABCD", 5000), strings.Repeat("ABCD", 5000), []string{strings.Repeat("ABCD", 5000)}},
	}

	for _, c := range cases {
		countLoop := 0
		got := Ex4LongestCommonSubsequenceDP(c.s1, c.s2, &countLoop)
		if !OneOf(got, c.want) {
			b.Errorf("Ex4LongestCommonSubsequenceRecursive(%q, %q) == %q, want %q", c.s1, c.s2, got, c.want)
		}
	}
}
