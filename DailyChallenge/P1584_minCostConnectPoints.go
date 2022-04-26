func minCostConnectPoints(points [][]int) int {
	h := &minheap{}
	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			heap.Push(h, [3]int{i, j, distance(points[i], points[j])})
		}
	}

	visited := make(map[int]bool)
	res := 0

	var list [][3]int
	for h.Len() > 0 {
		e := heap.Pop(h).([3]int)
		if visited[e[0]] && visited[e[1]] {
			continue
		}
		if res != 0 && !visited[e[0]] && !visited[e[1]] {
			list = append(list, e)
			continue
		}
		res += e[2]
		visited[e[0]] = true
		visited[e[1]] = true
		for _, v := range list {
			heap.Push(h, v)
		}
		list = [][3]int{}
	}
	return res
}

func distance(a []int, b []int) int {
	return abs(a[0]-b[0]) + abs(a[1]-b[1])
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

type minheap [][3]int

func (h minheap) Len() int {
	return len(h)
}

func (h minheap) Less(i int, j int) bool {
	return h[i][2] < h[j][2]
}

func (h minheap) Swap(i int, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *minheap) Push(a interface{}) {
	*h = append(*h, a.([3]int))
}

func (h *minheap) Pop() interface{} {
	l := len(*h)
	res := (*h)[l-1]
	*h = (*h)[:l-1]
	return res
}