package leetcode

type SegmentTree struct {
	data, tree, lazy []int
	left, right      int
	merge            func(i, j int) int
}

type NumArray struct {
	st *SegmentTree
}

// Constructor define
func Constructor(nums []int) NumArray {
	st := SegmentTree{}
	st.Init(nums, func(i, j int) int {
		return i + j
	})
	return NumArray{st: &st}
}

// Update define
func (this *NumArray) Update(i int, val int) {
	this.st.Update(i, val)
}

// SumRange define
func (this *NumArray) SumRange(i int, j int) int {
	return this.st.Query(i, j)
}

// Init define
func (st *SegmentTree) Init(nums []int, oper func(i, j int) int) {
	st.merge = oper
	data, tree, lazy := make([]int, len(nums)), make([]int, 4*len(nums)), make([]int, 4*len(nums))
	for i := 0; i < len(nums); i++ {
		data[i] = nums[i]
	}
	st.data, st.tree, st.lazy = data, tree, lazy
	if len(nums) > 0 {
		st.buildSegmentTree(0, 0, len(nums)-1)
	}
}

// 在 treeIndex 的位置创建 [left....right] 区间的线段树
func (st *SegmentTree) buildSegmentTree(treeIndex, left, right int) {
	if left == right {
		st.tree[treeIndex] = st.data[left]
		return
	}
	midTreeIndex, leftTreeIndex, rightTreeIndex := left+(right-left)>>1, st.leftChild(treeIndex), st.rightChild(treeIndex)
	st.buildSegmentTree(leftTreeIndex, left, midTreeIndex)
	st.buildSegmentTree(rightTreeIndex, midTreeIndex+1, right)
	st.tree[treeIndex] = st.merge(st.tree[leftTreeIndex], st.tree[rightTreeIndex])
}

func (st *SegmentTree) leftChild(index int) int {
	return 2*index + 1
}

func (st *SegmentTree) rightChild(index int) int {
	return 2*index + 2
}

func (st *SegmentTree) Query(left, right int) int {
	if len(st.data) > 0 {
		return st.queryInTree(0, 0, len(st.data)-1, left, right)
	}
	return 0
}

// 在以 treeIndex 为根的线段树中 [left...right] 的范围里，搜索区间 [queryLeft...queryRight] 的值
func (st *SegmentTree) queryInTree(treeIndex, left, right, queryLeft, queryRight int) int {
	if left == queryLeft && right == queryRight {
		return st.tree[treeIndex]
	}
	midTreeIndex, leftTreeIndex, rightTreeIndex := left+(right-left)>>1, st.leftChild(treeIndex), st.rightChild(treeIndex)
	if queryLeft > midTreeIndex {
		return st.queryInTree(rightTreeIndex, midTreeIndex+1, right, queryLeft, queryRight)
	} else if queryRight <= midTreeIndex {
		return st.queryInTree(leftTreeIndex, left, midTreeIndex, queryLeft, queryRight)
	}
	return st.merge(st.queryInTree(leftTreeIndex, left, midTreeIndex, queryLeft, midTreeIndex),
		st.queryInTree(rightTreeIndex, midTreeIndex+1, right, midTreeIndex+1, queryRight))
}

func (st *SegmentTree) Update(index, val int) {
	if len(st.data) > 0 {
		st.updateInTree(0, 0, len(st.data)-1, index, val)
	}
}

// 以 treeIndex 为根，更新 index 位置上的值为 val
func (st *SegmentTree) updateInTree(treeIndex, left, right, index, val int) {
	if left == right {
		st.tree[treeIndex] = val
		return
	}
	midTreeIndex, leftTreeIndex, rightTreeIndex := left+(right-left)>>1, st.leftChild(treeIndex), st.rightChild(treeIndex)
	if index > midTreeIndex {
		st.updateInTree(rightTreeIndex, midTreeIndex+1, right, index, val)
	} else {
		st.updateInTree(leftTreeIndex, left, midTreeIndex, index, val)
	}
	st.tree[treeIndex] = st.merge(st.tree[leftTreeIndex], st.tree[rightTreeIndex])
}
