package logic

import (
	"encoding/json"
	"fmt"
)

/*
给你一个大小为 m x n 的二进制网格 grid 。网格表示一个地图，其中，0 表示水，1 表示陆地。最初，grid 中的所有单元格都是水单元格（即，所有单元格都是 0）。

可以通过执行 addLand 操作，将某个位置的水转换成陆地。给你一个数组 positions ，其中 positions[i] = [ri, ci] 是要执行第 i 次操作的位置 (ri, ci) 。

返回一个整数数组 answer ，其中 answer[i] 是将单元格 (ri, ci) 转换为陆地后，地图中岛屿的数量。

岛屿 的定义是被「水」包围的「陆地」，通过水平方向或者垂直方向上相邻的陆地连接而成。你可以假设地图网格的四边均被无边无际的「水」所包围。


输入：m = 3, n = 3, positions = [[0,0],[0,1],[1,2],[2,1]]
输出：[1,1,2,3]
解释：
起初，二维网格 grid 被全部注入「水」。（0 代表「水」，1 代表「陆地」）
- 操作 #1：addLand(0, 0) 将 grid[0][0] 的水变为陆地。此时存在 1 个岛屿。
- 操作 #2：addLand(0, 1) 将 grid[0][1] 的水变为陆地。此时存在 1 个岛屿。
- 操作 #3：addLand(1, 2) 将 grid[1][2] 的水变为陆地。此时存在 2 个岛屿。
- 操作 #4：addLand(2, 1) 将 grid[2][1] 的水变为陆地。此时存在 3 个岛屿。
示例 2：

输入：m = 1, n = 1, positions = [[0,0]]
输出：[1]


提示：

1 <= m, n, positions.length <= 10^4
1 <= m * n <= 10^4
positions[i].length == 2
0 <= ri < m
0 <= ci < n


进阶：你可以设计一个时间复杂度 O(k log(mn)) 的算法解决此问题吗？（其中 k == positions.length）
*/

type UnionFind struct {
	Count   int
	parents []int
	weight  []int
}

func NewUF(num int) *UnionFind {
	uf := &UnionFind{}
	uf.parents = make([]int, num)
	uf.weight = make([]int, num)
	for i := 0; i < num; i++ {
		uf.parents[i] = i
		uf.weight[i] = 1
	}

	return uf
}

func (uf *UnionFind) find(index int) int {
	if uf.parents[index] != index {
		index = uf.parents[index]
		uf.parents[index] = uf.find(uf.parents[index])
	}
	return index
}

func (uf *UnionFind) union(a, b int) {
	roota, rootb := uf.find(a), uf.find(b)
	if roota == rootb {
		return
	}
	weightA := uf.weight[a]
	weightB := uf.weight[b]
	if weightA < weightB {
		uf.parents[roota] = rootb
		uf.weight[rootb] += uf.weight[roota]
	} else {
		uf.parents[rootb] = roota
		uf.weight[roota] += uf.weight[rootb]
	}
	uf.Count--
}

func (uf *UnionFind) IsConnected(a, b int) bool {
	return uf.find(a) == uf.find(b)
}

func TestNumIslands2() {
	a := [][]int{}
	//a = [][]int{{0, 0}, {0, 1}, {1, 2}, {2, 1}}
	jsonStr := "[[0,1],[1,2],[2,1],[1,0],[0,2],[0,0],[1,1]]" // 对应结果：[1,2,3,4,3,2,1]

	//jsonStr = "[[0,0],[0,1],[1,2],[1,2]]"

	err := json.Unmarshal([]byte(jsonStr), &a)
	if err != nil {
		fmt.Println(err)
	}

	result := numIslands2(3, 3, a)
	fmt.Println(result)
}

