package leetcode

import "sort"

func findUnsortedSubarray(nums []int) int {
	if sort.IntsAreSorted(nums) {
		return 0
	}
	cpy := make([]int, len(nums))
	copy(cpy, nums)
	sort.Ints(nums)
	left, right := 0, len(nums)-1
	for ; left < right; left++ {
		if cpy[left] != nums[left] {
			break
		}
	}
	for ; left < right; right-- {
		if cpy[right] != nums[right] {
			break
		}
	}
	return right - left + 1
}
