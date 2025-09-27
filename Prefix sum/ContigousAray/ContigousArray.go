package main

import "fmt"

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
func findMaxLength(nums []int) int {
	prefix := map[int]int{0: -1}
	count := 0
	maxLength := 0
	for i, n := range nums {
		if n == 1 {
			count++
		} else {
			count--
		}
		if _, ok := prefix[count]; ok {
			maxLength = max(maxLength, i-prefix[count])
		} else {
			prefix[count] = i
		}
	}
	return maxLength
}
func main() {
	nums := []int{0, 0, 1, 1, 1, 0, 0, 0}
	fmt.Println(findMaxLength(nums))
}
