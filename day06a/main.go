package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const maxInt = int(^uint(0) >> 1)

type position struct {
	x int
	y int
}

func main() {
	positions := getPositions()
	min := findMin(positions)
	positions = subtractAll(positions, min)
	max := findMax(positions)
	area := createArea(max)

	placeLevel0(area, positions)
	print(area)
	fmt.Println()

	isSomethingMarked := true
	for i := 0; isSomethingMarked; i++ {
		area, isSomethingMarked = expand(area)
		//print(area)
		fmt.Println(i)
	}

	fmt.Println("----")
	fmt.Println(topCount(area))
}

func topCount(area [][]string) int {
	counts := make(map[int]int)
	for _, ys := range area {
		for _, value := range ys {
			if !strings.Contains(value, "-") {
				continue
			}
			data := strings.Split(value, "-")
			id := atoi(data[0])
			counts[id]++
		}
	}
	max := 0
	for _, count := range counts {
		if max < count {
			max = count
		}
	}
	return max
}

func markPosition(area [][]string, pos position, id int, level int) bool {
	if pos.y < 0 || pos.x < 0 || pos.x > len(area)-1 || pos.y > len(area[0])-1 {
		return false
	}
	currentValue := area[pos.x][pos.y]
	if strings.Contains(currentValue, "X") {
		return false
	}
	if strings.Contains(currentValue, ".") {
		area[pos.x][pos.y] = fmt.Sprintf("%v-%v", id, level)
		return true
	}

	if strings.Contains(currentValue, "-") {
		data := strings.Split(currentValue, "-")
		if id != atoi(data[0]) && level == atoi(data[1]) {
			area[pos.x][pos.y] = " X "
			return true
		}
	}

	return false
}

func expand(area [][]string) ([][]string, bool) {

	newArea := make([][]string, len(area))
	for i := range area {
		newArea[i] = make([]string, len(area[i]))
		copy(newArea[i], area[i])
	}

	isSomethingMarked := false

	for x, ys := range area {
		for y, value := range ys {
			if !strings.Contains(value, "-") {
				continue
			}
			data := strings.Split(value, "-")
			id := atoi(data[0])
			level := atoi(data[1]) + 1

			isSomethingMarked = markPosition(newArea, position{x: x - 1, y: y}, id, level) || isSomethingMarked
			isSomethingMarked = markPosition(newArea, position{x: x + 1, y: y}, id, level) || isSomethingMarked
			isSomethingMarked = markPosition(newArea, position{x: x, y: y - 1}, id, level) || isSomethingMarked
			isSomethingMarked = markPosition(newArea, position{x: x, y: y + 1}, id, level) || isSomethingMarked
		}
	}

	return newArea, isSomethingMarked
}

func placeLevel0(area [][]string, positions []position) {
	for i, pos := range positions {
		area[pos.x][pos.y] = fmt.Sprintf("%v-0", i)
	}
}

func print(area [][]string) {
	for _, ys := range area {
		for _, value := range ys {
			fmt.Printf("%v ", value)
		}
		fmt.Println()
	}
}

func createArea(size position) [][]string {
	x := make([][]string, size.x+1)
	for i := range x {
		x[i] = make([]string, size.y+1)
		for j := range x[i] {
			x[i][j] = " . "
		}
	}
	return x
}

func subtractAll(positions []position, distance position) []position {
	newPositions := []position{}
	for _, oldPosition := range positions {
		newPositions = append(newPositions, position{x: oldPosition.x - distance.x, y: oldPosition.y - distance.y})
	}
	return newPositions
}

func findMin(positions []position) position {
	min := position{x: maxInt, y: maxInt}
	for _, position := range positions {
		if min.x > position.x {
			min.x = position.x
		}
		if min.y > position.y {
			min.y = position.y
		}
	}
	return min
}

func findMax(positions []position) position {
	max := position{x: 0, y: 0}
	for _, position := range positions {
		if max.x < position.x {
			max.x = position.x
		}
		if max.y < position.y {
			max.y = position.y
		}
	}
	return max
}

func getPositions() []position {
	fileHandle, _ := os.Open("input.txt")
	defer fileHandle.Close()

	positions := []position{}
	scanner := bufio.NewScanner(fileHandle)
	for scanner.Scan() {
		coords := strings.Split(scanner.Text(), ", ")
		positions = append(positions, position{x: atoi(coords[0]), y: atoi(coords[1])})
	}
	return positions
}

func atoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}
