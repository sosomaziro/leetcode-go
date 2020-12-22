package main

func main() {
	sum := twoSum([]int{1, 2, 3}, 5)
	for i, v := range sum {
		println(i, " - ", v)
	}
}

func twoSum(nums []int, target int) []int {
	table := map[int]int{}
	for i, v := range nums {
		if p, k := table[target-v]; k {
			return []int{p, i}
		}
		table[v] = i
	}
	return nil
}
