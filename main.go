package main

import (
	"fmt"
	"os"
)

type Usb interface {
	Start()
	Stop()
}

type Phone struct {
	name string
}

//让phone实现usb接口的方法

func (p Phone) Start() {
	fmt.Println("手机开始工作")
}
func (p Phone) Stop() {
	fmt.Println("手机停止工作")
}

type Camera struct {
	name string
}

func (c Camera) Start() {
	fmt.Println("相机开始工作")
}
func (c Camera) Stop() {
	fmt.Println("相机停止工作")
}

func main() {
	var usbArr [3]Usb
	usbArr[0] = Phone{"vivo"}
	usbArr[1] = Camera{"小米"}

	fmt.Println(usbArr)
	if phone, ok := usbArr[1].(Phone); ok {
		phone.Start()
	}
	os.Open("")

}
