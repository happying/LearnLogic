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
	resultMap := threeSum_V2(nums)
	end := time.Now()
	cost := end.Sub(start)
	fmt.Printf("函数执行耗时：%v \n", cost)
	fmt.Printf("结果：%v \n", len(resultMap))

}

/////////////////////

func threeSum_V2(nums []int) [][]int {
	var results [][]int
	var length = len(nums)
	if length < 3 {
		return results
	}
	sort.Ints(nums)
	// 为增加效率，针对某些极端情况做出额外处理
	minNum := nums[0] + nums[1] + nums[2]
	if minNum > 0 {
		// 最小的三个数相加都大于0了，没必要处理
		return results
	}
	if minNum == 0 {
		// 最小的三个数相加等于0了，那么也就意味着最小的第一个和第二个结合的最小的数，已经匹配到对的数了，剩下的那个数字必然固定，和它一样的不符合条件，和他不一样的一定比那个数大
		return [][]int{[]int{nums[0], nums[1], nums[2]}}
	}
	maxNum := nums[length-1] + nums[length-2] + nums[length-3]
	if maxNum < 0 {
		// 最大的3个数相加都小于0了，没必要处理了
		return results
	}
	if maxNum == 0 {
		// 最大的三个数相加等于0了，那么也就意味着最小的第一个和第二个结合的最小的数，已经匹配到对的数了，剩下的那个数字必然固定，和它一样的不符合条件，和他不一样的一定比那个数大
		return [][]int{[]int{nums[length-1], nums[length-2], nums[length-3]}}
	}

	for i := 0; i <= length-3; i++ {
		num_i := nums[i]
		if i > 0 && num_i == nums[i-1] {
			continue
		}
		k := length - 1
		j := i + 1
		for j < k {
			num_j := nums[j]
			num_k := nums[k]
			sum := num_j + num_k + num_i
			if sum == 0 {
				results = append(results, []int{num_i, num_j, num_k})
				for j = j + 1; j < k; j++ {
					if nums[j-1] != nums[j] {
						break
					}
				}
				for k = k - 1; k > j; k-- {
					if j >= k {
						break
					}
					if nums[k+1] != nums[k] {
						break
					}
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
