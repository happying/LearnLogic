//
//  ThreeSumClosest.swift
//  LearnLogic_swift
//
//  Created by fisher on 2024/4/10.
//

import Foundation

/*
 给你一个长度为 n 的整数数组 nums 和 一个目标值 target。请你从 nums 中选出三个整数，使它们的和与 target 最接近。

 返回这三个数的和。

 假定每组输入只存在恰好一个解。
 示例 1：

 输入：nums = [-1,2,1,-4], target = 1
 输出：2
 解释：与 target 最接近的和是 2 (-1 + 2 + 1 = 2) 。
 示例 2：

 输入：nums = [0,0,0], target = 1
 输出：0
  

 提示：

 3 <= nums.length <= 1000
 -1000 <= nums[i] <= 1000
 -10^4 <= target <= 10^4
 */

class ThreeSumClosest {

    func test() {
        var nums = [833,736,953,-584,-448,207,128,-445,126,248,871,860,333,-899,463,488,-50,-331,903,575,265,162,-733,648,678,549,579,-172,-897,562,-503,-508,858,259,-347,-162,-505,-694,300,-40,-147,383,-221,-28,-699,36,-229,960,317,-585,879,406,2,409,-393,-934,67,71,-312,787,161,514,865,60,555,843,-725,-966,-352,862,821,803,-835,-635,476,-704,-78,393,212,767,-833,543,923,-993,274,-839,389,447,741,999,-87,599,-349,-515,-553,-14,-421,-294,-204,-713,497,168,337,-345,-948,145,625,901,34,-306,-546,-536,332,-467,-729,229,-170,-915,407,450,159,-385,163,-420,58,869,308,-494,367,-33,205,-823,-869,478,-238,-375,352,113,-741,-970,-990,802,-173,-977,464,-801,-408,-77,694,-58,-796,-599,-918,643,-651,-555,864,-274,534,211,-910,815,-102,24,-461,-146];
        var target = -7111
        
//        nums = [4,0,5,-5,3,3,0,-4,-5]
//        target = -2
        
        nums = [321,413,82,812,-646,-858,729,609,-339,483,-323,-399,-82,-455,18,661,890,-328,-311,520,-865,-174,55,685,-636,462,-172,-696,-296,-832,766,-808,-763,853,482,411,703,655,-793,-121,-726,105,-966,-471,612,551,-257,836,-94,-213,511,317,-293,279,-571,242,-519,386,-670,-806,-612,-433,-481,794,712,378,-325,-564,477,169,601,971,-300,-431,-152,285,-899,978,-419,708,536,-816,-335,284,384,-922,-941,633,934,497,-351,62,392,-493,-44,-400,646,-912,-864,835,713,-12,322,-228,340,-42,-307,-580,-802,-914,-142,575,-684,-415,718,-579,759,579,732,-645,525,114,-880,-603,-699,-101,-738,-887,327,192,747,-614,393,97,-569,160,782,-69,235,-598,-116,928,-805,-76,-521,671,417,600,-442,236,831,637,-562,613,-705,-158,-237,-299,808,-734,364,919,251,-163,-343,899]
        target = 2218
        
        
        let result = threeSumClosest(nums, target)
        
        print(result)
        
    }
    
    func threeSumClosest(_ nums: [Int], _ target: Int) -> Int {
        let sortedArr = nums.sorted()
        let length = sortedArr.count
        var result = 0
        var isFirst = true
        
        for i in 0..<length-2 {
            if i > 0 && sortedArr[i] == sortedArr[i-1] {
                continue
            }
            let num_i = sortedArr[i]

            var j = i + 1
            var k = length - 1
            while j < k {
                let num_j = sortedArr[j]
                let num_k = sortedArr[k]
                let sum = num_i + num_k + num_j
                let diffNow = sum - target
                let diffResult = result - target
                if isFirst {
                    isFirst = false
                    result = sum
                    if diffNow == 0 {
                        // 差值不可能比0更小了，然后每组只有一个解，直接返回
                        return result
                    }
                } else if abs(diffNow) < abs(diffResult) || abs(diffNow) == abs(diffResult) {
                    result = sum
                    if diffNow == 0 {
                        // 差值不可能比0更小了，然后每组只有一个解，直接返回
                        return result
                    }
                }
                
                // 把相同的j和k都直接过掉
                if diffNow < 0 {
                    // 尽量向0逼近，所以这里差值小于0，应该要放大差值，所以最小值指针前移
                    j = j + 1
                    while j < k {
                        let tmp_1 = sortedArr[j - 1]
                        let tmp_2 = sortedArr[j]
                        if tmp_1 != tmp_2 {
                            break
                        }
                        j = j + 1
                    }
                } else {
                    k = k - 1
                    while j < k {
                        let tmp_1 = sortedArr[k + 1]
                        let tmp_2 = sortedArr[k]
                        if tmp_1 != tmp_2 {
                            break
                        }
                        k = k - 1
                    }
                }
            }
        }
        
        return result
    }
}
