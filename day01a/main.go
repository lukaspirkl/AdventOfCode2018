package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	fileHandle, _ := os.Open("input.txt")
	defer fileHandle.Close()

	scanner := bufio.NewScanner(fileHandle)
	sum := 0
	for scanner.Scan() {
		text := scanner.Text()

		number, err := strconv.Atoi(text)
		if err != nil {
			panic(err)
		}
		sum += number
	}
	fmt.Println(sum)
}
