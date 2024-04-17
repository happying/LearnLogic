//
//  ThreeSumClosest.m
//  LearnLogic_oc
//
//  Created by fisher on 2024/4/10.
//

#import "ThreeSumClosest.h"
#import "BaseTool.h"
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

@implementation ThreeSumClosest

- (NSInteger)threeSumClosestWithNums:(NSArray *)nums target:(NSInteger)target {
    NSInteger result = NSIntegerMax;
    NSInteger count = nums.count;
    NSArray *sortedArr = [BaseTool sortNumArray:nums];
    
    for (NSInteger i = 0; i < count - 2; i++) {
        NSInteger num_i = [sortedArr[i] integerValue];
        if (i > 0 && num_i == [sortedArr[i-1] integerValue]) {
            continue;
        }
        
        NSInteger small = i + 1;
        NSInteger big = count - 1;
        
        while (small < big) {
            NSInteger num_small =[sortedArr[small] integerValue];
            NSInteger num_big =[sortedArr[big] integerValue];
            
            NSInteger sum =  num_small + num_big;
            NSInteger diff = sum - target;
            if (diff == 0) {
                return sum;
            } else if (labs(diff) < labs(result - target)) {
                result = sum;
            }
            
            if (diff > 0) {
                big = big - 1;
                while (small < big) {
                    NSInteger tmp_1 =[sortedArr[big] integerValue];
                    NSInteger tmp_2 =[sortedArr[big + 1] integerValue];
                    if (tmp_1 != tmp_2) {
                        break;
                    }
                    big = big - 1;
                }
            } else {
                small = small + 1;
                while (small < big) {
                    NSInteger tmp_1 =[sortedArr[small] integerValue];
                    NSInteger tmp_2 =[sortedArr[small - 1] integerValue];
                    if (tmp_1 != tmp_2) {
                        break;
                    }
                    small = small + 1;
                }
            }
        }
    }
    
    return result;
}

@end
