package main

import "fmt"

func main() {
	//1.定义了一个管道10ge数据int
	intChan := make(chan int, 10)
	for i := 0; i < 10; i++ {
		intChan <- i
	}
	stringChan := make(chan string, 5)

	for i := 0; i < 5; i++ {
		stringChan <- "hello" + fmt.Sprintf("%d", i)
	}
	//传统的方法在遍历管道的时候，如果不关闭会阻塞而导致deadlock
	//问题，在实际开发中，我们可能不好确定什么关闭该管道
	//可以使用select方式来解决
	for {
		select {
		//注意：这里如果intChan一直没有
		}
	}
}
