package logic

/*
给你一个整数数组 nums ，判断是否存在三元组 [nums[i], nums[j], nums[k]] 满足 i != j、i != k 且 j != k ，同时还满足 nums[i] + nums[j] + nums[k] == 0 。请
你返回所有和为 0 且不重复的三元组。
注意：答案中不可以包含重复的三元组。
*/

func ThreeSum(nums []int) [][]int {
	for index, num := range nums {
		target := 0 - num
		tmpMap := TwoSum(nums, target)

	}
	return nil
}

func twoSum(nums []int, target int, thirdId int) map[int][]int {
	numMap := map[int]int{}
	resultMap := map[int][]int{}
	for index, num := range nums {
		num2 := target - num
		if index2, ok := numMap[num2]; ok {
			if index == index2 || index == thirdId || index2 == thirdId {
				continue
			}

			// 一个简单的排序，拿小的index来做map key,方便去重
			mapKey := index
			if index2 < index {
				mapKey = index2
			}
			if _, ok := resultMap[mapKey]; !ok && index != index2 {
				resultMap[mapKey] = []int{index, index2}
			}
		}
		numMap[num] = index
	}

	return resultMap
}
