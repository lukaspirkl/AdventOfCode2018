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

	m := make(map[int]bool)
	sum := 0
	m[0] = true

	for {
		fileHandle.Seek(0, 0)
		scanner := bufio.NewScanner(fileHandle)
		for scanner.Scan() {
			text := scanner.Text()

			number, err := strconv.Atoi(text)
			if err != nil {
				panic(err)
			}

			sum += number
			fmt.Printf("%d = %d\n", number, sum)

			if _, present := m[sum]; present {
				fmt.Println(sum)
				return
			}

			m[sum] = true
		}
		fmt.Println("repeat")
	}
}
