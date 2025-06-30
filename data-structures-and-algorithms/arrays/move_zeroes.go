package main

func moveZeroes(nums []int) []int {
	write := 0
	for i := range nums {
		if nums[i] != 0 {
			if i != write {
				nums[write] = nums[i]
				// or nums[write], nums[i] = nums[i], nums[write]
				// ex: i = 1, write = 0
				// ['0, 1', 0, 2] after swap ['1, 0', 0, 2]
				//
			}

			write++
		}
	}

	for write < len(nums) {
		nums[write] = 0
		write++
	}

	return nums
}
