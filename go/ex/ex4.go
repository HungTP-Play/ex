package ex

import (
	"strings"
)

func Ex4LongestCommonSubsequenceRecursive(s1, s2 string, countLoop *int) string {
	if len(s1) == 0 || len(s2) == 0 {
		return ""
	}

	*countLoop++

	s1Len := len(s1)

	// Case 1: The last character is not included
	withoutLast := Ex4LongestCommonSubsequenceRecursive(s1[:s1Len-1], s2, countLoop)

	// Case 2: The last character is included
	lastChar := s1[s1Len-1]
	index := strings.Index(s2, string(lastChar))
	if index == -1 {
		return withoutLast
	}

	withLast := Ex4LongestCommonSubsequenceRecursive(s1[:s1Len-1], s2[:index], countLoop)
	withLast += string(lastChar)

	if len(withLast) >= len(withoutLast) {
		return withLast
	}

	return withoutLast
}

func Ex4LongestCommonSubsequenceDP(s1, s2 string, countLoop *int) string {
	mem := make(map[int]map[int]string)
	for i := 0; i <= len(s1); i++ {
		mem[i] = make(map[int]string)
	}

	for i := 0; i <= len(s1); i++ {
		*countLoop++
		for j := 0; j <= len(s2); j++ {
			if i == 0 || j == 0 {
				mem[i][j] = ""
			} else {
				if s1[i-1] == s2[j-1] {
					mem[i][j] = mem[i-1][j-1] + string(s1[i-1])
				} else {
					if len(mem[i-1][j]) > len(mem[i][j-1]) {
						mem[i][j] = mem[i-1][j]
					} else {
						mem[i][j] = mem[i][j-1]
					}
				}
			}
		}
	}

	return mem[len(s1)][len(s2)]
}
