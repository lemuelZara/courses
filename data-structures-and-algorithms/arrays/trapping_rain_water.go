package main

func trap(height []int) int {
	n := len(height)
	if n == 0 {
		return 0
	}

	water := 0
	maxLeft := 0
	maxRight := 0

	left, right := 0, n-1
	for left < right {
		if height[left] < height[right] {
			if height[left] >= maxLeft {
				maxLeft = height[left]
			}
			water += max(maxLeft-height[left], 0)
			left++

		} else {
			if height[right] >= maxRight {
				maxRight = height[right]
			}
			water += max(maxRight-height[right], 0)
			right--
		}
	}

	return water
}
