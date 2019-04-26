package main

import "fmt"

const serialNumber int = 1723

func main() {

	maxValue := 0
	maxX := 0
	maxY := 0
	for x := 1; x <= 300-2; x++ {
		for y := 1; y <= 300-2; y++ {
			value := calculateSquare(x, y)
			if maxValue < value {
				maxValue = value
				maxX = x
				maxY = y
			}
		}
	}

	fmt.Printf("%v,%v has total power %v", maxX, maxY, maxValue)
}

func calculateSquare(x, y int) int {
	sum := 0
	for xx := 0; xx <= 2; xx++ {
		for yy := 0; yy <= 2; yy++ {
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
