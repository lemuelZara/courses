package main

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	maxActual, maxGlobal := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		maxActual = max(nums[i], maxActual+nums[i])
		maxGlobal = max(maxActual, maxGlobal)
	}

	return maxGlobal
}
