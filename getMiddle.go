package main

func main() {
	println(findMedianSortedArrays([]int{1, 2, 3, 4, 7, 8, 9, 10}, []int{5, 6}))
	println(findMedianSortedArrays([]int{1, 2}, []int{3, 4}))
	println(findMedianSortedArrays([]int{1, 3}, []int{2}))
	println(findMedianSortedArrays([]int{1}, []int{1}))
}

func findMedianSortedArrays(nums1, nums2 []int) float64 {
	l1 := len(nums1)
	l2 := len(nums2)
	l := l1 + l2
	if l%2 == 1 {
		return float64(getMedNum(nums1, nums2, l/2+1))
	} else {
		return float64(getMedNum(nums1, nums2, l/2)+getMedNum(nums1, nums2, l/2+1)) / 2.0
	}
}

func getMedNum(n1, n2 []int, m int) int {
	i1, i2 := 0, 0
	for {
		if i1 == len(n1) {
			return n2[m+i2-1]
		}
		if i2 == len(n2) {
			return n1[m+i1-1]
		}
		if m == 1 {
			return minimum(n1[i1], n2[i2])
		}

		half := m / 2
		ni1 := minimum(half+i1, len(n1)) - 1
		ni2 := minimum(half+i2, len(n2)) - 1
		c1, c2 := n1[ni1], n2[ni2]
		if c1 > c2 {
			m = m - (ni2 - i2 + 1)
			i2 = ni2 + 1
		} else {
			m = m - (ni1 - i1 + 1)
			i1 = ni1 + 1
		}

	}
}

func minimum(a int, b int) int {
	if a > b {
		return b
	}
	return a
}
