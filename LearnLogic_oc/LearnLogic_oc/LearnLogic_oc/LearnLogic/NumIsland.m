//
//  NumIsland.m
//  LearnLogic_oc
//
//  Created by fisher on 2024/4/11.
//

#import "NumIsland.h"

@implementation NumIsland

- (void)testNumIsland {
    NSArray*  islands = @[
        [@[@"1", @"1", @"1", @"1", @"0"] mutableCopy],
        [@[@"1", @"1", @"0", @"1", @"0"] mutableCopy],
        [@[@"1", @"1", @"0", @"0", @"0"] mutableCopy],
        [@[@"0", @"0", @"0", @"0", @"0"] mutableCopy]
    ];
    
    NSInteger num = [self NumIslandWith:islands];
    
    printf("%ld", (long)num);
}

- (NSInteger)NumIslandWith:(NSArray *)island {
    NSInteger result = 0;
    
    for (NSInteger vertical = 0; vertical < [island count]; vertical++) {
        NSArray* horArr = island[vertical];
        for (NSInteger horizontal = 0; horizontal < [horArr count]; horizontal++) {
            if (![horArr[horizontal] isEqual:@"1"]) {
                continue;
            }
            result++;
            [NumIsland DeepFirstSearch:island vertical:vertical horizontal:horizontal];
        }
    }
    
    
    return result;
}

+ (void)DeepFirstSearch:(NSArray *)island vertical:(NSInteger)vertical horizontal:(NSInteger)horizontal {
    if (vertical < 0 || horizontal < 0 || vertical >= island.count || horizontal >= [island[vertical] count] || ![island[vertical][horizontal] isEqual:@"1"]) {
        return;
    }
    [island[vertical] setObject:@"2" atIndex:horizontal];
    [self DeepFirstSearch:island vertical:vertical+1 horizontal:horizontal];
    [self DeepFirstSearch:island vertical:vertical-1 horizontal:horizontal];
    [self DeepFirstSearch:island vertical:vertical horizontal:horizontal+1];
    [self DeepFirstSearch:island vertical:vertical horizontal:horizontal-1];
    
}

@end
