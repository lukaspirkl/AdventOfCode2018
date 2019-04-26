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

	count := 0
	for x := 0; x < max.x; x++ {
		for y := 0; y < max.y; y++ {
			sum := 0
			for _, pos := range positions {
				sum += distance(position{x: x, y: y}, pos)
			}
			if sum < 10000 {
				count++
			}
		}
	}

	fmt.Println(count)
}

func distance(a, b position) int {
	return abs(a.x-b.x) + abs(a.y-b.y)
}

func abs(x int) int {
	if x < 0 {
		return -x
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
