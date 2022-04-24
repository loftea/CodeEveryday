package leetcode

import "fmt"

func trap(height []int) int {
	res := 0
	tmp := 0
	// begin := false
	leftHeight := height[0]
	leftIndex := 0
	for i := 1; i < len(height); i++ {
		if leftHeight <= height[i] {
			// begin = false
			leftHeight = height[i]
			res += tmp
			fmt.Printf("add %d\n", tmp)
		} else {
			// begin = true
			tmp += leftHeight - height[i]
		}

	}
	return res
}
