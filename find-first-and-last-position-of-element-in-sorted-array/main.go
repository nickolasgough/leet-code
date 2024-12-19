package main

func searchRange(nums []int, target int) []int {
	firstIndex := searchRangeHelper(nums, target, true)
	lastIndex := searchRangeHelper(nums, target, false)
	return []int{firstIndex, lastIndex}
}

func searchRangeHelper(nums []int, target int, findFirst bool) int {
	numsLen := len(nums)
	lastIndex := -1
	if numsLen <= 3 {
		for i, n := range nums {
			if n == target {
				if findFirst {
					return i
				}
				lastIndex = i
			}
		}
		return lastIndex
	}

	splitIndex := numsLen / 2
	firstHalf := nums[0:splitIndex]
	secondHalf := nums[splitIndex:numsLen]
	index := -1
	if findFirst {
		if containsTarget(firstHalf, target) {
			return searchRangeHelper(firstHalf, target, findFirst)
		}
		index = searchRangeHelper(secondHalf, target, findFirst)
	} else {
		if !containsTarget(secondHalf, target) {
			return searchRangeHelper(firstHalf, target, findFirst)
		}
		index = searchRangeHelper(secondHalf, target, findFirst)
	}
	if index != -1 {
		index += splitIndex
	}
	return index
}

func containsTarget(nums []int, target int) bool {
	numsLen := len(nums)
	firstEl := nums[0]
	lastEl := nums[numsLen-1]
	return firstEl <= target && target <= lastEl
}
