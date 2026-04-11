package main

import (
	"strings" // 导入 strings 包，使用 Fields 函数分割单词

	"golang.org/x/tour/wc"
)

// WordCount 统计字符串中每个单词的出现次数
func WordCount(s string) map[string]int {
	// 1. 创建一个空的 map，用于存储单词和对应的计数
	wordMap := make(map[string]int)

	// 2. 使用 strings.Fields 分割字符串，按空格/制表符/换行符分割成单词切片
	words := strings.Fields(s)

	// 3. 遍历单词切片，统计每个单词的数量
	for _, word := range words {
		wordMap[word]++ // 单词每出现一次，对应计数 +1
	}

	// 4. 返回统计结果
	return wordMap
}

func main() {
	// 运行测试函数，验证 WordCount 是否正确
	wc.Test(WordCount)
}
