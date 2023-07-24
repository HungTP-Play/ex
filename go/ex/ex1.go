package ex

var loopCount int = 0

// Return total number of distinct subsequences of s that are equal to t.
//
// Example:
//
//	ex1("abc", "ab") = 1
//	ex1("babgbag", "bag") = 5 => ["ba g","bag", "ba    g", "b    ag", "  b  ag"]
func Ex1Recursive(s, t string) int {
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
		count += Ex1Recursive(s[1:], t[1:])
	}
	count += Ex1Recursive(s[1:], t)
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

func Ex1Memoization(s, t string) int {
	if s == "" && t == "" {
		return 1
	}

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

	return count
}

// Dynamic programming solution
func Ex1Dp(s string, t string) int {
	if s == "" && t == "" {
		return 1
	}

	m, n := len(t), len(s)

	// Initialize the dp array
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
		dp[i][0] = 0
	}
	for j := 0; j <= n; j++ {
		dp[0][j] = 1
	}

	// Compute the dp array
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if t[i-1] == s[j-1] {
				dp[i][j] = dp[i-1][j-1] + dp[i][j-1]
			} else {
				dp[i][j] = dp[i][j-1]
			}
		}
	}

	return dp[m][n]
}
