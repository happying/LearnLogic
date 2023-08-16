package logic

import (
	"fmt"
	"sort"
)

/*
给你一个长度为 n 的整数数组 nums 和 一个目标值 target。请你从 nums 中选出三个整数，使它们的和与 target 最接近。
返回这三个数的和。
假定每组输入只存在恰好一个解。
*/

func TestThreeSumClosest() {
	nums := []int{-1, 2, 1, -4}
	result := threeSumClosest(nums, 1)
	fmt.Println("result is: ", result)
}

func threeSumClosest(nums []int, target int) int {
	result := 0
	resultDiff := 0
	resultAbs := 0
	length := len(nums)
	sort.Ints(nums)
	for i := 0; i < length-2; i++ {
		num_i := nums[i]
		j := i + 1
		k := length - 1
		num_j := nums[j]
		num_k := nums[k]
		if i == 0 {
			result = num_i - nums[j] - nums[k]
			resultDiff = target - result
			if resultDiff == 0 {
				return result
			}
			resultAbs = abs(resultDiff)
		} else if nums[i-1] == num_i {
			continue
		}

		for j < k {
			num_j = nums[j]
			num_k = nums[k]
			tmpResult := num_i + num_j + num_k
			tmpResultDiff := target - tmpResult
			if tmpResultDiff == 0 {
				return tmpResult
			}
			tmpAbs := abs(tmpResultDiff)
			if tmpAbs < resultAbs {
				resultDiff = tmpResultDiff
				resultAbs = tmpAbs
			}
			if tmpResultDiff > 0 {
				j++
				for j < k && nums[j] == nums[j-1] {
					j++
				}
			} else {
				k--
				for j < k && nums[k] == nums[k+1] {
					k--
				}
			}

		}
	}

	return result
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
