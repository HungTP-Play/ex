package ex

import (
	"testing"
	"time"
)

func EqualIgnoreOrder(a, b []Item) bool {
	if len(a) != len(b) {
		return false
	}
	for _, v := range a {
		found := false
		for _, w := range b {
			if v == w {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}

func TestEx3_01Knapsack(t *testing.T) {
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

	loopCount := 0

	for _, c := range cases {
		recursiveStartTime := time.Now()
		got := Knapsack(c.Items, c.Capacity, false, &loopCount)
		recursiveEndTime := time.Now()
		t.Logf(" Knapsack Recursive: %v\n", recursiveEndTime.Sub(recursiveStartTime))
		t.Logf(" Knapsack Recursive Loop Count: %d\n", loopCount)
		loopCount = 0

		if !EqualIgnoreOrder(got, c.Want) {
			t.Errorf("Knapsack(%v, %d, false) == %v, want %v", c.Items, c.Capacity, got, c.Want)
		}

		dpStartTime := time.Now()
		got = Knapsack(c.Items, c.Capacity, true, &loopCount)
		dpEndTime := time.Now()
		t.Logf("Knapsack DP: %v\n", dpEndTime.Sub(dpStartTime))
		t.Logf("Knapsack DP Loop Count: %d\n", loopCount)
		loopCount = 0

		if !EqualIgnoreOrder(got, c.Want) {
			t.Errorf("Knapsack(%v, %d, false) == %v, want %v", c.Items, c.Capacity, got, c.Want)
		}
	}
}
