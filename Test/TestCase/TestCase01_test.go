package main

import (
	"testing"
)

func TestAddUpper(t *testing.T) {

	res := addUpper(1, 2)
	if res != 3 {
		//fmt.Println("err!")
		t.Fatalf("期望值为 = %v, 实际值为 = %v\n", 3, res)
	}

	//如果正确，则输出日志
	t.Logf("执行正确")
}
