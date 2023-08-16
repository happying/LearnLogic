package logic

import (
	"sort"
	"strconv"
)

/*
给你一个整数数组 nums ，判断是否存在三元组 [nums[i], nums[j], nums[k]] 满足 i != j、i != k 且 j != k ，同时还满足 nums[i] + nums[j] + nums[k] == 0 。请
你返回所有和为 0 且不重复的三元组。
注意：答案中不可以包含重复的三元组。
*/

func TestThreeSum() {
	nums := []int{0, 0, 0, 0}
	//nums := []int{-1, 0, 1, 2, -1, -4}
	resultMap := threeSum(nums)
	println(resultMap)
}

func threeSum(nums []int) [][]int {
	var results [][]int
	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] { // 跳过重复元素
			continue
		}
		j, k := i+1, len(nums)-1
		for j < k {
			sum := nums[i] + nums[j] + nums[k]
			if sum == 0 {
				results = append(results, []int{nums[i], nums[j], nums[k]})
				j++
				k--
				for j < k && nums[j] == nums[j-1] { // 跳过重复元素
					j++
				}
				for j < k && nums[k] == nums[k+1] { // 跳过重复元素
					k--
				}
			} else if sum < 0 {
				j++
			} else {
				k--
			}
		}
	}
	return results
}

////////////////////////
func threeSum_V1(nums []int) [][]int {
	resultMap := map[string][]int{}
	numMap := map[string][]int{}
	sort.Ints(nums)
	for index, num := range nums {
		target := 0 - num
		tmpMap, tmpNumMap := twoSum(nums, target, index, num, numMap)
		resultMap = mergeMaps(resultMap, tmpMap)
		numMap = mergeMaps(numMap, tmpNumMap)
	}

	resultArr := [][]int{}
	for _, ints := range resultMap {
		resultArr = append(resultArr, ints)
	}

	return resultArr
}

func twoSum(nums []int, target int, thirdId int, thirdNum int, thirdNumMap map[string][]int) (map[string][]int, map[string][]int) {
	numMap := map[int]int{}
	resultMap := map[string][]int{}

	for index := thirdId + 1; index < len(nums); index++ {
		num := nums[index]
		num2 := target - num
		for index2 := index + 1; index2 < len(nums); index2++ {
			if index == index2 || index == thirdId || index2 == thirdId {
				continue
			}

			// 将index从小到大排序，并生成string作为key，避免重复
			key := mapKey(index, index2, thirdId)
			if _, ok := resultMap[key]; !ok {
				numMapKey := mapKey(num, num2, thirdNum)
				if _, ok := thirdNumMap[numMapKey]; !ok {
					resultMap[key] = []int{num, num2, thirdNum}
				}
				thirdNumMap[numMapKey] = []int{}
			}
		}
		numMap[num] = index
	}

	return resultMap, thirdNumMap
}

func mergeMaps(m1, m2 map[string][]int) map[string][]int {
	// 创建一个新的map，用于存放合并后的结果
	merged := make(map[string][]int)
	// 先将第一个map中的键值对添加到merged中
	for k, v := range m1 {
		merged[k] = v
	}
	// 再将第二个map中的键值对添加到merged中
	for k, v := range m2 {
		merged[k] = v
	}
	return merged
}

func mapKey(a, b, c int) string {
	arr := []int{a, b, c}
	sort.Ints(arr)
	str := ""
	for _, num := range arr {
		if str == "" {
			str = strconv.Itoa(num)
			continue
		}
		str = str + "_" + strconv.Itoa(num)
	}

	return str
}

func min(a, b, c int) int {
	if a <= b && a <= c {
		return a
	} else if b <= a && b <= c {
		return b
	} else {
		return c
	}
}
