package main

func removeDuplicates(nums []int) []int {
	l, r := 0, 0
	for r < len(nums) {
		if nums[l] != nums[r] {
			l++
			nums[l] = nums[r]
		}

		r++
	}

	return nums[:l+1]
}
