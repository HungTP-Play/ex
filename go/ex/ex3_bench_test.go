package ex

import "testing"

func BenchmarkEx3Recursive(b *testing.B) {
	cases := []struct {
		Items    []Item
		Capacity int
		Want     []Item
	}{
		// Test Case 1: Capacity is very small compared to item weights
		{
			Items:    []Item{{W: 10, V: 20}, {W: 15, V: 30}, {W: 25, V: 40}},
			Capacity: 5,
			Want:     []Item{},
		},

		// Test Case 2: All items have the same weight
		{
			Items:    []Item{{W: 10, V: 20}, {W: 10, V: 30}, {W: 10, V: 40}},
			Capacity: 30,
			Want:     []Item{{W: 10, V: 40}, {W: 10, V: 30}, {W: 10, V: 20}},
		},

		// Test Case 3: All items have the same value
		{
			Items:    []Item{{W: 10, V: 50}, {W: 20, V: 50}, {W: 30, V: 50}},
			Capacity: 40,
			Want:     []Item{{W: 10, V: 50}, {W: 20, V: 50}},
		},

		// Test Case 4: Randomly generated items
		{
			Items:    []Item{{W: 10, V: 20}, {W: 20, V: 30}, {W: 30, V: 40}, {W: 40, V: 50}, {W: 50, V: 60}},
			Capacity: 100,
			Want:     []Item{{W: 40, V: 50}, {W: 30, V: 40}, {W: 20, V: 30}, {W: 10, V: 20}},
		},

		// Test Case 5: Small number of items
		{
			Items:    []Item{{W: 5, V: 10}, {W: 10, V: 20}, {W: 15, V: 30}},
			Capacity: 25,
			Want:     []Item{{W: 15, V: 30}, {W: 10, V: 20}},
		},

		// Test Case 6: Large number of items
		{
			Items: []Item{{W: 10, V: 20}, {W: 20, V: 30}, {W: 30, V: 40}, {W: 40, V: 50}, {W: 50, V: 60},
				{W: 60, V: 70}, {W: 70, V: 80}, {W: 80, V: 90}, {W: 90, V: 100}, {W: 100, V: 110}},
			Capacity: 500,
			Want: []Item{{W: 100, V: 110}, {W: 90, V: 100}, {W: 80, V: 90}, {W: 70, V: 80}, {W: 60, V: 70},
				{W: 10, V: 20}, {W: 40, V: 50}, {W: 30, V: 40}, {W: 20, V: 30}},
		},

		// Test Case 7: Values increase exponentially with weight
		{
			Items:    []Item{{W: 2, V: 1}, {W: 4, V: 4}, {W: 8, V: 27}, {W: 16, V: 256}, {W: 32, V: 3125}},
			Capacity: 50,
			Want:     []Item{{W: 32, V: 3125}, {W: 16, V: 256}, {W: 2, V: 1}},
		},
	}

	for _, c := range cases {
		loopCount := 0
		for i := 0; i < b.N; i++ {
			KnapsackRecursive(c.Items, c.Capacity, &loopCount)
		}
	}
}

func BenchmarkEx3DP(b *testing.B) {
	cases := []struct {
		Items    []Item
		Capacity int
		Want     []Item
	}{
		// Test Case 1: Capacity is very small compared to item weights
		{
			Items:    []Item{{W: 10, V: 20}, {W: 15, V: 30}, {W: 25, V: 40}},
			Capacity: 5,
			Want:     []Item{},
		},

		// Test Case 2: All items have the same weight
		{
			Items:    []Item{{W: 10, V: 20}, {W: 10, V: 30}, {W: 10, V: 40}},
			Capacity: 30,
			Want:     []Item{{W: 10, V: 40}, {W: 10, V: 30}, {W: 10, V: 20}},
		},

		// Test Case 3: All items have the same value
		{
			Items:    []Item{{W: 10, V: 50}, {W: 20, V: 50}, {W: 30, V: 50}},
			Capacity: 40,
			Want:     []Item{{W: 10, V: 50}, {W: 20, V: 50}},
		},

		// Test Case 4: Randomly generated items
		{
			Items:    []Item{{W: 10, V: 20}, {W: 20, V: 30}, {W: 30, V: 40}, {W: 40, V: 50}, {W: 50, V: 60}},
			Capacity: 100,
			Want:     []Item{{W: 40, V: 50}, {W: 30, V: 40}, {W: 20, V: 30}, {W: 10, V: 20}},
		},

		// Test Case 5: Small number of items
		{
			Items:    []Item{{W: 5, V: 10}, {W: 10, V: 20}, {W: 15, V: 30}},
			Capacity: 25,
			Want:     []Item{{W: 15, V: 30}, {W: 10, V: 20}},
		},

		// Test Case 6: Large number of items
		{
			Items: []Item{{W: 10, V: 20}, {W: 20, V: 30}, {W: 30, V: 40}, {W: 40, V: 50}, {W: 50, V: 60},
				{W: 60, V: 70}, {W: 70, V: 80}, {W: 80, V: 90}, {W: 90, V: 100}, {W: 100, V: 110}},
			Capacity: 500,
			Want: []Item{{W: 100, V: 110}, {W: 90, V: 100}, {W: 80, V: 90}, {W: 70, V: 80}, {W: 60, V: 70},
				{W: 10, V: 20}, {W: 40, V: 50}, {W: 30, V: 40}, {W: 20, V: 30}},
		},

		// Test Case 7: Values increase exponentially with weight
		{
			Items:    []Item{{W: 2, V: 1}, {W: 4, V: 4}, {W: 8, V: 27}, {W: 16, V: 256}, {W: 32, V: 3125}},
			Capacity: 50,
			Want:     []Item{{W: 32, V: 3125}, {W: 16, V: 256}, {W: 2, V: 1}},
		},
	}

	for _, c := range cases {
		loopCount := 0
		for i := 0; i < b.N; i++ {
			KnapsackDP(c.Items, c.Capacity, &loopCount)
		}
	}
}
