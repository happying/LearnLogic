package logic

import (
	"fmt"
	"time"
)

func TimeCost(testFunc func()) {
	//nums := BigLengthIntArr
	//nums := BigLengthZeroIntArr
	//nums := []int{0, 0, 0, 0}
	start := time.Now()
	testFunc()
	end := time.Now()
	cost := end.Sub(start)
	fmt.Printf("函数执行耗时：%v \n", cost)

}
