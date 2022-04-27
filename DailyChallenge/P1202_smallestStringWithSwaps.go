package leetcode

import (
	"sort"
)

func smallestStringWithSwaps(s string, pairs [][]int) string {
	length := len(s)
	str := []byte(s)
	res := make([]byte, length)
	fa := make([]int, length)
	cos := make(map[int][]int, 0) // connect components
	for i := range fa {
		fa[i] = i
		cos[i] = []int{i}
	}
	for _, pair := range pairs {
		union(fa, pair[0], pair[1], cos)
	}

	for _, set := range cos {
		sort.Ints(set)
		cp := make([]byte, len(set))
		for j := range set {
			cp[j] = str[set[j]]
		}
		// fmt.Println(cp)
		sort.Slice(cp, func(i, j int) bool { return cp[i] < cp[j] })
		for i, index := range set {
			res[index] = cp[i]
			// fmt.Printf("res[%d]=str[%d]=%c\n", cp[i], index, str[index])
		}
	}
	return string(res)
}

func union(fa []int, a, b int, co map[int][]int) { // a<=b
	a = find(fa, a)
	b = find(fa, b)
	if a == b {
		return
	}
	fa[b] = a
	co[a] = append(co[a], co[b]...)
	delete(co, b)
}

func find(fa []int, a int) int {
	if fa[a] == a {
		return a
	}
	fa[a] = find(fa, fa[a])
	return fa[a]
}
