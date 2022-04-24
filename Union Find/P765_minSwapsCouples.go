package leetcode

func minSwapsCouples(row []int) int {
	return solve(row)
}

func solve(row []int) int {
	if len(row) == 0 {
		return 0
	}
	a := row[0]

	b := coupleOf(a)

	if row[1] == b {
		return solve(row[2:])
	} else {
		for i := 2; i < len(row); i++ {
			if row[i] == b {
				row[i] = row[1]
				subrow := row[2:]
				return solve(subrow) + 1
			}
		}
	}
	return 0
}

// return the couple of a
func coupleOf(a int) int {
	if a%2 == 0 {
		return a + 1
	}
	return a - 1
}
