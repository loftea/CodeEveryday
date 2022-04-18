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
			// 如果 H[currByte] 都没有被更新过，但是仍然有 LeftIndex <= H[currByte] 成立，
			// 不应该更新，因为这个时候 0<=0 成立，
			// 但是 LeftIndex 不应该变成 0+1=1，而仍然应该是 0
			LeftIndex = H[currByte]
			sLength = index - LeftIndex + 1
			result = max(sLength, result)
			H[currByte] = index + 1
			continue
		}

		sLength++
		result = max(result, sLength)
		H[currByte] = index + 1
		continue

	}

	return result
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
