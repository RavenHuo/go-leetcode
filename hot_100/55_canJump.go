package main

// dfs(i)表示为i时可以跳的最远距离
// dfs(i) = max(dfs(i-1)-1,nums[i])
// 上一个最远能跳dfs(i-1)，跳到当前，减1
func canJump(nums []int) bool {
	maxPlace := 0
	for i := 0; i < len(nums); i++ {
		if maxPlace >= i {
			maxPlace = canJumpMax(maxPlace, i+nums[i])
		} else {
			return false
		}
	}
	return true
}

func canJumpMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	canJump([]int{3, 0, 8, 2, 0, 0, 1})
}
