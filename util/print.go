package util

import (
	"fmt"
	"strings"
)

func PrintUpdLine(os string, ln int, s string) {
	move(os, ln, s)
}

func move(os string, ln int, s string) {
	fmt.Print(os)
	lineTotal := countLF(os) + 1
	revLn := lineTotal - ln
	// 光标上移 revLn && 清除从光标到行尾 && 重新输入s
	rs := fmt.Sprintf("\033[%dA\r\033[K%s", revLn, s)
	// 光标下移 回归到底部
	rs += fmt.Sprintf("\033[%dB", revLn)
	fmt.Println(rs)
}

func countLF(s string) int {
	return strings.Count(s, "\n")
}
