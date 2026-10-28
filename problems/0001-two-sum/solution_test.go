package twosum

import (
	"slices"
	"testing"
)

func TestSolution(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   []int
	}{
		{name: "exemple 1", nums: []int{2, 7, 11, 15}, target: 9, want: []int{0, 1}},
		{name: "exemple 2", nums: []int{3, 2, 4}, target: 6, want: []int{1, 2}},
		{name: "exemple 3", nums: []int{3, 3}, target: 6, want: []int{0, 1}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := twoSum(tt.nums, tt.target)
			slices.Sort(got)

			if !slices.Equal(got, tt.want) {
				t.Errorf("twoSum(%v, %d) = %v, want %v", tt.nums, tt.target, got, tt.want)
			}
		})
	}
}
