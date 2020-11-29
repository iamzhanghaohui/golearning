package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("d:/train.txt")
	if err != nil {
		fmt.Println("open file err=", err)
	}

	//当函数退出时，要及时的关闭file

	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadSlice('\n')
		if err == io.EOF {
			break
		}
		fmt.Println(str)
	}

	fmt.Println("文件读取结束")
}
