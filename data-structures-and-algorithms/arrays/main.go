package main

import "fmt"

func main() {
	fmt.Println(twoSumHash([]int{3, 2, 4}, 6))
	fmt.Println(twoSumTwoPointers([]int{3, 2, 4}, 6))
	fmt.Println(removeDuplicates([]int{1, 2}))
	fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
	fmt.Println(containsNearbyDuplicate([]int{1, 2, 3, 1, 2, 3}, 2))
	fmt.Println(moveZeroes([]int{1, 0, 0, 2, 0, 3}))
	fmt.Println(trap([]int{2, 0, 0, 2}))
	fmt.Println(subarraySum([]int{1, -1, 0}, 0))
}
