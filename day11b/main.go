package main

import (
	"fmt"
	"runtime"
)

const serialNumber int = 1723

func main() {

	maxValue := 0
	maxX := 0
	maxY := 0
	maxSize := 0
	for size := 1; size < 300; size++ {
		for x := 1; x <= 300-size; x++ {
			for y := 1; y <= 300-size; y++ {
				value := calculateSquare(x, y, size)
				if maxValue < value {
					maxValue = value
					maxX = x
					maxY = y
					maxSize = size
				}
			}
		}
		runtime.GC()
		fmt.Println(size)
	}
	fmt.Printf("%v,%v,%v has total power %v", maxX, maxY, maxSize+1, maxValue)
}

func calculateSquare(x, y, size int) int {
	sum := 0
	for xx := 0; xx <= size; xx++ {
		for yy := 0; yy <= size; yy++ {
			sum += calculate(x+xx, y+yy)
		}
	}
	return sum
}

func calculate(x, y int) int {
	rackID := x + 10
	num := ((rackID * y) + serialNumber) * rackID
	return (num / 100 % 10) - 5
}
