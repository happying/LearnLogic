//
//  NumIsland.swift
//  LearnLogic_swift
//
//  Created by fisher on 2024/4/11.
//

import Foundation

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

class NumIslands {
    
    func test() {
        var islands =  [[Character]]()
        
        islands = [
            ["1", "1", "1", "1", "0"],
            ["1", "1", "0", "1", "0"],
            ["1", "1", "0", "0", "0"],
            ["0", "0", "0", "0", "0"]
        ]
        
        let result = numIslands_v2(islands)
        print(result)
    }

    enum Direction: Int {
        case left = 0
        case right = 1
        case top = 3
        case bottom = 4
    }

    func numIslands(_ grid: [[Character]]) -> Int {
        
        var gridTmp = Array(grid)
        var result = 0
        
        for index1 in 0..<gridTmp.count {
            for index2 in 0..<gridTmp[index1].count {
                if gridTmp[index1][index2] == "0" || gridTmp[index1][index2] == "a"{
                    continue
                }
                result = result + 1
                check(grid: &gridTmp, vertical: index1, horizontal: index2, direction: Direction.left)
                check(grid: &gridTmp, vertical: index1, horizontal: index2, direction: Direction.right)
                check(grid: &gridTmp, vertical: index1, horizontal: index2, direction: Direction.top)
                check(grid: &gridTmp, vertical: index1, horizontal: index2, direction: Direction.bottom)
            }
        }
        
        return result
    }

    func check(grid: inout [[Character]], vertical: Int, horizontal: Int, direction: Direction) {
        var newVer = vertical
        var newHor = horizontal
        if direction == Direction.right {
            newHor = newHor + 1
        } else if direction == Direction.left {
            newHor = newHor - 1;
        } else if direction == Direction.top {
            newVer = newVer - 1;
        } else {
            newVer = newVer + 1
        }
        if newVer < 0 || newHor < 0 ||  newVer >= grid.count || newHor >= grid[vertical].count {
            // 越界结束
            return
        }
        if grid[newVer][newHor] == "0" || grid[newVer][newHor] == "a" {
            // 触碰到边界，可以停止了
            return
        }
        // 这里转一个标记，避免重复检查
        grid[newVer][newHor] = "a"
        if direction == Direction.right {
            // 从左向右，所以左边不检查，检查剩下的三边
            check(grid: &grid, vertical: newVer, horizontal: newHor, direction: Direction.right)
            check(grid: &grid, vertical: newVer, horizontal: newHor, direction: Direction.top)
            check(grid: &grid, vertical: newVer, horizontal: newHor, direction: Direction.bottom)
        } else if direction == Direction.left {
            check(grid: &grid, vertical: newVer, horizontal: newHor, direction: Direction.left)
            check(grid: &grid, vertical: newVer, horizontal: newHor, direction: Direction.top)
            check(grid: &grid, vertical: newVer, horizontal: newHor, direction: Direction.bottom)
        } else if direction == Direction.top {
            check(grid: &grid, vertical: newVer, horizontal: newHor, direction: Direction.left)
            check(grid: &grid, vertical: newVer, horizontal: newHor, direction: Direction.top)
            check(grid: &grid, vertical: newVer, horizontal: newHor, direction: Direction.right)
        } else {
            check(grid: &grid, vertical: newVer, horizontal: newHor, direction: Direction.left)
            check(grid: &grid, vertical: newVer, horizontal: newHor, direction: Direction.bottom)
            check(grid: &grid, vertical: newVer, horizontal: newHor, direction: Direction.right)
        }
        
    }
    
    
    func numIslands_v2(_ grid: [[Character]]) -> Int {
        var grid = grid
        var result = 0
        for index1 in 0..<grid.count {
            for index2 in 0..<grid[index1].count {
                if grid[index1][index2] != "1" {
                    continue
                }
                result = result + 1
                dfs(&grid, vertical: index1, horizontal: index2)
            }
        }
        
        return result
    }
    
    func dfs (_ grid: inout [[Character]], vertical: Int, horizontal: Int) {
        if vertical < 0 || horizontal < 0 || vertical >= grid.count || horizontal >= grid[vertical].count || grid[vertical][horizontal] != "1" {
            return
        }
        grid[vertical][horizontal] = "2"
        dfs(&grid, vertical: vertical, horizontal: horizontal - 1)
        dfs(&grid, vertical: vertical, horizontal: horizontal + 1)
        dfs(&grid, vertical: vertical + 1, horizontal: horizontal)
        dfs(&grid, vertical: vertical - 1, horizontal: horizontal)
    }

}

