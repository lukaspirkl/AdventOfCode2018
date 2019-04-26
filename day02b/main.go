package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fileHandle, _ := os.Open("input.txt")
	defer fileHandle.Close()

	ids := []string{}

	scanner := bufio.NewScanner(fileHandle)
	for scanner.Scan() {
		current := scanner.Text()
		for _, id := range ids {
			if isOneLetterDifference(current, id) {
				fmt.Println(id)
				fmt.Println(current)
			}
		}
		ids = append(ids, current)
	}
}

func isOneLetterDifference(a string, b string) bool {
	if len(a) != len(b) {
		panic("strings have different length")
	}
	sameCount := 0
	for index := range a {
		if b[index] == a[index] {
			sameCount++
		}
	}
	return sameCount+1 == len(a)
}
