package ex

// This is 0/1 Knapsack problem.

type Item struct {
	W int
	V int
}

func TotalValue(items []Item) int {
	total := 0
	for _, item := range items {
		total += item.V
	}
	return total
}

func KnapsackRecursive(items []Item, capacity int, loopCount *int) []Item {
	if len(items) == 0 || capacity == 0 {
		return []Item{}
	}
	*loopCount++

	// Case 1: The last item is not included
	withoutLast := KnapsackRecursive(items[:len(items)-1], capacity, loopCount)

	// Case 2: The last item is included
	index := len(items) - 1
	lastItem := items[index]
	if lastItem.W > capacity {
		return withoutLast
	}

	newCapacity := capacity - lastItem.W
	withLast := KnapsackRecursive(items[:index], newCapacity, loopCount)
	withLast = append(withLast, lastItem)

	if TotalValue(withLast) > TotalValue(withoutLast) {
		return withLast
	}
	return withoutLast
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func KnapsackDP(items []Item, capacity int, loopCount *int) []Item {
	mem := make([][]int, len(items)+1)
	for i := range mem {
		mem[i] = make([]int, capacity+1)
	}

	for i := 1; i <= len(items); i++ {
		for j := 1; j <= capacity; j++ {
			*loopCount++
			if items[i-1].W > j {
				mem[i][j] = mem[i-1][j]
			} else {
				remainCapacity := j - items[i-1].W
				prevMax := mem[i-1][remainCapacity]
				itemIndex := i - 1
				mem[i][j] = max(prevMax+items[itemIndex].V, mem[i-1][j])
			}
		}
	}

	// Backtrack
	result := []Item{}
	i := len(items)
	j := capacity
	for i > 0 && j > 0 {
		*loopCount++
		if mem[i][j] != mem[i-1][j] {
			result = append(result, items[i-1])
			j -= items[i-1].W
		}
		i--
	}

	return result
}

func Knapsack(items []Item, capacity int, dp bool, loopCount *int) []Item {
	if dp {
		return KnapsackDP(items, capacity, loopCount)
	}
	return KnapsackRecursive(items, capacity, loopCount)
}
