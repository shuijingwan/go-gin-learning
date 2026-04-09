package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0              // 初始猜测值
	const epsilon = 1e-10 // 精度控制

	for {
		prev := z
		z -= (z*z - x) / (2 * z)

		// 判断变化是否足够小
		if math.Abs(z-prev) < epsilon {
			break
		}
	}

	return z
}

func main() {
	fmt.Println("My sqrt:", Sqrt(2))
	fmt.Println("Math sqrt:", math.Sqrt(2))
}
