package main

import (
	"Test2/model"
	"Test2/service"
	"fmt"
)

type customerView struct {
	//接收用户输入
	key string
	//表示是否循环的显示主菜单
	loop bool

	customerService *service.CustomerService
}

// List 显示所有的客户信息
func (cv *customerView) list() {
	//首先获取到当前所有的客户信息（在切片中）
	customers := cv.customerService.List()
	//显示
	fmt.Println("---------------------------客户列表----------------------")
	fmt.Println("编号\t姓名\t性别\t年龄\t电话\t\t邮箱")
	for i := 0; i < len(customers); i++ {
		fmt.Println(customers[i].GetInfo())
	}
	fmt.Println("-------------------------客户列表完成---------------------")
}

func (cv *customerView) add() {
	fmt.Println("-----------------添加客户------------------")
	fmt.Printf("姓名:")
	name := ""
	fmt.Scanln(&name)
	fmt.Printf("性别:")
	gender := ""
	fmt.Scanln(&gender)
	fmt.Printf("年龄:")
	age := 0
	fmt.Scanln(&age)
	fmt.Printf("电话:")
	phone := ""
	fmt.Scanln(&phone)
	fmt.Printf("邮箱:")
	email := ""
	fmt.Scanln(&email)

	//id号不是由用户输入，而是由系统自动计算，id是唯一的
	customer := model.NewCustomer2(name, gender, age, phone, email)
	if cv.customerService.Add(customer) {
		fmt.Println("添加完成")
	}
}

func (cv *customerView) delete() {
	fmt.Println("---------------------删除用户-----------------")
	fmt.Printf("请输入待删除的编号（-1退出）：")
	id := -1
	fmt.Scanln(&id)
	if id == -1 {
		return
	}
	fmt.Printf("请确认待删除的编号（y/n）:")
	for {
		choice := ""
		fmt.Scanln(&choice)
		if choice == "n" || choice == "N" {
			break
		} else if choice == "y" || choice == "Y" {
			if cv.customerService.Delete(id) {
				fmt.Println("删除成功")
			} else {
				fmt.Println("删除失败")
			}
			break
		} else {
			fmt.Printf("请输入正确字符（y/n）")
		}
	}
}

// 退出
func (cv *customerView) exit() {
	fmt.Printf("是否退出（y/n）：")
	for {
		fmt.Scanln(&cv.key)
		if cv.key == "y" || cv.key == "Y" {
			cv.loop = false
			break
		} else if cv.key == "n" || cv.key == "N" {
			cv.loop = true
			break
		} else {
			fmt.Printf("输入有误，重新输入（y/n）：")
		}
	}

}

// 更新信息
func (cv *customerView) update() {
	fmt.Println("-------------------------------修改客户----------------------")
	fmt.Printf("请选择待修改的客户编号（-1退出）")
	id := -1
	fmt.Scanln(&id)
	if id == -1 {
		return
	}

	//判断该用户存在否
	index := cv.customerService.FindById(id)
	if index == -1 {
		fmt.Println("该用户不存在")
		return
	}

	//如果存在
	customer := cv.customerService.List()
	fmt.Printf("姓名（%v）：", customer[index].Name)
	name := ""
	fmt.Scanln(&name)
	fmt.Printf("性别（%v）：", customer[index].Gender)
	gender := ""
	fmt.Scanln(&gender)
	fmt.Printf("年龄（%v）：", customer[index].Age)
	age := 0
	fmt.Scanln(&age)
	fmt.Printf("电话（%v）：", customer[index].Phone)
	phone := ""
	fmt.Scanln(&phone)
	fmt.Printf("邮箱（%v）：", customer[index].Email)
	email := ""
	fmt.Scanln(&email)

	temp := model.NewCustomer2(name, gender, age, phone, email)
	if cv.customerService.Update(index, temp) {
		fmt.Println("已更新")
	} else {
		fmt.Println("更新失败")
	}

	fmt.Println("-------------------------------修改完成----------------------")
}

// 显示主菜单
func (cv *customerView) mainMenu() {
	for {
		fmt.Println("----------------客户端管理软件--------------")
		fmt.Println("				 1.添加客户")
		fmt.Println("				 2.修改客户")
		fmt.Println("				 3.删除客户")
		fmt.Println("				 4.客户列表")
		fmt.Println("				 5.退	出")
		fmt.Printf("请选择（1-5）:")

		fmt.Scanln(&cv.key)
		switch cv.key {
		case "1":
			cv.add()
		case "2":
			cv.update()
		case "3":
			cv.delete()
		case "4":
			cv.list()
		case "5":
			cv.exit()
		default:
			fmt.Println("请重新输入！")
		}

		if !cv.loop {
			break
		}
	}
	fmt.Println("已成功退出！")
}

func main() {
	//在main函数中，创建一个customerView,并运行主菜单
	customerView := customerView{
		key:  "",
		loop: true,
	}
	//对结构体中customerService字段的初始化
	customerView.customerService = service.NewCustomerService()

	customerView.mainMenu()
}
