package main

func moveZeroes(nums []int) {
	if len(nums) <= 1 {
		return
	}
	i := 0
	j := i + 1
	for i < len(nums) && j < len(nums) {
		if nums[i] == 0 {
			if nums[j] == 0 {
				j++
			} else { // 交换
				n := nums[i]
				m := nums[j]
				nums[i] = m
				nums[j] = n
				// 交换之后全部都右移
				i++
				j++
			}
		} else {
			// 不是0的话得全都都右移
			i++
			j++
		}
	}
}

func main() {
	moveZeroes([]int{0, 1, 0, 3, 12})
}
