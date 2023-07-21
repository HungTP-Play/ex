package ex

// Give n train stations.
// The train can only go forward.
// The cost of going from station i to station j is cost[i][j].
// Find the minimum cost of going from station 0 to station n-1.
//
// Example:
//
// 	Ex2(3, [][]int{
// 		[]int{0, 1, 3},
// 		[]int{1, 0, 1},
// 		[]int{2, 1, 0},
// 	})
//
// 	=> 2 (0 -> 1 -> 2)

// Subproblem: Find the minimum cost of going from station i to station j.
//
// Recurrence relation: Let's choose 3 stations (0->2), the minimum cost of going from
//
//	= Min {
//		 cost(0,2),
//		 cost(0,1) + cost(1,2),
//	}
//
// And with 4 stations (0->3), the minimum cost of going from
//
//	= Min {
//		 cost(0,3),
//		 cost(0,1) + cost(1,3),
//		 cost(0,2) + cost(2,3),
//	}
//
// And so on.
func Ex2_Recursive(i, j int, cost [][]int, loopCount *int) int {
	*loopCount++
	if i >= j {
		return 0
	}

	if i == j-1 {
		return cost[i][j]
	}

	min := cost[i][j]
	for k := i + 1; k < j; k++ {
		costCalc := Ex2_Recursive(i, k, cost, loopCount) + Ex2_Recursive(k, j, cost, loopCount)
		if costCalc < min {
			min = costCalc
		}
	}
	return min
}

func Ex2_Memoization(i, j int, cost [][]int, loopCount *int, memo [][]int) int {
	*loopCount++
	if i >= j {
		return 0
	}

	if i == j-1 {
		return cost[i][j]
	}

	if memo[i][j] != 0 {
		return memo[i][j]
	}

	memo[i][j] = cost[i][j]
	for k := i + 1; k < j; k++ {
		costCalc := Ex2_Memoization(i, k, cost, loopCount, memo) + Ex2_Memoization(k, j, cost, loopCount, memo)
		if costCalc < memo[i][j] {
			memo[i][j] = costCalc
		}
	}
	return memo[i][j]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Ex2_DynamicProgramming(i, j int, cost [][]int, loopCount *int, memo [][]int) int {
	memo[0][1] = cost[0][1]
	for k := 2; k <= j; k++ {
		*loopCount++
		min0k := cost[0][k]
		temp := memo[0][k-1] + cost[k-1][k]
		memo[0][k] = min(min0k, temp)
	}
	return memo[0][j]
}
