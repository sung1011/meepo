package util

import (
	"fmt"
	"strings"
)

func Print() {
	os := "111\n222\n333\n444\n555\n666\n777\n888\n999\n"
	move(os, 3, "la")
}

func move(os string, line int, s string) {
	fmt.Print(os)
	lineTotal := countLF(os) + 1
	line = lineTotal - line
	fmt.Printf("\033[%dA\r\033[K%s", line, s)
	fmt.Printf("\033[%dB", line)
}

func countLF(s string) int {
	return strings.Count(s, "\n")
}
