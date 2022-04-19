type RangeModule struct {
	Root *BSTNode
}

type BSTNode struct {
	Interval    []int
	Left, Right *BSTNode
}

func Constructor() RangeModule {
	return RangeModule{}
}

func (this *RangeModule) AddRange(left int, right int) {
	interval := []int{left, right - 1}
	this.Root = insert(this.Root, interval)
}

func (this *RangeModule) RemoveRange(left int, right int) {
	interval := []int{left, right - 1}
	this.Root = delete(this.Root, interval)
}

func (this *RangeModule) QueryRange(left int, right int) bool {
	return query(this.Root, []int{left, right - 1})
}

func insert(root *BSTNode, interval []int) *BSTNode {
	if root == nil {
		return &BSTNode{interval, nil, nil}
	}
	if root.Interval[0] <= interval[0] && interval[1] <= root.Interval[1] {
		return root
	}
	if interval[0] < root.Interval[0] {
		root.Left = insert(root.Left, []int{interval[0], min(interval[1], root.Interval[0]-1)})
	}
	if root.Interval[1] < interval[1] {
		root.Right = insert(root.Right, []int{max(interval[0], root.Interval[1]+1), interval[1]})
	}
	return root
}

func delete(root *BSTNode, interval []int) *BSTNode {
	if root == nil {
		return nil
	}
	if interval[0] < root.Interval[0] {
		root.Left = delete(root.Left, []int{interval[0], min(interval[1], root.Interval[0]-1)})
	}
	if root.Interval[1] < interval[1] {
		root.Right = delete(root.Right, []int{max(interval[0], root.Interval[1]+1), interval[1]})
	}
	if interval[1] < root.Interval[0] || root.Interval[1] < interval[0] {
		return root
	}
	if interval[0] <= root.Interval[0] && root.Interval[1] <= interval[1] {
		if root.Left == nil {
			return root.Right
		} else if root.Right == nil {
			return root.Left
		} else {
			pred := root.Left
			for pred.Right != nil {
				pred = pred.Right
			}
			root.Interval = pred.Interval
			root.Left = delete(root.Left, pred.Interval)
			return root
		}
	}
	if root.Interval[0] < interval[0] && interval[1] < root.Interval[1] {
		left := &BSTNode{[]int{root.Interval[0], interval[0] - 1}, root.Left, nil}
		right := &BSTNode{[]int{interval[1] + 1, root.Interval[1]}, nil, root.Right}
		left.Right = right
		return left
	}
	if interval[0] <= root.Interval[0] {
		root.Interval[0] = interval[1] + 1
	}
	if root.Interval[1] <= interval[1] {
		root.Interval[1] = interval[0] - 1
	}
	return root
}

func query(root *BSTNode, interval []int) bool {
	if root == nil {
		return false
	}
	if interval[1] < root.Interval[0] {
		return query(root.Left, interval)
	}
	if root.Interval[1] < interval[0] {
		return query(root.Right, interval)
	}
	left := true
	if interval[0] < root.Interval[0] {
		left = query(root.Left, []int{interval[0], root.Interval[0] - 1})
	}
	right := true
	if root.Interval[1] < interval[1] {
		right = query(root.Right, []int{root.Interval[1] + 1, interval[1]})
	}
	return left && right
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

/**
 * Your RangeModule object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddRange(left,right);
 * param_2 := obj.QueryRange(left,right);
 * obj.RemoveRange(left,right);
 */