//
//  BaseTool.m
//  LearnLogic_oc
//
//  Created by fisher on 2024/4/10.
//

#import "BaseTool.h"

@implementation BaseTool

+ (NSArray *)sortNumArray:(NSArray *)array {
    NSArray *sortedArr = [array sortedArrayUsingComparator:^NSComparisonResult(id  _Nonnull obj1, id  _Nonnull obj2) {
        NSInteger obj1Num = [obj1 integerValue];
        NSInteger obj2Num = [obj2 integerValue];
        if (obj1Num < obj2Num) {
            return NSOrderedAscending;
        } else if (obj1Num > obj2Num) {
            return NSOrderedDescending;
        }
        
        return NSOrderedSame;
    }];
    
    return sortedArr;
}


@end
