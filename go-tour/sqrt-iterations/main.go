package main

import (
	"fmt"
)

func sqrt(x float64) float64 {
	z := 1.0 // // 初始猜测值
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
		fmt.Println("iteration", i, "z =", z)
	}

	return z
}

func main() {
	fmt.Println("result:", sqrt(2))
}
