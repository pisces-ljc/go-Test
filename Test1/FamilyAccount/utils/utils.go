package utils

import "fmt"

type FamilyAccount struct {
	key     string
	loop    bool
	balance float64
	money   float64
	note    string
	flag    bool
	details string
}

// NewFamilyAccount 编写一个工厂模式的构造方法，返回一个FamilyAccount实例
func NewFamilyAccount() *FamilyAccount {
	return &FamilyAccount{
		key:     "",
		loop:    true,
		balance: 10000.0,
		money:   0.0,
		note:    "",
		flag:    false,
		details: "收支\t账户金额\t收支金额\t说明",
	}
}

// 显示明细
func (fa *FamilyAccount) showDetails() {
	fmt.Println("--------------当前收支明细-----------------")
	if fa.flag {
		fmt.Println(fa.details)
	} else {
		fmt.Println("当前没有收支情况")
	}
}

// 登记收入
func (fa *FamilyAccount) income() {
	fmt.Println("本次收入金额：")
	fmt.Scanln(&fa.money)
	fa.balance += fa.money
	fmt.Println("本次收入说明：")
	fmt.Scanln(&fa.note)
	//记录收入情况
	fa.details += fmt.Sprintf("\n收入\t%v\t\t%v\t\t%v", fa.balance, fa.money, fa.note)
	fa.flag = true
}

// 登记支出
func (fa *FamilyAccount) pay() {
	fmt.Println("登记支出..")
	fmt.Scanln(&fa.money)
	if fa.money > fa.balance {
		fmt.Println("金额不足")
		return
	}
	fa.balance -= fa.money
	fmt.Println("本次支出说明：")
	fmt.Scanln(&fa.note)
	fa.details += fmt.Sprintf("\n支出\t%v\t\t%v\t\t%v", fa.balance, fa.money, fa.note)
	fa.flag = true
}

// 退出
func (fa *FamilyAccount) exit() {
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
		fa.loop = false
	}
}

// MainMenu 显示主菜单
func (fa *FamilyAccount) MainMenu() {
	for {
		fmt.Println("\n--------------家庭收支记账软件-----------------")
		fmt.Println("			   1.收支明细")
		fmt.Println("			   2.登记收入")
		fmt.Println("			   3.登记支出")
		fmt.Println("			   4.退出软件")
		fmt.Println()
		fmt.Println("请选择（1-4）：")
		fmt.Scanln(&fa.key)
		switch fa.key {
		case "1":
			fa.showDetails()
		case "2":
			fa.income()
		case "3":
			fa.pay()
		case "4":
			fa.exit()
		default:
			fmt.Println("请输入正确选项..")
		}

		if !fa.loop {
			break
		}
	}
}
