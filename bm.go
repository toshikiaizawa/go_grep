package main

import (
	"bufio"
	"fmt"
	"os"
)

// 前処理：ずれ表（Bad Character Rule）
func buildBadCharTable(pattern string) [256]int {
	var table [256]int
	for i := 0; i < 256; i++ {
		table[i] = -1
	}
	for i := 0; i < len(pattern); i++ {
		table[pattern[i]] = i
	}
	return table
}

func boyerMooreSearch(text, pattern string, badCharTable [256]int) bool {
	n := len(text)
	m := len(pattern)
	if m == 0 {
		return false
	}
	s := 0
	for s <= n-m {
		j := m - 1
		for j >= 0 && pattern[j] == text[s+j] {
			j--
		}
		if j < 0 {
			return true
		}
		badChar := badCharTable[text[s+j]]
		shift := j - badChar
		if shift < 1 {
			shift = 1
		}
		s += shift
	}
	return false
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run bm.go <pattern>")
		os.Exit(1)
	}
	pattern := os.Args[1]
	badCharTable := buildBadCharTable(pattern)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if boyerMooreSearch(line, pattern, badCharTable) {
			fmt.Println(line)
		}
	}
}