func numIslands2(m int, n int, positions [][]int) []int {

	uf := NewUF(m * n)
	visited := make([]bool, m*n)
	directions := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	result := []int{}

	isInLand := func(x, y int) bool {
		if x >= 0 && y >= 0 && x < m && y < n {
			return true
		}
		return false
	}

	for _, position := range positions {
		x := position[0]
		y := position[1]
		index := x*n + y
		//if !isInLand(x, y) {
		//	continue
		//}
		if visited[index] {
			result = append(result, uf.Count)
			continue
		}
		visited[index] = true
		uf.Count++
		for _, direction := range directions {
			newX := x + direction[0]
			newY := y + direction[1]
			newIndex := newX*n + newY
			if isInLand(newX, newY) && visited[newIndex] && !uf.IsConnected(index, newIndex) {
				uf.union(index, newIndex)
			}
		}
		result = append(result, uf.Count)
	}

	return result
}

////////////

func numIslands2_v2(m int, n int, positions [][]int) []int {
	sea := generateArray(m, n)
	islandCont := []int{}
	islandToBorder := map[int]map[int]int{}
	borderToIsland := map[int]map[int]int{}
	tmpAAA := 10000
	for _, position := range positions {
		mIndex := position[0]
		nIndex := position[1]

		oldSeaArea := sea[mIndex][nIndex]
		if oldSeaArea == 1 {
			// 本来就是岛屿，没必要操作
			islandCont = append(islandCont, len(islandToBorder))

			continue
		} else if oldSeaArea == 0 {
			// 本来是海洋的，直接新增一个岛屿
			islandKey := mIndex*tmpAAA + nIndex
			// 新增的岛屿，上下左右属于它的边界
			borders := handleBorders(m, n, mIndex, nIndex, sea, tmpAAA, borderToIsland, islandKey)
			islandToBorder[islandKey] = borders
			sea[mIndex][nIndex] = 1
			islandCont = append(islandCont, len(islandToBorder))

			continue
		}
		// 到了这里，意味着新增的陆地肯定属于某个岛屿，需要做相应的合并操作处理
		borderToIslandKey := mIndex*tmpAAA + nIndex

		// 找到边界所勾连的所有岛屿，这些岛屿将要合并为一个
		islandKeys := map[int]int{}
		firstIslandKey := 0
		if islands, ok := borderToIsland[borderToIslandKey]; ok {
			for island, _ := range islands {
				islandM := island / tmpAAA
				islandN := island % tmpAAA
				islandKey := islandM*tmpAAA + islandN
				islandKeys[islandKey] = 1
				if firstIslandKey == 0 {
					firstIslandKey = islandKey
				}
			}
		}

		// 如果就一个，那也不需要进行合并了
		if len(islandKeys) == 1 {
			// 这个边界已经变成岛屿的一部分了，将它从边界map里移除
			delete(borderToIsland, borderToIslandKey)
			// 检查并新增上下左右4个边界
			borders := handleBorders(m, n, mIndex, nIndex, sea, tmpAAA, borderToIsland, firstIslandKey)
			islandToBorder[firstIslandKey] = borders
			sea[mIndex][nIndex] = 1
			islandCont = append(islandCont, len(islandToBorder))

			continue
		}
		sea[mIndex][nIndex] = 1

		// 涉及到岛屿合并了,这里使用任意一个岛屿的初坐标作为合并后岛屿的坐标都是没问题的，直接采用firstIslandKey

		// 找到所有需要合并岛屿的所有边界
		allCombineIslandBorders := map[int]int{}
		for key, _ := range islandKeys {
			borders := islandToBorder[key]
			for borderKey, _ := range borders {
				allCombineIslandBorders[borderKey] = 1
			}
			// 移除已经被合并的岛屿记录
			if key != firstIslandKey {
				delete(islandToBorder, key)
			}
		}

		islandToBorder[firstIslandKey] = allCombineIslandBorders

		// 检查所有边界，将原有的需要合并的岛屿，替换成合并后的岛屿
		for _, singleBorderIslandMap := range borderToIsland {
			for islandKey22, _ := range islandKeys {
				if islandKey22 == firstIslandKey {
					continue
				}
				if _, ok := singleBorderIslandMap[islandKey22]; ok {
					delete(singleBorderIslandMap, islandKey22)
					singleBorderIslandMap[firstIslandKey] = 1
				}
			}
		}

		borders := handleBorders(m, n, mIndex, nIndex, sea, tmpAAA, borderToIsland, firstIslandKey)
		islandToBorder[firstIslandKey] = borders
		sea[mIndex][nIndex] = 1
		islandCont = append(islandCont, len(islandToBorder))
	}

	return islandCont
}

