package main

func searchRange(nums []int, target int) []int {
	res := make([]int, 2)
	res[0] = -1
	res[0] = -1
	i := 0
	j := len(nums) - 1
	for i <= j {
		if nums[i] < target {
			i++
		}
		if nums[j] > target {
			j--
		}
		if nums[i] == target && nums[j] == target {
			res[0] = i
			res[1] = j
			return res
		}
	}
	return res
}
func main() {
	searchRange([]int{5, 7, 7, 8, 8, 10}, 8)
}
