package main

import "fmt"

// Định nghĩa struct NumArr để lưu mảng prefix sum
type NumArr struct {
	prefix []int // prefix[i] = tổng từ nums[0..i-1]
}

// Constructor: tạo đối tượng NumArr từ mảng nums
func Constructor(nums []int) NumArr {
	// Cấp phát mảng prefix có độ dài = len(nums)+1
	// Thêm 1 phần tử để prefix[0] = 0 (giúp tính nhanh subarray sum)
	prefix := make([]int, len(nums)+1)

	prefix[0] = 0 // khởi tạo phần tử đầu tiên bằng 0

	// Tính prefix sum: prefix[i+1] = prefix[i] + nums[i]
	for i := 0; i < len(nums); i++ {
		prefix[i+1] = prefix[i] + nums[i]
	}

	// Trả về một struct NumArr với field prefix
	return NumArr{prefix}
}

// Method SumArr: tính tổng từ nums[left..right]
// Sử dụng con trỏ *NumArr để tránh copy struct lớn
func (variable *NumArr) SumArr(left int, right int) int {
	// Công thức: sum[left..right] = prefix[right+1] - prefix[left]
	return variable.prefix[right+1] - variable.prefix[left]
}

func main() {
	// Mảng đầu vào
	nums := []int{-2, 0, 3, 2, 1}

	// Gọi Constructor để tạo đối tượng NumArr với prefix sum
	object := Constructor(nums)

	// In ra tổng từ nums[0..4] = -2 + 0 + 3 + 2 + 1 = 4
	fmt.Println(object.SumArr(0, 4)) // Kết quả: 4
}
