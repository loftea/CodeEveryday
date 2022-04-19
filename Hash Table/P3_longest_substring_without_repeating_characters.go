package leetcode

import "math"

func lengthOfLongestSubstring(s string) int {
	LeftIndex := 0
	sLength := 0
	length := len(s)
	H := make([]int, math.MaxUint8)
	result := 0

	for index := 0; index < length; index++ {
		currByte := int(s[index])

		if LeftIndex < H[currByte] {
			// 更新窗口
			LeftIndex = H[currByte]
		}

		sLength = index - LeftIndex + 1
		result = max(sLength, result)
		H[currByte] = index + 1

	}

	return result
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
