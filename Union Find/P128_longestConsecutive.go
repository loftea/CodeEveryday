package leetcode

type node struct {
	father      *node
	son         *node
	is_Ancestor bool
}

func (n *node) count() int {
	if n.son == nil {
		return 1
	}
	return n.son.count() + 1
}

func (n *node) updateFather(father *node) {
	n.father = father
	n.is_Ancestor = false
}

func (n *node) updateSon(son *node) {
	n.son = son
}

func longestConsecutive(nums []int) int {
	union := map[int]*node{}
	res := 0

	for _, val := range nums {
		union[val] = &node{father: nil, is_Ancestor: true, son: nil}
	}

	for _, val := range nums {
		if _, isPresent := union[val+1]; isPresent {
			union[val].updateFather(union[val+1])
			union[val+1].updateSon(union[val])
		}
	}

	for _, val := range nums {
		if union[val].is_Ancestor {
			res = max(res, union[val].count())
		}
	}

	return res

}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
