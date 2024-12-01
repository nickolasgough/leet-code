package main

func maxArea(h []int) int {
	hLen := len(h)
	leftIndex := 0
	rightIndex := hLen - 1
	maxArea := 0
	for leftIndex != rightIndex {
		length := rightIndex - leftIndex
		leftHeight := h[leftIndex]
		rightHeight := h[rightIndex]
		height := 0
		if leftHeight < rightHeight {
			leftIndex += 1
			height = leftHeight
		} else {
			rightIndex -= 1
			height = rightHeight
		}
		area := length * height
		if area > maxArea {
			maxArea = area
		}
	}
	return maxArea
}
