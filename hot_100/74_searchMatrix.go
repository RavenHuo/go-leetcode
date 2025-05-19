package main

func searchMatrix(matrix [][]int, target int) bool {
	for _, nums := range matrix {
		if len(nums) <= 0 {
			continue
		}
		if nums[0] <= target && nums[len(nums)-1] >= target {
			i := 0
			j := len(nums) - 1
			for i <= j {
				if nums[i] == target || nums[j] == target {
					return true
				}
				if nums[i] < target {
					i++
				}
				if nums[j] > target {
					j--
				}
			}
		}
	}
	return false
}

func main() {
	searchMatrix([][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, 3)
}