func handleBorders(m int, n int, mIndex int, nIndex int, sea [][]int, tmpAAA int, borderToIsland map[int]map[int]int, islandKey int) map[int]int {
	borders := map[int]int{}
	if mIndex > 0 {
		tmpM := mIndex - 1
		tmpN := nIndex
		if sea[tmpM][tmpN] != 1 {
			sea[tmpM][tmpN] = sea[tmpM][tmpN] - 1
			borderKey := tmpM*tmpAAA + tmpN
			borders[borderKey] = 1
			if borderToIsland[borderKey] == nil {
				borderToIsland[borderKey] = map[int]int{}
			}
			borderToIsland[borderKey][islandKey] = 1
		}
	}
	if mIndex+1 < m {
		tmpM := mIndex + 1
		tmpN := nIndex
		if sea[tmpM][tmpN] != 1 {
			sea[tmpM][tmpN] = sea[tmpM][tmpN] - 1
			borderKey := tmpM*tmpAAA + tmpN
			borders[borderKey] = 1
			if borderToIsland[borderKey] == nil {
				borderToIsland[borderKey] = map[int]int{}
			}
			borderToIsland[borderKey][islandKey] = 1
		}
	}

	if nIndex > 0 {
		tmpM := mIndex
		tmpN := nIndex - 1
		if sea[tmpM][tmpN] != 1 {
			sea[tmpM][tmpN] = sea[tmpM][tmpN] - 1
			borderKey := tmpM*tmpAAA + tmpN
			borders[borderKey] = 1
			if borderToIsland[borderKey] == nil {
				borderToIsland[borderKey] = map[int]int{}
			}
			borderToIsland[borderKey][islandKey] = 1
		}
	}
	if nIndex+1 < n {
		tmpM := mIndex
		tmpN := nIndex + 1
		if sea[tmpM][tmpN] != 1 {
			sea[tmpM][tmpN] = sea[tmpM][tmpN] - 1
			borderKey := tmpM*tmpAAA + tmpN
			borders[borderKey] = 1
			if borderToIsland[borderKey] == nil {
				borderToIsland[borderKey] = map[int]int{}
			}
			borderToIsland[borderKey][islandKey] = 1
		}
	}
	return borders
}

func generateArray(m, n int) [][]int {
	arr := make([][]int, m)
	for i := 0; i < m; i++ {
		arr[i] = make([]int, n)
	}
	return arr
}

func numIslands2_v1(m int, n int, positions [][]int) []int {
	sea := generateArray(m, n)
	islandCont := []int{}
	lastIslandCount := 0
	for _, position := range positions {
		mIndex := position[0]
		nIndex := position[1]

		oldSeaArea := sea[mIndex][nIndex]
		if oldSeaArea == 1 {
			// 本来就是岛屿，没必要操作
			continue
		}

		islandChange := 1 + oldSeaArea
		lastIslandCount = lastIslandCount + islandChange
		islandCont = append(islandCont, lastIslandCount)
		sea[mIndex][nIndex] = 1

		if mIndex > 0 {
			top := sea[mIndex-1][nIndex]
			if top != 1 { // 不是岛屿，必是海洋
				sea[mIndex-1][nIndex] = top - 1
			}
		}
		if mIndex+1 < m {
			bottom := sea[mIndex+1][nIndex]
			if bottom != 1 { // 不是岛屿，必是海洋
				sea[mIndex+1][nIndex] = bottom - 1
			}
		}

		if nIndex > 0 {
			left := sea[mIndex][nIndex-1]
			if left != 1 { // 不是岛屿，必是海洋
				sea[mIndex][nIndex-1] = left - 1
			}
		}
		if nIndex+1 < m {
			right := sea[mIndex][nIndex+1]
			if right != 1 { // 不是岛屿，必是海洋
				sea[mIndex][nIndex+1] = right - 1
			}
		}

	}

	return islandCont
}
