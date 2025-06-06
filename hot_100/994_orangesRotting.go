package main

var orangesRottingGrid [][]int
var orangesRottingRes int

func orangesRotting(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}
	orangesRottingGrid = grid
	orangesRottingRes = 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 2 {
				orangesRottingDfs(i, j)
				break
			}
		}
	}
	return orangesRottingRes
}

func orangesRottingDfs(i, j int) {
	if i < 0 || i > len(orangesRottingGrid)-1 || j < 0 || j > len(orangesRottingGrid[0])-1 || orangesRottingGrid[i][j] == 0 {
		return
	}
	orangesRottingDfs(i-1, j)
	orangesRottingDfs(i+1, j)
	orangesRottingDfs(i, j-1)
	orangesRottingDfs(i, j+1)
	if orangesRottingGrid[i][j] == 1 {
		orangesRottingGrid[i][j] = 2
		orangesRottingRes++
	}
}
