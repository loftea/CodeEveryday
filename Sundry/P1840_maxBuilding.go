package sundry

import (
	"math"
	"sort"
)

func maxBuilding(n int, restrictions [][]int) int {
	restrictions = append(restrictions, []int{1, 0})
	length := len(restrictions)
	sort.Slice(restrictions, func(i, j int) bool {
		return restrictions[i][0] < restrictions[j][0]
	}) //排序

	reveal(restrictions)
	res := toBe(restrictions, length)
	res = max(res, calcHeight(n, restrictions[length-1]))
	return res
}

func calcHeight(x int, f []int) int {
	return abs(x-f[0]) + f[1]
}

func reveal(A [][]int) {
	length := len(A)
	for i := 0; i < length; i++ {
		res := math.MaxInt
		for j := 0; j < length; j++ {
			res = min(res, calcHeight(A[i][0], A[j]))
		}
		A[i][1] = res
	}
}

func toBe(A [][]int, length int) int {
	if length == 1 {
		return A[0][1]
	}
	res := 0
	for i := 0; i < length-1; i++ {
		res = max(res, abs(A[i][1]-A[i][0]+A[i+1][1]+A[i+1][0])>>1)
	}
	return res
}

func abs(x int) int {
	if x > 0 {
		return x
	}
	return -1 * x
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
