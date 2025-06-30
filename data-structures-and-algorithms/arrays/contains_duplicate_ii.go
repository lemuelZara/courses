package main

import "math"

func containsNearbyDuplicate(nums []int, k int) bool {
	unique := map[int]int{}
	for i, v := range nums {
		if v, ok := unique[v]; ok {
			if math.Abs(float64(v-i)) <= float64(k) {
				return true
			}
		}

		unique[v] = i
	}

	return false
}
