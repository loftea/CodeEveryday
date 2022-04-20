func maxBuilding(n int, restrictions [][]int) int {
	// no restrictions, increase one by one
	if len(restrictions) == 0 {
		return n - 1
	}
	// sort by index
	sort.Slice(restrictions, func(i, j int) bool {
		return restrictions[i][0] < restrictions[j][0]
	})
	// make height valid
	for _, restrict := range restrictions {
		if restrict[1] >= restrict[0] {
			restrict[1] = restrict[0] - 1
		}
	}
	// Heap item: [indexOfRestrictions, indexOfBuildings, restrictHeight]
	h := &Heap{}
	for i, v := range restrictions {
		heap.Push(h, []int{i, v[0], v[1]})
	}
	processed := make([]bool, len(restrictions))
	for h.Len() != 0 {
		elem := heap.Pop(h).([]int)
		if processed[elem[0]] == true {
			continue
		}
		processed[elem[0]] = true
		// refresh restrictions
		restrictions[elem[0]][1] = elem[2]
		if elem[0] > 0 && processed[elem[0]-1] == false {
			leftElem := restrictions[elem[0]-1]
			maxLeftElemHeight := elem[2] + elem[1] - leftElem[0]
			if leftElem[1] > maxLeftElemHeight {
				leftElem[1] = maxLeftElemHeight
			}
			heap.Push(h, []int{elem[0] - 1, leftElem[0], leftElem[1]})
		}
		if elem[0] < len(restrictions)-1 && processed[elem[0]+1] == false {
			rightElem := restrictions[elem[0]+1]
			maxRightElemHeight := elem[2] + rightElem[0] - elem[1]
			if rightElem[1] > maxRightElemHeight {
				rightElem[1] = maxRightElemHeight
			}
			heap.Push(h, []int{elem[0] + 1, rightElem[0], rightElem[1]})
		}
	}
	// add restrict for last building if not exist
	if restrictions[len(restrictions)-1][0] != n {
		restrictions = append(restrictions, []int{n, restrictions[len(restrictions)-1][1] + n - restrictions[len(restrictions)-1][0]})
	}
	// calculate max building height
	maxHeight := 0
	prevIndex, preHeight := 1, 0
	for _, restrict := range restrictions {
		if restrict[1] != preHeight+restrict[0]-prevIndex && restrict[1] != preHeight-restrict[0]+prevIndex {

		}
		if curHeight := getMaxHeight(preHeight, restrict[1], restrict[0]-prevIndex); curHeight > maxHeight {
			maxHeight = curHeight
		}
		prevIndex, preHeight = restrict[0], restrict[1]
	}
	return maxHeight
}

func getMaxHeight(left, right, gap int) int {
	if left < right {
		gap -= right - left
		left = right
	} else if left > right {
		gap -= left - right
		right = left
	}
	return left + gap/2
}

type Heap [][]int

func (p Heap) Len() int            { return len(p) }
func (p Heap) Less(i, j int) bool  { return p[i][2] < p[j][2] }
func (p Heap) Swap(i, j int)       { p[i], p[j] = p[j], p[i] }
func (p *Heap) Push(i interface{}) { *p = append(*p, i.([]int)) }
func (p *Heap) Pop() interface{}   { v := (*p)[len(*p)-1]; *p = (*p)[:len(*p)-1]; return v }