package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name string
	Age  int
}

func reflectTest01(b interface{}) {
	//通过反射获取的传入的变量的type，kind 值
	//1.先获取到reflect.Type
	rTyp := reflect.TypeOf(b)
	fmt.Println("rType=", rTyp)
	//2.获取到 reflect，Value
	rVal := reflect.ValueOf(b)

	n2 := 2 + rVal.Int()
	fmt.Println("n2=", n2)

	fmt.Printf("rVal = %v rVal type = %T\n", rVal, rVal)

	iV := rVal.Interface()
	num2 := iV.(int)
	fmt.Println("num2=", num2)
}

func reflectTest02(b interface{}) {
	//通过反射获取的传入的变量的type，kind，值
	//1.先获取到reflect.Type
	rTyp := reflect.TypeOf(b)
	fmt.Println("rType=", rTyp)
	fmt.Println("rType=", rTyp)

	//2.获取到reflect.Value
	rVal := reflect.ValueOf(b)
	//下面我们将 rVal 转成 interface{}
	iV := rVal.Interface()

	fmt.Printf("iv=%v iv type=%T \n", iV, iV)

	//将interface{} 通过断言转成需要的类型
	//这里，我们就简单实用了一带检测的类型断言
	stu, ok := iV.(Student)
	if ok {
		fmt.Printf("stu.Name=%v\n", stu.Name)
	}

}

func main() {

	stu := Student{
		Name: "tom",
		Age:  20,
	}
	reflectTest02(stu)

}
