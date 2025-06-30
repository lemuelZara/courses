package main

func subarraySum(nums []int, k int) int {
	if len(nums) == 1 && nums[0] != k {
		return 0
	}

	sums := map[int]int{}
	sums[0] = 1
	actual := 0
	r := 0
	for _, v := range nums {
		actual += v
		if v, ok := sums[actual-k]; ok {
			r += v
		}

		sums[actual]++
	}

	return r
}
