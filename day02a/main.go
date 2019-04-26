package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fileHandle, _ := os.Open("input.txt")
	defer fileHandle.Close()

	countTwo := 0
	countThree := 0

	scanner := bufio.NewScanner(fileHandle)
	for scanner.Scan() {
		m := createMap(scanner.Text())
		if isExactCount(m, 2) {
			countTwo++
		}
		if isExactCount(m, 3) {
			countThree++
		}
	}

	fmt.Println(countTwo * countThree)
}

func createMap(text string) map[rune]int {
	m := make(map[rune]int)

	for _, r := range text {
		if _, present := m[r]; present {
			m[r]++
		} else {
			m[r] = 1
		}
	}

	return m
}

func isExactCount(m map[rune]int, count int) bool {
	for _, value := range m {
		if value == count {
			return true
		}
	}
	return false
}
