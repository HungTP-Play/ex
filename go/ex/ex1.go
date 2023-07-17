package ex

import "fmt"

var loopCount int = 0

// Return total number of distinct subsequences of s that are equal to t.
//
// Example:
//
//	ex1("abc", "ab") = 1
//	ex1("babgbag", "bag") = 5 => ["ba g","bag", "ba    g", "b    ag", "  b  ag"]
func Ex1(s, t string) int {
	loopCount++
	if len(s) < len(t) {
		return 0
	}
	if len(t) == 0 {
		return 1
	}
	if len(s) == len(t) {
		if s == t {
			return 1
		}
		return 0
	}

	firstS := s[0]
	firstT := t[0]
	count := 0
	if firstS == firstT {
		count += Ex1(s[1:], t[1:])
	}
	count += Ex1(s[1:], t)
	fmt.Printf("loopCount: %d\n", loopCount)
	return count
}

type Stack struct {
	stack []struct {
		i int
		j int
	}
}

func (s *Stack) Push(i, j int) {
	s.stack = append(s.stack, struct {
		i int
		j int
	}{i, j})
}

func (s *Stack) Pop() (int, int) {
	if len(s.stack) == 0 {
		return -1, -1
	}
	last := s.stack[len(s.stack)-1]
	s.stack = s.stack[:len(s.stack)-1]
	return last.i, last.j
}

func (s *Stack) Empty() bool {
	return len(s.stack) == 0
}

func (s *Stack) Size() int {
	return len(s.stack)
}

func Ex1_2(s, t string) int {
	sLen := len(s)
	tLen := len(t)
	stack := Stack{}

	i := 0 // for t
	j := 0 // for s

	count := 0
	loopCount := 0

	for i < tLen {
		for j < sLen {
			loopCount++
			if s[j] == t[i] {
				if i == tLen-1 {
					// Last char of t
					count++
				} else {
					// Push to stack
					stack.Push(i, j)
					i++
				}
			}
			j++
		}

		if stack.Empty() {
			break
		}

		i, j = stack.Pop()
		j++
	}

	fmt.Printf("loopCount: %d\n", loopCount)
	return count
}
