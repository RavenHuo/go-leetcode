package main

func rotate(nums []int, k int) {
	i := 0
	for i < k {
		// 先交换，再依次往前推
		n := nums[len(nums)-1]
		for k := len(nums) - 1; k > 0; k-- {
			nums[k] = nums[k-1]
		}
		nums[0] = n
		i++
	}
}
func main() {
	rotate([]int{1, 2, 3, 4, 5, 6, 7}, 3)
}
