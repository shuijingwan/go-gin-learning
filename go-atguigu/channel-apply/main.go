package main

import (
	"fmt"
	// 此处无需引入time包，彻底摆脱手动休眠的依赖
)

// writeData 协程：向intChan写入50个整数（生产者）
func writeData(intChan chan int) {
	// 循环写入50个整数，逻辑简单但需注意循环边界
	for i := 1; i <= 50; i++ {
		// 向管道写入数据，有缓冲管道会自动调节写入节奏
		intChan <- i
		fmt.Println("writeData ", i)
		// 此处可注释time.Sleep，验证Channel的缓冲特性，无需手动控制写入速度
		// time.Sleep(time.Second)
	}
	// 关键操作：写入完成后关闭管道，告知readData协程“数据已写完”
	close(intChan)
}

// readData 协程：从intChan读取数据，读取完成后向exitChan发送信号（消费者）
func readData(intChan chan int, exitChan chan bool) {
	// 无限循环读取管道数据，直到管道关闭且无数据可读
	for {
		// 核心语法：v接收数据，ok判断管道是否关闭（true=有数据/未关闭，false=管道关闭且无数据）
		v, ok := <-intChan
		if !ok { // 管道关闭且无数据，说明读取完成，退出循环
			break
		}
		fmt.Printf("readData 读到数据=%v\n", v)
	}
	// readData 读取完数据后，即任务完成
	exitChan <- true
	// 关闭exitChan（可选，此处关闭是为了让主线程读取信号后正常退出，避免阻塞）
	close(exitChan)
}

func main() {
	// 创建两个管道
	// 1. 创建数据管道intChan：有缓冲管道，容量设为10，平衡读写节奏
	// 复盘：容量可根据实际需求调整，此处10足够承载writeData的写入速度，避免频繁阻塞
	intChan := make(chan int, 10)
	// 2. 创建信号管道exitChan：用于传递“协程完成”信号，容量设为1即可（仅需传递1个信号）
	exitChan := make(chan bool, 1)

	// 启动两个协程，共享同一个intChan，实现协同工作
	go writeData(intChan)
	go readData(intChan, exitChan)

	// 主线程监听exitChan，等待readData协程发送完成信号，精准退出
	// 复盘：此处替代了此前的time.Sleep，彻底解决休眠时间估算不准的问题
	for {
		_, ok := <-exitChan
		if !ok { // exitChan关闭且无信号，说明所有协程都已完成
			break
		}
	}
	// 主线程退出前可添加提示，验证协同效果
	fmt.Println("所有协程执行完成，主线程正常退出")
}
