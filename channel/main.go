package main

import "fmt"

func main() {
	var intChan chan int
	intChan = make(chan int, 3)

	fmt.Println(intChan)

	intChan <- 10
	num := 211
	intChan <- num
	intChan <- 50

	fmt.Printf("channel len = %v,cap =%v \n", len(intChan), cap(intChan))
	close(intChan)
	for v := range intChan {
		fmt.Println("v=", v)
	}

	num1 := <-intChan
	num2 := <-intChan
	fmt.Println(num1)
	fmt.Println(num2)
	fmt.Printf("channel len = %v,cap =%v \n", len(intChan), cap(intChan))
}
