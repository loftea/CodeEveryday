package dailychallenge

import "sort"

func maxOperations(nums []int, k int) int {
	sort.Ints(nums)
	l, r := 0, len(nums)-1
	res := 0
	for l < r {
		if nums[l]+nums[r] < k {
			l++
		} else if nums[l]+nums[r] > k {
			r--
		} else {
			res++
			l++
			r--
		}
	}
	return res
}
