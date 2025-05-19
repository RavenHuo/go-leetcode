package main

func productExceptSelf(nums []int) []int {
	// 左边的结果
	lRes := make([]int, len(nums))
	// 右边的结果
	rRes := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		if i > 0 {
			lRes[i] = nums[i] * lRes[i-1]
		} else {
			lRes[i] = nums[i]
		}
	}
	for i := len(nums) - 1; i >= 0; i-- {
		if i < len(nums)-1 {
			rRes[i] = nums[i] * rRes[i+1]
		} else {
			rRes[i] = nums[i]
		}
	}
	resultList := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		l := 1
		r := 1
		if i-1 >= 0 {
			l = lRes[i-1]
		}
		if i+1 <= len(nums)-1 {
			r = rRes[i+1]
		}
		resultList[i] = l * r
	}
	return resultList
}

func main() {
	productExceptSelf([]int{0, 0})
}
