package main

import "fmt"

func find(nums []int, k int) int {
	prefix := make(map[int]int, len(nums))
	count := 0
	sum := 0
	for _, v := range nums {
		sum += v
		checkSum := sum - k
		count += prefix[checkSum]
		prefix[sum]++
	}
	return count
}

func main() {
	nums := []int{1, 2, 3, 2, 1}
	k := 2
	fmt.Println(find(nums, k))
}
