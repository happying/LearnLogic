//
//  ThreeSum.m
//  LearnLogic_oc
//
//  Created by fisher on 2024/4/10.
//

#import "ThreeSum.h"
#import "BaseTool.h"

@implementation ThreeSum


- (void)testThreeSum {
    NSArray *testArr = @[@(-1),@(0),@(1),@(2),@(-1),@(-4)];
//    testArr = @[@0, @1, @1];
//    testArr = @[@0, @0, @0];
    NSArray *result = [self threeSum:testArr];
    NSLog(@"result:%@", result);
}

- (NSArray *)threeSum:(NSArray *)nums {
    NSArray *sortedArr = [BaseTool sortNumArray:nums];
    NSInteger length = sortedArr.count;
    
    NSMutableArray *result = [@[] mutableCopy];
    
    for (NSInteger i = 0; i < length - 2; i++) {
        while (i > 0) {
            NSInteger num_i_current = [[sortedArr objectAtIndex:i] integerValue];
            NSInteger num_i_reduce_1 = [sortedArr[i-1] integerValue];
            if (num_i_current == num_i_reduce_1) {
                i++;
            } else {
                break;
            }
        }
        NSInteger num_i = [[sortedArr objectAtIndex:i] integerValue];

        NSInteger j = i + 1;
        NSInteger k = length - 1;
        
        while (j < k) {
            NSInteger num_j = [sortedArr[j] integerValue];
            NSInteger num_k = [sortedArr[k] integerValue];
            NSInteger sum = num_i + num_j + num_k;
            if (sum == 0) {
                [result addObject:@[@(num_i), @(num_j), @(num_k)]];
                j++;
                while (j < k) {
                    NSInteger num_j_tmp = [sortedArr[j] integerValue];
                    NSInteger num_j_reduce_1 = [sortedArr[j-1] integerValue];
                    if (num_j_tmp != num_j_reduce_1) {
                        break;
                    }
                    j++;
                }
                
                k--;
                while (j < k) {
                    NSInteger num_k_tmp = [sortedArr[k] integerValue];
                    NSInteger num_k_add_1 = [sortedArr[k+1] integerValue];
                    if (num_k_tmp != num_k_add_1) {
                        break;
                    }
                    k--;
                }
            } else if (sum < 0) {
                j++;
            } else {
                k--;
            }
        }
    }
    
    return result;
}



@end
