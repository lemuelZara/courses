package main

import (
	"sort"
)

func twoSumHash(nums []int, target int) []int {
	need := map[int]int{}

	for i, v := range nums {
		want := target - v
		if j, ok := need[want]; ok {
			return []int{j, i}
		}

		need[v] = i
	}

	return []int{}
}

type pair struct {
	value, idx int
}

func twoSumTwoPointers(nums []int, target int) []int {
	pairs := make([]pair, len(nums))
	for i, v := range nums {
		pairs[i] = pair{idx: i, value: v}
	}

	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].value < pairs[j].value
	})

	left, right := 0, len(pairs)-1
	for left < right {
		sum := pairs[left].value + pairs[right].value

		if sum == target {
			return []int{pairs[left].idx, pairs[right].idx}
		}

		if sum > target {
			right--
		} else {
			left++
		}
	}

	return []int{}
}
