package backtracking

type solve struct {
	path   []int
	dim    int
	res    [][]int
	tables [][]string
}

func (s *solve) check(x int) bool {
	last := len(s.path)
	for i := 0; i < last; i++ {
		switch abs(x - s.path[i]) {
		case 0:
			return false
		case (last - i):
			return false
		}
	}
	return true
}

func (s *solve) step() bool {
	length := len(s.path)
	if length == s.dim {
		s.add()
		return s.back()
	} else {
		for i := 0; i < s.dim; i++ {
			if s.check(i) {
				s.path = append(s.path, i)
				return false
			}
		}
		return s.back()
	}
}

func (s *solve) add() {
	cpy := make([]int, len(s.path))
	copy(cpy, s.path)
	s.res = append(s.res, cpy)
}

func (s *solve) back() bool { //回退查找字典序下一个可行的path，如果找不到则说明查找结束
	if len(s.path) == 0 {
		return true
	}
	tmp := s.path[len(s.path)-1]
	s.path = s.path[:len(s.path)-1] //删除一个元素

	for i := tmp + 1; i < s.dim; i++ {
		if s.check(i) {
			s.path = append(s.path, i)
			return false
		}
	}

	return s.back()
}

func (s *solve) makeTables() {
	for _, x := range s.res {
		table := []string{}
		for i := 0; i < s.dim; i++ {
			line := ""
			for j := 0; j < s.dim; j++ {
				if x[i] == j {
					line += "Q"
				} else {
					line += "."
				}
			}
			table = append(table, line)
		}
		s.tables = append(s.tables, table)
	}
}

// func (s *solve) ShowPath() {
// 	fmt.Printf("path = %v now\n", s.path)
// }

// func (s *solve) ShowRes() {
// 	fmt.Printf("res = %v now\n", s.res)
// }

func NewSolve(n int) solve {
	return solve{dim: n}
}

func abs(x int) int {
	if x > 0 {
		return x
	}
	return -1 * x
}

func solveNQueens(n int) [][]string {
	s := NewSolve(n)
	for !s.step() {
	}
	s.makeTables()
	return s.tables
}

// func main() {
// 	solveNQueens(4)
// }
