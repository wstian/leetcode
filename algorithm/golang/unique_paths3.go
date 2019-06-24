package main

const _UNVISITED = 0
const _STARTING = 1
const _ENDING = 2
const _BLOCKED = -1
const _VISITED = -2


func uniquePathsIII(grid [][]int) int {
	start_i := -1
	start_j := -1
	for i := range(grid) {
		for j := range(grid[i]) {
			if grid[i][j] == _STARTING {
				start_i = i;
				start_j = j;
			}
		}
	}
	return _uniquePaths(grid, start_i, start_j)
}

func _uniquePaths(grid [][]int, i int, j int) int {
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) || grid[i][j] == _BLOCKED || grid[i][j] == _VISITED {
		return 0
	}
	if grid[i][j] == _ENDING {
		return _check(grid)
	}
	grid[i][j] = _VISITED
	cnt := 0
	cnt += _uniquePaths(grid, i - 1, j)
	cnt += _uniquePaths(grid, i + 1, j)
	cnt += _uniquePaths(grid, i, j - 1)
	cnt += _uniquePaths(grid, i, j + 1)
	grid[i][j] = _UNVISITED
	return cnt
}

func _check(grid [][]int) int {
	for i := range(grid) {
		for j := range(grid[i]) {
			if grid[i][j] == _UNVISITED {
				return 0
			}
		}
	}
	return 1
}
