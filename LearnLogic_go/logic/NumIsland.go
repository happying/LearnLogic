package logic

import "fmt"

/*
给你一个由 '1'（陆地）和 '0'（水）组成的的二维网格，请你计算网格中岛屿的数量。

岛屿总是被水包围，并且每座岛屿只能由水平方向和/或竖直方向上相邻的陆地连接形成。

此外，你可以假设该网格的四条边均被水包围。

示例 1：

输入：grid = [
  {"1","1","1","1","0"},
  {"1","1","0","1","0"},
  {"1","1","0","0","0"},
  {"0","0","0","0","0"}
]
输出：1
示例 2：

输入：grid = [
  {"1","1","0","0","0"},
  {"1","1","0","0","0"},
  {"0","0","1","0","0"},
  {"0","0","0","1","1"}
]
输出：3



提示：
m == grid.length
n == grid[i].length
1 <= m, n <= 300
grid[i][j] 的值为 '0' 或 '1'

["1","0","1","1","1"]
["1","0","1","0","1"]
["1","1","1","0","1"]

["1","1","1","1","0"],
["1","1","0","1","0"],
["1","1","0","0","0"],
["0","0","0","0","0"]]

*/

func TestNumIslands() {

	a := [][]string{{"1", "1", "1"}, {"0", "1", "0"}, {"1", "1", "1"}}
	a = [][]string{{"1", "0", "1", "1", "1"}, {"1", "0", "1", "0", "1"}, {"1", "1", "1", "0", "1"}}
	a = [][]string{{"1", "1", "1", "1", "0"}, {"1", "1", "0", "1", "0"}, {"1", "1", "0", "0", "0"}, {"0", "0", "0", "0", "0"}}
	b := [][]byte{}
	for _, subStr := range a {
		subB := []byte{}
		for _, b2 := range subStr {
			subB = append(subB, []byte(b2)[0])
		}
		b = append(b, subB)
	}
	result := numIslands(b)
	fmt.Println(result)
}

func numIslands(grid [][]byte) int {
	iLength := len(grid)
	jLength := len(grid[0])

	var dfs func(i, j int)
	dfs = func(i, j int) {
		if i < 0 || j < 0 || i >= iLength || j >= jLength || grid[i][j] != '1' {
			return
		}
		grid[i][j] = '2'
		dfs(i+1, j)
		dfs(i-1, j)
		dfs(i, j+1)
		dfs(i, j-1)
	}

	totalCount := 0
	for i := 0; i < iLength; i++ {
		for j := 0; j < jLength; j++ {
			if grid[i][j] == '1' {
				totalCount++
				dfs(i, j)
			}
		}
	}

	return totalCount
}

func numIslands_v1(grid [][]byte) int {
	iLength := len(grid)
	jLength := len(grid[0])
	totalCacheMap := map[int]int{}
	totalCount := 0
	for i := 0; i < iLength; i++ {
		for j := 0; j < jLength; j++ {
			key := i*10000 + j
			if _, ok := totalCacheMap[key]; ok {
				continue
			}
			//  []byte("0")[0] = 48
			if grid[i][j] == 48 {
				totalCacheMap[key] = 0
				continue
			}
			totalCount++
			checkTop(grid, i, j, iLength, jLength, totalCacheMap)
			checkBottom(grid, i-1, j, iLength, jLength, totalCacheMap)
		}
	}

	return totalCount
}

func checkTop(grid [][]byte, i int, j int, iLen int, jLen int, cacheMap map[int]int) {
	if i < 0 || j < 0 {
		return
	}
	if i >= iLen || j >= jLen {
		return
	}
	key := i*10000 + j
	if _, ok := cacheMap[key]; ok {
		return
	}
	if grid[i][j] == 48 {
		cacheMap[key] = 0
		return
	}

	cacheMap[key] = 1
	checkTop(grid, i+1, j, iLen, jLen, cacheMap)
	checkRight(grid, i, j+1, iLen, jLen, cacheMap)
	checkLeft(grid, i, j-1, iLen, jLen, cacheMap)

	return
}

func checkBottom(grid [][]byte, i int, j int, iLen int, jLen int, cacheMap map[int]int) {
	if i < 0 || j < 0 {
		return
	}
	if i >= iLen || j >= jLen {
		return
	}
	key := i*10000 + j
	if _, ok := cacheMap[key]; ok {
		return
	}
	if grid[i][j] == 48 {
		cacheMap[key] = 0
		return
	}

	cacheMap[key] = 1
	checkBottom(grid, i-1, j, iLen, jLen, cacheMap)
	checkRight(grid, i, j+1, iLen, jLen, cacheMap)
	checkLeft(grid, i, j-1, iLen, jLen, cacheMap)

	return
}

func checkRight(grid [][]byte, i int, j int, iLen int, jLen int, cacheMap map[int]int) {
	if i < 0 || j < 0 {
		return
	}
	if i >= iLen || j >= jLen {
		return
	}
	key := i*10000 + j
	if _, ok := cacheMap[key]; ok {
		return
	}
	if grid[i][j] == 48 {
		cacheMap[key] = 0
		return
	}

	cacheMap[key] = 1
	checkTop(grid, i+1, j, iLen, jLen, cacheMap)
	checkRight(grid, i, j+1, iLen, jLen, cacheMap)
	checkBottom(grid, i-1, j, iLen, jLen, cacheMap)

	return
}

func checkLeft(grid [][]byte, i int, j int, iLen int, jLen int, cacheMap map[int]int) {
	if i < 0 || j < 0 {
		return
	}
	if i >= iLen || j >= jLen {
		return
	}
	key := i*10000 + j
	if _, ok := cacheMap[key]; ok {
		return
	}
	if grid[i][j] == 48 {
		cacheMap[key] = 0
		return
	}

	cacheMap[key] = 1
	checkTop(grid, i+1, j, iLen, jLen, cacheMap)
	checkLeft(grid, i, j-1, iLen, jLen, cacheMap)
	checkBottom(grid, i-1, j, iLen, jLen, cacheMap)
	return
}
