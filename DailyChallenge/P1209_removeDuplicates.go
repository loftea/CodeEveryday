package main

import "fmt"

func removeDuplicates(s string, k int) string {
	str := []byte(s)
	for !removeLoop(&str, k) {
	}
	return string(str)
}

func removeLoop(s *[]byte, k int) bool {
	l, r, count := 0, 0, 0
	length := len(*s)
	for l < length {
		count = 1
		for r < length-1 && (*s)[r] == (*s)[r+1] {
			count++
			r++
		}
		if count >= k {
			*s = append((*s)[:(l+(count%k))], (*s)[r+1:]...)
			return false
		}
		l = r + 1
		r++
	}
	return true
}

func main() {
	s := "deeedbbcccbdaa"
	k := 3
	fmt.Println(removeDuplicates(s, k))
}
