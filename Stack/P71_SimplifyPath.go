package leetcode

import "strings"

func simplifyPath(path string) string {
	stack := []string{}
	folders := strings.Split(path, "/")
	for _, folder := range folders {
		if folder == "." || folder == "" {
		} else if folder == ".." {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		} else {
			stack = append(stack, folder)
		}
	}
	return "/" + strings.Join(stack, "/")
}
