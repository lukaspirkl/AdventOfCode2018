package main

import "fmt"

type train struct {
	direction             int
	tile                  *tile
	intersectionDirection int
	isCrashed             bool
}

func newTrain(direction int, t *tile) *train {
	return &train{
		direction: direction,
		tile:      t,
	}
}

func (t *train) move() {
	currentTile := t.tile

	goTo := func(direction int) {
		currentTile.train = nil
		destination := currentTile.getNeighbour(direction)
		if destination.train != nil {
			fmt.Printf("CRASH (%d,%d)\n", destination.x, destination.y)
			destination.train.isCrashed = true
			destination.train = nil
			t.isCrashed = true
			return
		}
		destination.train = t
		t.tile = destination
		t.direction = direction
	}

	if currentTile.isIntersection() {
		switch t.intersectionDirection {
		case straight:
			goTo(t.direction)
		case left:
			goTo(t.directionOnLeft())
		case right:
			goTo(t.directionOnRight())
		}
		t.intersectionDirection = nextIntersectionDirection(t.intersectionDirection)
		return
	}
	if !currentTile.isIntersection() && currentTile.isTrackTo(t.direction) {
		goTo(t.direction)
		return
	}
	if !currentTile.isIntersection() && currentTile.isTrackTo(t.directionOnLeft()) {
		goTo(t.directionOnLeft())
		return
	}
	if !currentTile.isIntersection() && currentTile.isTrackTo(t.directionOnRight()) {
		goTo(t.directionOnRight())
		return
	}
}

func (t *train) directionOnLeft() int {
	return (t.direction + 1) % 4
}

func (t *train) directionOnRight() int {
	return (t.direction + 3) % 4
}

func (t *train) String() string {
	switch t.direction {
	case north:
		return string('^')
	case south:
		return string('v')
	case west:
		return string('<')
	case east:
		return string('>')
	default:
		panic("unknown direction")
	}
}
