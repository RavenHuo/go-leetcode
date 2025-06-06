package main

var globalGrid [][]byte

func numIslands(grid [][]byte) int {
	globalGrid = grid
	if len(globalGrid) == 0 {
		return 0
	}
	res := 0
	for i := 0; i < len(globalGrid); i++ {
		for j := 0; j < len(globalGrid[i]); j++ {
			if globalGrid[i][j] == '1' {
				res += 1
				numIslandsDfs(i, j)
			}
		}
	}
	return res
}
func numIslandsDfs(i, j int) {
	if i < 0 || i > len(globalGrid)-1 || j < 0 || j > len(globalGrid[0])-1 || globalGrid[i][j] != 1 {
		return
	}
	globalGrid[i][j] = '2'
	numIslandsDfs(i-1, j)
	numIslandsDfs(i+1, j)
	numIslandsDfs(i, j+1)
}

func main() {
	numIslands([][]byte{[]byte{'1', '1', '1', '1', '0'}})
}
