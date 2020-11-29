package main

import "fmt"

func writeData(intChan chan int) {
	for i := 1; i < 50; i++ {
		intChan <- i
		fmt.Println("writeData", i)
	}
	close(intChan)
}

func readData(intChan chan int, exitChan chan bool) {
	for {
		v, ok := <-intChan
		if !ok {
			break
		}
		fmt.Printf("readData 读到数据=%v\n", v)
	}
	//读取完之后就任务完成
	exitChan <- true
	close(exitChan)
}

func main() {
	//创建两个管道
	intChan := make(chan int, 10)
	exitChan := make(chan bool, 1)

	go writeData(intChan)
	go readData(intChan, exitChan)

	//time.Sleep(time.Second *10)
	for {
		_, ok := <-exitChan
		if !ok {
			break
		}
	}
}
