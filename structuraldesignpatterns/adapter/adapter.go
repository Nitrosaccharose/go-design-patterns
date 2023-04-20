package adapter

import (
	"strconv"
	"strings"
)

// Reverser 反转器接口
type Reverser interface {
	reverse(str string) string
}

// StringReverser 字符串反转器
type StringReverser struct{}

func (sr *StringReverser) reverse(str string) string {
	var result []string
	for _, c := range str {
		result = append([]string{string(c)}, result...)
	}
	return strings.Join(result, "")
}

// NumberAdapter 数字适配器
type NumberAdapter struct {
	reverser Reverser
}

func (n *NumberAdapter) reverse(num int) int {
	// 将数字转换为字符串
	numStr := strconv.Itoa(num)
	// 调用 StringReverser 的 reverse 方法
	reversedStr := n.reverser.reverse(numStr)
	// 将字符串转换为数字
	reversedNum, _ := strconv.Atoi(reversedStr)
	return reversedNum
}
