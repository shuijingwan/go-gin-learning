package main

import (
	"fmt"
	"math"
)

// ErrNegativeSqrt 自定义错误类型
type ErrNegativeSqrt float64

// Error 实现 error 接口
// 注意：必须将 e 转换为 float64(e)，否则 fmt.Sprint(e) 会再次调用 Error() 方法，导致无限递归
func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

// Sqrt 计算平方根，如果输入为负数则返回错误
func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	return math.Sqrt(x), nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
