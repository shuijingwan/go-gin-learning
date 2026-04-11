package main

import "golang.org/x/tour/pic"

// Pic 函数：生成 dx * dy 的图片切片
func Pic(dx, dy int) [][]uint8 {
	// 创建外层切片，长度 dy
	image := make([][]uint8, dy)

	// 遍历每一行
	for y := range image {
		// 创建每一行的切片，长度 dx
		row := make([]uint8, dx)

		// 遍历每一列
		for x := range row {
			// 你可以任选下面一种公式，效果不同
			// row[x] = uint8((x + y) / 2)  // 渐变
			// row[x] = uint8(x * y)        // 格子状
			row[x] = uint8(x ^ y) // 斜纹（最漂亮）
		}

		image[y] = row
	}

	return image
}

func main() {
	pic.Show(Pic)
}
