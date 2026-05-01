package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// Walk 中序遍历树 t，将所有值发送到通道 ch
func Walk(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}
	// 递归遍历左子树
	Walk(t.Left, ch)
	// 发送当前节点的值到通道
	ch <- t.Value
	// 递归遍历右子树
	Walk(t.Right, ch)
}

// Same 判断树 t1 和 t2 是否包含相同的值
func Same(t1, t2 *tree.Tree) bool {
	// 创建两个通道用于接收树的值
	ch1 := make(chan int)
	ch2 := make(chan int)

	// 启动两个 goroutine 分别遍历两棵树
	go func() {
		Walk(t1, ch1)
		close(ch1) // 关闭通道，表示遍历完成
	}()

	go func() {
		Walk(t2, ch2)
		close(ch2) // 关闭通道，表示遍历完成
	}()

	// 同时从两个通道读取值并比较
	for {
		v1, ok1 := <-ch1
		v2, ok2 := <-ch2

		// 如果两个通道都关闭了，说明遍历完成且所有值都相等
		if !ok1 && !ok2 {
			return true
		}

		// 如果一个通道关闭而另一个没有，或者值不相等，返回 false
		if ok1 != ok2 || v1 != v2 {
			return false
		}
	}
}

func main() {
	// 测试 Walk 函数
	fmt.Println("测试 Walk 函数：")
	ch := make(chan int)
	go Walk(tree.New(1), ch)
	for i := 0; i < 10; i++ {
		fmt.Print(<-ch, " ")
	}
	fmt.Println()

	// 测试 Same 函数
	fmt.Println("\n测试 Same 函数：")
	fmt.Println("Same(tree.New(1), tree.New(1)) =", Same(tree.New(1), tree.New(1)))
	fmt.Println("Same(tree.New(1), tree.New(2)) =", Same(tree.New(1), tree.New(2)))
}
