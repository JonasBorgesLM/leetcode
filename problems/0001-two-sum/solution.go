// Package twosum resolve o problema 1 do LeetCode.
// https://leetcode.com/problems/two-sum/
package twosum

func twoSum(nums []int, target int) []int {
	seen := make(map[int]int)

	for i, num := range nums {
		complement := target - num

		if idx, exists := seen[complement]; exists {
			return []int{i, idx}
		}

		seen[num] = i
	}

	return nil
}
