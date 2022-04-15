package main

import (
	"fmt"
)

type NumArray struct {
	nums []int
	sums []int
}

func main() {
	nums := []int{1, 2, 3, 4, 5}
	obj := Constructor(nums)
	fmt.Println(obj)

	index, val := 1, 3
	obj.Update(index, val)
	fmt.Println(obj)

	left, right := 1, 2
	param_2 := obj.SumRange(left, right)
	fmt.Println(param_2)

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
