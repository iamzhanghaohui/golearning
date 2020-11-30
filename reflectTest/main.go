package main

import (
	"fmt"
	"reflect"
)

//定义一个Monster结构体

type Monster struct {
	Name  string  `json:"name"`
	Age   int     `json:"monster_age"`
	Score float32 `json:"成绩"`
	Sex   string
}

//方法，返回两个数的和
func (s Monster) GetSum(n1, n2 int) int {
	return n1 + n2
}

//方法，接收四个值，给s赋值
func (s Monster) Set(name string, age int, score float32, sex string) {
	s.Name = name
	s.Age = age
	s.Score = score
	s.Sex = sex
}

//方法，显示s的值
func (s Monster) Print() {
	fmt.Println("-----start~-----")
	fmt.Println(s)
	fmt.Println("---end~---")
}

func TestStruct(a interface{}) {
	//获取reflect.Type类型
	typ := reflect.TypeOf(a)
	//获取reflect.Value类型
	val := reflect.ValueOf(a)
	//获取到a对应的类别
	kd := val.Kind()
	//如果传入的不是strut,就退出
	if kd != reflect.Struct {
		fmt.Println("expect struct")
		return
	}
	//获取该结构体有多少个字段
	num := val.NumField()
	//应该是4个
	fmt.Printf("struct has %d fields\n", num)
	for i := 0; i < num; i++ {
		fmt.Printf("Field %d:值为=%v\n", i, val.Field(i))
		//获取struct表情，注意需要通过reflect.Type 来获取tag标签的值
		tagVal := typ.Field(i).Tag.Get("json")
		//如果该字段于tag标签就显示，否则就不显示
		if tagVal != "" {
			fmt.Printf("Field %d:tag 为=%d\n", i, tagVal)
		}
	}
	//获取到该结构体有多少个方法
	numOfMethod := val.NumMethod()

	fmt.Printf("struct has %d methods\n", numOfMethod)
	//var params []reflect.Value
	val.Method(1).Call(nil) //获取到第二个方法 然后调用它

	//调用结构体的第1个方法 Method(0)

	var params []reflect.Value //声明了 []reflect.Value
	params = append(params, reflect.ValueOf(10))
	params = append(params, reflect.ValueOf(40))
	//传入的参数是 []reflect.Value 返回[]reflect.Value
	res := val.Method(0).Call(params)
	fmt.Println("res=", res[0].Int())
}
func main() {
	var a Monster = Monster{
		Name:  "Free",
		Age:   100,
		Score: 11.1,
	}
	TestStruct(a)

}
