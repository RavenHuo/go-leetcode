package main

func rotate(nums []int, k int) {
	if len(nums) == 0 || len(nums) == 1 {
		return
	}
	i := k%len(nums) + 1
	j := 0
	for i < len(nums) {
		n := nums[i]
		nums[i] = nums[j]
		nums[j] = n
		i++
		j++
	}
	rotate(nums[k:], 1)
}
func main() {
	rotate([]int{1, 2, 3, 4, 5, 6, 7}, 3)
}
