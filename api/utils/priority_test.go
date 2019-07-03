package utils

import "testing"

func TestSum(t *testing.T) {
	tables := []struct {
		priority int
		weight int
		result int

	}{
		{1, 1, 1},
		{1, 2, 1},
		{6, 6, 3},
		{10, 5, 2},
		{10, 10, 3},
	}

	for _, table := range tables {
		total := CalculatePriority(table.priority, table.weight)
		if total != table.result {
			t.Errorf("Sum of (%d+%d) was incorrect, got: %d, want: %d.", table.priority, table.weight, total, table.result)
		}
	}
}