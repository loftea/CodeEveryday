package leetcode

import (
	"math"
	"sort"
)

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)

	result, left, right, index1, index2, addNum, length := make([][]int, 0), 0, 0, 0, 0, 0, len(nums)

	pre_c, pre_d := math.MaxInt32, math.MaxInt32

	for index1 = 1; index1 < length-2; index1++ {
		for index2 = index1 + 1; index2 < length-1; index2++ {

			if nums[index1] == pre_c && nums[index2] == pre_d {
				continue
			}

			left, right = 0, length-1

			for left < index1 && right > index2 {

				if left > 0 && nums[left] == nums[left-1] {
					left++
					continue
				}

				if right < length-1 && nums[right] == nums[right+1] {
					right--
					continue
				}

				addNum = nums[left] + nums[right] + nums[index1] + nums[index2]

				if addNum == target {
					result = append(result,
						[]int{nums[left],
							nums[index1],
							nums[index2],
							nums[right]})
					left++
					right--
				} else if addNum > target {
					right--
				} else {
					left++
				}

			}

			pre_c = nums[index1]
			pre_d = nums[index2]
		}
	}
	return result
}
