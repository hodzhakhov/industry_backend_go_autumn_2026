package main

import "math"

type Stats struct {
	Count         int
	Sum, Min, Max int64
}

func Calc(nums []int64) Stats {
	ans := Stats{}

	if len(nums) < 2 {
		return ans
	}

	count := len(nums) - 1
	var sum int64 = 0
	var minDiff int64 = math.MaxInt64
	var maxDiff int64 = math.MinInt64

	for i := 1; i < len(nums); i++ {
		diff := nums[i] - nums[i - 1] 
		sum += diff
		minDiff = min(minDiff, diff)
		maxDiff = max(maxDiff, diff)
	}

	ans.Count = count
	ans.Sum = sum
	ans.Min = minDiff
	ans.Max = maxDiff

	return ans
}
