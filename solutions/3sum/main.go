package main

import (
	"slices"
)

// https://leetcode.com/problems/3sum/

func threeSum(nums []int) [][]int {
	slices.Sort(nums)
	firstPositiveIndex := slices.IndexFunc(nums, func(n int) bool {
		return n > 0
	})
	result := make([][]int, 0)
	if firstPositiveIndex == -1 {
		return result
	}

	leftIndex := 0
	rightIndex := len(nums) - 1
	leftEl := nums[leftIndex]
	rightEl := nums[rightIndex]
	for leftIndex != rightIndex && (leftEl <= 0 && rightEl >= 0) {
		temp := leftEl + rightEl
		nextEl := 0
		if temp > 0 {
			nextEl = nums[leftIndex+1]
		} else if temp < 0 {
			nextEl = nums[rightIndex-1]
		} else if hasZero {
			nextEl = 0
		}
		if leftEl+nextEl+rightEl == 0 {
			triplet := []int{leftEl, nextEl, rightEl}
			result = append(result, triplet)
		}

		leftEl = nums[leftIndex]
		rightEl = nums[rightIndex]
	}
}
