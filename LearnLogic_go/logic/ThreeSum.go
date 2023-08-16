package logic

import (
	"fmt"
	"sort"
	"strconv"
	"time"
)

/*
给你一个整数数组 nums ，判断是否存在三元组 [nums[i], nums[j], nums[k]] 满足 i != j、i != k 且 j != k ，同时还满足 nums[i] + nums[j] + nums[k] == 0 。请
你返回所有和为 0 且不重复的三元组。
注意：答案中不可以包含重复的三元组。
*/

func TestThreeSum() {
	//nums := BigLengthIntArr
	//nums := BigLengthZeroIntArr
	//nums := []int{0, 0, 0, 0}
	nums := []int{-1, 0, 1, 2, -1, -4}
	start := time.Now()
	resultMap := threeSum_V1(nums)
	end := time.Now()
	cost := end.Sub(start)
	fmt.Printf("函数执行耗时：%v \n", cost)
	fmt.Printf("结果：%v \n", len(resultMap))

}

func threeSum(nums []int) [][]int {
	var results [][]int
	jump1 := 0
	jump2 := 0
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
					jump1++
					j++
				}
				for j < k && nums[k] == nums[k+1] { // 跳过重复元素
					jump2++
					k--
				}
			} else if sum < 0 {
				j++
			} else {
				k--
			}
		}
	}
	fmt.Printf("a: %v b: %v \n", jump1, jump2)

	return results
}

////////////////////////
func threeSum_V1(nums []int) [][]int {
	resultMap := map[string][]int{}
	existsNumMap := map[string][]int{}
	numMap := map[int]int{}
	length := len(nums)
	resultArr := [][]int{}

	sort.Ints(nums)
	for index := 0; index < length; index++ {
		num := nums[index]
		if index != 0 && nums[index-1] == num {
			continue
		}
		target := 0 - num
		//twoSum(nums, target, index, num, &existsNumMap, &resultMap, &numMap)
		for index2 := index + 1; index2 < length; index2++ {
			num2 := nums[index2]
			//if nums[index2-1] == num2 {
			//	continue
			//}
			num3 := target - num2
			if index3, ok := numMap[num3]; ok {
				if index2 == index3 || index3 == index {
					continue
				}
				key := mapKey(index2, index3, index)
				if _, ok := resultMap[key]; !ok {
					numMapKey := mapKey(num2, num3, num)
					if _, ok := existsNumMap[numMapKey]; !ok {
						singleArr := []int{num2, num3, num}
						resultMap[key] = singleArr
						resultArr = append(resultArr, singleArr)
					}
					existsNumMap[numMapKey] = []int{}
				}
			}
			numMap[num2] = index2
		}
	}

	return resultArr
}

func twoSum(nums []int, target int, thirdId int, thirdNum int, existsNumMap *map[string][]int, resultMap *map[string][]int, numMap *map[int]int) {
	//numMap := map[int]int{}
	length := len(nums)
	for index2 := thirdId + 1; index2 < length; {
		num2 := nums[index2]
		num3 := target - num2
		if index3, ok := (*numMap)[num3]; ok {
			if index2 == index3 || index3 == thirdId {
				continue
			}

			key := mapKey(index2, index3, thirdId)
			if _, ok := (*resultMap)[key]; !ok {
				numMapKey := mapKey(num2, num3, thirdNum)
				if _, ok := (*existsNumMap)[numMapKey]; !ok {
					(*resultMap)[key] = []int{num2, num3, thirdNum}
				}
				(*existsNumMap)[numMapKey] = []int{}
			}
		}
		(*numMap)[num2] = index2
	}
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

func min(a, b, c int) int {
	if a <= b && a <= c {
		return a
	} else if b <= a && b <= c {
		return b
	} else {
		return c
	}
}
