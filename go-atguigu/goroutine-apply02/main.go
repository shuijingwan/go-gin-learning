package main

import (
	"fmt"
	"time"
)

func main() {

	start := time.Now().Unix()
	for num := 1; num <= 200000; num++ {

		flag := true //假设是素数
		//判断num是不是素数
		for i := 2; i < num; i++ {
			if num%i == 0 { //说明该num不是素数
				flag = false
				break
			}
		}

		if flag {
			//将这个数就放入到primeChan
			//primeChan<- num
			//将结果输出
			//fmt.Printf("素数=%d\n", num)
		}

	}
	end := time.Now().Unix()
	fmt.Println("普通的方法耗时=", end-start)

}
