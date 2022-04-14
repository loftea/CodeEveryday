/*
给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有和为 0 且不重复的三元组。

注意：答案中不可以包含重复的三元组。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/3sum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
import "sort"

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	result, left, right, index, addNum, length := make([][]int, 0), 0, 0, 0, 0, len(nums)
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
			if addNum == 0 {
				result = append(result, []int{nums[left], nums[index], nums[right]})
				left++
				right--
			} else if addNum > 0 {
				right--
			} else {
				left++
			}
		}
	}
	return result
}