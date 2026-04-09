package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0              // 初始猜测值
	const epsilon = 1e-10 // 精度控制

	count := 0

	for {
		prev := z
		z -= (z*z - x) / (2 * z)

		count++

		// 判断变化是否足够小
		if math.Abs(z-prev) < epsilon {
			break
		}
	}

	fmt.Println("iterations:", count)

	return z
}

func main() {
	fmt.Println("result:", Sqrt(2))
}
