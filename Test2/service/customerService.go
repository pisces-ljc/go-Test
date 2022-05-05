package service

import (
	"Test2/model"
)

// CustomerService 该结构体完成对Customer的操作，包括增删改查
type CustomerService struct {
	customers []model.Customer
	//声明字段，表示当前切片含有多少个客户
	//该字段可以作为新客户的ID+1
	customerNum int
}

// NewCustomerService 编写一个方法，可以返回 *CustomerService
func NewCustomerService() *CustomerService {
	//为了能看到有客户在切片中，先初始化一个客户
	customerService := &CustomerService{}
	customerService.customerNum = 1
	customer := model.NewCustomer(
		1, "张三", "男", 18, "13111111111", "xxxxx@zzz.com")

	customerService.customers = append(customerService.customers, customer)
	return customerService
}

//	返回客户切片
func (cs *CustomerService) List() []model.Customer {
	return cs.customers
}

// Add 添加客户到customer切片中
func (cs *CustomerService) Add(customer model.Customer) bool {
	//确定分配ID的规则
	cs.customerNum += 1
	customer.Id = cs.customerNum
	cs.customers = append(cs.customers, customer)
	return true
}

// FindById 查找 根据ID查找切片中对应的位置
func (cs *CustomerService) FindById(id int) int {
	index := -1
	for i := 0; i < len(cs.customers); i++ {
		if cs.customers[i].Id == id {
			index = i
		}
	}

	return index
}

// Delete 删除
func (cs *CustomerService) Delete(id int) bool {
	index := cs.FindById(id)
	if index == -1 {
		return false
	}

	cs.customers = append(cs.customers[:index], cs.customers[index+1:]...)
	return true
}

func (cs *CustomerService) Update(id int, customer model.Customer) bool {
	if customer.Name != "" {
		cs.customers[id].Name = customer.Name
	}

	if customer.Age != 0 {
		cs.customers[id].Age = customer.Age
	}

	if customer.Gender != "" {
		cs.customers[id].Gender = customer.Gender
	}

	if customer.Phone != "" {
		cs.customers[id].Phone = customer.Phone
	}

	if customer.Email != "" {
		cs.customers[id].Email = customer.Email
	}
	
	return true
}
