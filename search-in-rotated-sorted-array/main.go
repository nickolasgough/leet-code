package main

import (
	"slices"
)

func search(nums []int, target int) int {
	numsLen := len(nums)
	if numsLen <= 3 {
		for i, n := range nums {
			if n == target {
				return i
			}
		}
		return -1
	}

	isRotated := !isSorted(nums)
	if !isRotated {
		index, found := slices.BinarySearch(nums, target)
		if found {
			return index
		}
		return -1
	}

	isLenEven := numsLen%2 == 0
	splitIndex := numsLen / 2
	var firstHalf, secondHalf []int
	if isLenEven {
		firstHalf = nums[0:splitIndex]
		secondHalf = nums[splitIndex:numsLen]
	} else {
		middleEl := nums[splitIndex]
		if middleEl == target {
			return splitIndex
		}
		firstHalf = nums[0:splitIndex]
		secondHalf = nums[splitIndex+1 : numsLen]
	}

	if shouldSearch(firstHalf, target) || shouldNotSearch(secondHalf, target) {
		return search(firstHalf, target)
	}
	index := search(secondHalf, target)
	if index == -1 {
		return index
	}
	if isLenEven {
		return index + splitIndex
	}
	return index + splitIndex + 1
}

func shouldSearch(nums []int, target int) bool {
	return isSorted(nums) && containsTarget(nums, target)
}

func shouldNotSearch(nums []int, target int) bool {
	return isSorted(nums) && !containsTarget(nums, target)
}

func isSorted(nums []int) bool {
	numsLen := len(nums)
	firstEl := nums[0]
	lastEl := nums[numsLen-1]
	return firstEl < lastEl
}

func containsTarget(nums []int, target int) bool {
	numsLen := len(nums)
	firstEl := nums[0]
	lastEl := nums[numsLen-1]
	return firstEl <= target && target <= lastEl
}
