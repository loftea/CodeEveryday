package leetcode

func trap(height []int) int {
	res, left, right, leftHeight, rightHeight := 0, 0, len(height)-1, 0, 0
	for left <= right {
		if height[left] <= height[right] {
			if height[left] > leftHeight {
				leftHeight = height[left]
			} else {
				res += leftHeight - height[left]
			}
			left++
		} else {
			if height[right] >= rightHeight {
				rightHeight = height[right]
			} else {
				res += rightHeight - height[right]
			}
			right--
		}
	}
	return res
}
