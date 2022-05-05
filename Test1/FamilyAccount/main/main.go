package main

import (
	"Test1/FamilyAccount/utils"
	"fmt"
)

//面向对象来处理
//把记账软件功能，封装到一个结构体中，调用该结构体的方法来记账
func main() {
	fmt.Println("面向对象实现")
	familyAccount := utils.NewFamilyAccount()
	familyAccount.MainMenu()
}
