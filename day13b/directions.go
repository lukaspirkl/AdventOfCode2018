package main

// neighbours
//   0
//  1X3
//   2

const (
	north = 0
	west  = 1
	south = 2
	east  = 3
)

const (
	left     = 0
	straight = 1
	right    = 2
)

func nextIntersectionDirection(current int) int {
	return (current + 1) % 3
}

func oppositDirection(direction int) int {
	return (direction + 2) % 4
}
