package leetcode

import (
	"sort"
)

func accountsMerge(accounts [][]string) [][]string {
	var masterGroups = make(map[*group]bool, 0)
	fa := make(map[string]*group, 0)

	for _, account := range accounts {
		if _, ok := fa[account[1]]; !ok {
			fa[account[1]] = &group{name: account[0], father: nil}
			masterGroups[fa[account[1]]] = true
		}

		for i := 2; i < len(account); i++ {
			if _, ok := fa[account[i]]; ok {
				union(fa[account[i]], fa[account[1]], masterGroups)
			}
			fa[account[i]] = fa[account[1]]
		}
	}

	res := make(map[*group][]string, 0)

	for email, father := range fa {
		a := find(father)
		res[a] = append(res[a], email)
	}

	output := [][]string{}

	for gp := range masterGroups {
		line := []string{gp.name}
		sort.Strings(res[gp])
		line = append(line, res[gp]...)
		output = append(output, line)
	}

	return output

}

type group struct {
	father *group
	name   string
}

func find(x *group) *group {
	if x.father != nil {
		x = find(x.father)
	}
	return x
}

func union(x, y *group, masterGroups map[*group]bool) {
	x = find(x)
	y = find(y)
	if x == y {
		return
	} else {
		y.father = x
		delete(masterGroups, y)
	}
}
