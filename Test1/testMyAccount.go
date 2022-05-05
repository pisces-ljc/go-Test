package main

import (
	"fmt"
)

/*
	基本功能
	1.完成显示主菜单，并可以退出
	2.完成显示明细和登记收入的功能
	todo：需要显示明细，因此需要定义变量来记录detail string 、余额、收入
	3.完成登记支出的功能
*/
func main() {

	//声明变量用于接收用户的输入
	key := ""
	//声明变量退出for循环
	loop := true

	//定义账户的余额
	balance := 10000.0
	//每次收支的金额
	money := 0.0
	//每次收支的说明
	note := ""
	//定义一个变量，记录是否存在收支行为
	flag := false
	//收支详情：当有收支时，对其拼接处理
	details := "收支\t账户金额\t收支金额\t说明"

	//显示主菜单
	for {
		fmt.Println("\n--------------家庭收支记账软件-----------------")
		fmt.Println("			   1.收支明细")
		fmt.Println("			   2.登记收入")
		fmt.Println("			   3.登记支出")
		fmt.Println("			   4.退出软件")
		fmt.Println()
		fmt.Println("请选择（1-4）：")
		fmt.Scanln(&key)
		switch key {
		case "1":
			fmt.Println("--------------当前收支明细-----------------")
			if flag {
				fmt.Println(details)
			} else {
				fmt.Println("当前没有收支情况")
			}
		case "2":
			fmt.Println("本次收入金额：")
			fmt.Scanln(&money)
			balance += money
			fmt.Println("本次收入说明：")
			fmt.Scanln(&note)
			//记录收入情况
			details += fmt.Sprintf("\n收入\t%v\t\t%v\t\t%v", balance, money, note)
			flag = true
		case "3":
			fmt.Println("登记支出..")
			fmt.Scanln(&money)
			if money > balance {
				fmt.Println("金额不足")
				break
			}
			balance -= money
			fmt.Println("本次支出说明：")
			fmt.Scanln(&note)
			details += fmt.Sprintf("\n支出\t%v\t\t%v\t\t%v", balance, money, note)
			flag = true
		case "4":
			fmt.Println("确认要退出吗？(y/n)")
			choice := ""
			for {
				fmt.Scanln(&choice)
				if choice == "y" || choice == "n" {
					break
				}
				fmt.Println("输入错误，请重新输入（y/n）")
			}
			if choice == "y" {
				loop = false
			}
		default:
			fmt.Println("请输入正确选项..")
		}

		if !loop {
			break
		}
	}

	fmt.Println("成功退出软件！")
}
