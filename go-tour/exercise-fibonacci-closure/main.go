package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
// fibonacci 函数返回一个闭包，用于生成连续的斐波那契数
func fibonacci() func() int {
	// 闭包会保留这两个变量的状态
	a, b := 0, 1

	return func() int {
		// 先保存当前 a 的值（要返回的值）
		res := a
		// 更新 a 和 b 到下一个状态
		a, b = b, a+b
		// 返回结果
		return res
	}
}

func main() {
	f := fibonacci() // 创建一个闭包实例
	// 循环调用 10 次，输出前 10 个斐波那契数
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
