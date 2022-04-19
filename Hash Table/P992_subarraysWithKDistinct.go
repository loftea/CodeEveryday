package leetcode

// import "fmt"

func subarraysWithKDistinct(A []int, k int) int {
	length := len(A)
	left, mid, right, count, res, H := 0, 0, 0, 0, 0, make([]int, length+1)
	for ; right < length; right++ {
		x := A[right]

		if H[x] <= left { // x not in window
			count++
		} else if H[x] == mid+1 {
			mid++
			for mid+1 < H[A[mid]] {
				mid++
			}
		}

		switch count {
		case k + 1:
			left = mid + 1
			mid++
			for mid+1 < H[A[mid]] {
				mid++
			}
			count--
			fallthrough

		case k:
			res += mid - left + 1
		}

		H[x] = right + 1
		// fmt.Println(left,mid,right)
	}

	return res
}
