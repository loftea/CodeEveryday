package leetcode

type NumArray struct {
	nums []int
	sums []int
}

func Constructor(nums []int) NumArray {
	a := new(NumArray)
	a.nums = nums
	a.sums = []int{nums[0]}
	for i := 1; i < len(nums); i++ {
		a.sums = append(a.sums, a.sums[i-1]+nums[i])
	}
	return *a
}

func (this *NumArray) Update(index int, val int) {
	old := this.nums[index]
	this.nums[index] = val
	delta := val - old
	for i := index; i < len(this.sums); i++ {
		this.sums[i] += delta
	}
}

func (this *NumArray) SumRange(left int, right int) int {
	if left == 0 {
		return this.sums[right]
	}
	return this.sums[right] - this.sums[left-1]
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * obj.Update(index,val);
 * param_2 := obj.SumRange(left,right);
 */
