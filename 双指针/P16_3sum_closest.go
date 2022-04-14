import (
	"sort"
	"math"
	"fmt"
)

func threeSumClosest(nums []int, target int) int {
	result := math.MaxInt32
	delta := result - target
	sort.Ints(nums)
	left, right, index, addNum, length, tmp := 0, 0, 0, 0, len(nums), delta
	for index = 1; index < length-1; index++ {
		left, right = 0, length-1
		if index > 1 && nums[index] == nums[index-1] {
			left = index - 1
		}
		for left < index && right > index {
			if left > 0 && nums[left] == nums[left-1] {
				left++
				continue
			}
			if right < length-1 && nums[right] == nums[right+1] {
				right--
				continue
			}
			addNum = nums[left] + nums[right] + nums[index]
			// fmt.Printf("index=%d\n", index)
			// fmt.Printf("%d+%d+%d=%d\n", nums[left], nums[index], nums[right], addNum)
			tmp = abs(addNum - target)
			if tmp < delta {
				result = addNum
				delta = tmp
			}

			if tmp == 0 {
				return addNum
			}

			if addNum > target {
				right--
			} else {
				left++
			}
		}
	}
	return result

}

func abs(x int) int {
	if x >= 0 {
		return x
	}
	return -1 * x
}
