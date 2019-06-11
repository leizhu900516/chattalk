package utils

import (
	"fmt"
	"testing"
)

//加权轮询算法 weighted round robin
// 客服权重默认都为0
// 返回当前匹配到的客服信息
// go test -v
// go test -v wrr_test.go wrr.go
// go test -v -test.run TestWrr
/* 算法结果样例
目标 a [5, 2, 1] [-3, 2, 1]
目标 b [2, 4, 2] [2, -4, 2]
目标 a [7, -2, 3] [-1, -2, 3]
目标 a [4, 0, 4] [-4, 0, 4]
目标 c [1, 2, 5] [1, 2, -3]
目标 a [6, 4, -2] [-2, 4, -2]
目标 b [3, 6, -1] [3, -2, -1]
目标 a [8, 0, 0] [0, 0, 0]
*/

func TestWrr(t *testing.T)  {
	nodes := []*Node{
		&Node{"a",0,5},
		&Node{"b",0,2},
		&Node{"c",0,1},
	}
	for i:=0;i<8;i++{
		best :=	Wrr(nodes)
		fmt.Println(best.Uid,best)
	}
}


