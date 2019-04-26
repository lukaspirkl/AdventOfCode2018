package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	real()
}

func real() {
	area := parse("input.txt")
	for i := 0; true; i++ {
		area.moveTrains()
		fmt.Println(i)
	}
}

func test() {
	area := parse("test.txt")
	area.print()
	for {
		area.moveTrains()
		area.print()
	}
}

func parse(fileName string) area {
	fileHandle, _ := os.Open(fileName)
	defer fileHandle.Close()

	maxX, maxY := size(fileHandle)
	area := newArea(maxX, maxY)

	fileHandle.Seek(0, 0)
	scanner := bufio.NewScanner(fileHandle)

	for y := 0; scanner.Scan(); y++ {
		line := scanner.Text()
		for x, r := range []rune(line) {
			area.parseRunes(x, y, r)
		}
	}

	area.connectNeighbours()

	return area
}

func size(fileHandle *os.File) (int, int) {
	scanner := bufio.NewScanner(fileHandle)
	maxX := 0
	maxY := 0
	for scanner.Scan() {
		line := scanner.Text()
		lenX := len([]rune(line))
		if maxX < lenX {
			maxX = lenX
		}
		maxY++
	}
	return maxX, maxY
}
