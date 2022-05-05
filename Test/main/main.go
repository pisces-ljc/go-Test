package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type Student struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func reflectTest01(b interface{}) {
	rVal := reflect.ValueOf(b)
	a := rVal.String()
	//fmt.Println("a = ", a)
	var stu Student
	err := json.Unmarshal([]byte(a), &stu)
	if err != nil {
		fmt.Println("test err = ", err)
		return
	}
	fmt.Println("name = ", stu.Name, "age = ", stu.Age)
}

func (s Student) Set(name string, age int) {
	s.Name = name
	s.Age = age
}

func (s Student) Get() string {
	str := fmt.Sprintf("name = %v, age =%v", s.Name, s.Age)
	return str
}

func reflectTest02(b interface{}) {
	rVal := reflect.ValueOf(b)
	rType := reflect.TypeOf(b)
	stu := b.(Student)
	fmt.Println(stu.Name)
	//kd:reflect.struct
	kd := rVal.Kind()
	if kd != reflect.Struct {
		return
	}
	fmt.Println("kd =", kd)
	//获取结构体中的字段
	num := rVal.NumField()
	fmt.Println("num =", num)

	for i := 0; i < num; i++ {
		//获取标签的值
		tag := rType.Field(i).Tag.Get("json")
		fmt.Println("tag", i, "=", tag)
	}

	//获取方法的数量
	mNum := rVal.NumMethod()
	//call调用方法，方法顺序是按照ASCII进行排序的
	str := rVal.Method(0).Call(nil)
	fmt.Println("方法个数 =", mNum)
	fmt.Println(str)
	fmt.Printf("%T", str)
}

func main() {
	stu := Student{
		Name: "tom",
		Age:  10,
	}

	reflectTest02(stu)
	//fmt.Println(stu.Get())
	//data, err := json.Marshal(stu)
	//if err != nil {
	//	fmt.Println("err = ", err)
	//}
	//fmt.Println(string(data))
	//
	//str := string(data)
	//fmt.Println("str =", str)
	//reflectTest01(str)
}
