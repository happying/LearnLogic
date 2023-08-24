package logic

import "fmt"

/*
在一条环路上有 n 个加油站，其中第 i 个加油站有汽油 gas[i] 升。
你有一辆油箱容量无限的的汽车，从第 i 个加油站开往第 i+1 个加油站需要消耗汽油 cost[i] 升。你从其中的一个加油站出发，开始时油箱为空。
给定两个整数数组 gas 和 cost ，如果你可以按顺序绕环路行驶一周，则返回出发时加油站的编号，否则返回 -1 。如果存在解，则 保证 它是 唯一 的。
*/

func TestGas() {
	gas := GasLong
	cost := CostLong
	gas = []int{1, 2, 3, 4, 5}
	cost = []int{3, 4, 5, 1, 2}
	gas = []int{0, 0, 0, 0, 0}
	cost = []int{0, 0, 0, 0, 2}
	result := canCompleteCircuit(gas, cost)
	fmt.Println(result)
}

//[1,2,3,4,5]
//[3,4,5,1,2]

func canCompleteCircuit(gas []int, cost []int) int {
	total := 0
	locateIndex := 0
	length := len(gas)
	loopCount := 0
	for loopCount < length && locateIndex < length {
		index := (locateIndex + loopCount) % length
		diff := gas[index] - cost[index]
		total = total + diff
		if total >= 0 {
			loopCount++
			if loopCount == length {
				return locateIndex
			}
			continue
		} else {
			locateIndex = locateIndex + loopCount + 1
			total = 0 // 重置
			loopCount = 0
		}
	}

	return -1
}

func canCompleteCircuit_v1(gas []int, cost []int) int {
	subtracts := []int{}
	startIndex := []int{}
	length := len(gas)
	for i := 0; i < length; i++ {
		gasNum := gas[i]
		cosNum := cost[i]
		diff := gasNum - cosNum
		subtracts = append(subtracts, diff)
		if diff >= 0 {
			startIndex = append(startIndex, i)
		}
	}

	for _, startPoint := range startIndex {
		total := 0
		for countNum := 0; countNum < length; {
			current := startPoint + countNum
			if current >= length {

				current = current - length
			}
			total = total + subtracts[current]
			if total < 0 {
				break
			}
			countNum++
			if countNum >= len(subtracts) { // 跑完一圈了
				return startPoint
			}
		}
	}

	return -1
}

/*
	来源于最佳答案，仔细思考一下，这个解法在原解法的基础上，更近一步优化了
	优化来源于总量控制，一个二分的思想，如果总油数小于总消耗数，则必然无解，反之则必然有解
	则 if totalSum < 0 判断总数是否充足，在总数充足的情况下，只需要找到某一个起点，其一直到终点为止的汽油都足够(递归反证，如果前半段还存在幺儿子，那么此题必然无解，无解的话与第一个判断的结论相悖，所以必然有解，则找出的第一个到终点仍然够油耗的节点则为解)
*/
func canCompleteCircuit_from_result(gas []int, cost []int) int {
	totalSum := 0
	start := 0
	curSum := 0
	for i := 0; i < len(gas); i++ {
		curSum += gas[i] - cost[i]   // 当前油-到下一个的油
		totalSum += gas[i] - cost[i] // 这个是总体的油
		if curSum < 0 {
			start = i + 1 // 因为当前油不足以支持以该加油站为起点，必须靠totalSum支撑
			curSum = 0    // 还原
		}
	}
	if totalSum < 0 {
		return -1
	}
	return start
}
