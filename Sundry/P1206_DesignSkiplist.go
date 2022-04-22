import (
	"math"
	"math/rand"
)

const (
	MaxLevel = 32
	P        = 1 / math.E
)

type Skiplist struct {
	head  *Node
	level int
}

func Constructor() Skiplist {
	return Skiplist{
		head:  newNode(0, 32),
		level: 1,
	}
}

type Node struct {
	score int
	next  []*Node
}

func newNode(score int, level int) *Node {
	return &Node{
		score: score,
		next:  make([]*Node, level),
	}
}

func randomLevel() int {
	level := 1
	for rand.Float64() < P {
		level++
	}
	if level > MaxLevel {
		return MaxLevel
	}
	return level
}

func (s *Skiplist) Search(target int) bool {
	x := s.head
	for i := s.level - 1; i >= 0; i-- {
		for x.next[i] != nil {
			if x.next[i].score == target {
				return true
			} else if x.next[i].score < target {
				x = x.next[i]
			} else {
				break
			}
		}
	}
	return false
}

func (s *Skiplist) Add(score int) {
	update := make([]*Node, MaxLevel)

	for i, x := s.level-1, s.head; i >= 0; i-- {
		for x.next[i] != nil && x.next[i].score < score {
			x = x.next[i]
		}
		update[i] = x
	}

	lvl := randomLevel()
	n := newNode(score, lvl)
	if lvl > s.level {
		for i := s.level; i < lvl; i++ {
			update[i] = s.head
		}
		s.level = lvl
	}

	for i := 0; i < lvl; i++ {
		n.next[i] = update[i].next[i]
		update[i].next[i] = n
	}
}

func (s *Skiplist) Erase(score int) bool {
	update := make([]*Node, MaxLevel)

	for i, x := s.level-1, s.head; i >= 0; i-- {
		for x.next[i] != nil && x.next[i].score < score {
			x = x.next[i]
		}
		update[i] = x
	}

	node := update[0].next[0]
	if node == nil || node.score != score {
		return false
	}

	for i := 0; i < len(node.next); i++ {
		update[i].next[i] = node.next[i]
	}

	for s.level > 1 && s.head.next[s.level-1] == nil {
		s.level--
	}

	return true
}