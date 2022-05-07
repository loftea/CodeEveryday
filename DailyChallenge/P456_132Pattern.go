// package dailychallenge
package main

import "fmt"

type atom struct {
	index int
	val   int
}

func find132pattern(nums []int) bool {
	a := make([]atom, len(nums))
	b := make([]atom, len(nums))

	for i, num := range nums {
		if i == 0 {
			a[i] = atom{index: 0, val: num}
			b[i] = atom{index: 0, val: num}
		} else {
			if b[i-1].index > a[i-1].index &&
				b[i-1].val > num && num > a[i-1].val {
				return true
			}

			if a[i-1].val <= num {
				a[i] = a[i-1]
			} else {
				a[i].val = num
				a[i].index = i
			}

			if b[i-1].val > num {
				b[i] = b[i-1]
			} else {
				b[i].val = num
				b[i].index = i
			}
		}
	}
	return false
}

func main() {
	nums := []int{3, 5, 0, 3, 4}
	fmt.Println(find132pattern(nums))
}
