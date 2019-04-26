package main

type tile struct {
	x          int
	y          int
	neighbours []bool
	visual     rune
	train      *train
	area       area
}

func newTile(x, y int, a area) *tile {
	return &tile{
		x:          x,
		y:          y,
		neighbours: make([]bool, 4),
		visual:     ' ',
		area:       a,
	}
}

func (t *tile) allowNeighbours(indexes ...int) {
	for _, index := range indexes {
		t.neighbours[index] = true
	}
}

func (t *tile) isTrackTo(directions ...int) bool {
	for _, direction := range directions {
		if !t.neighbours[direction] {
			return false
		}
	}
	return true
}

func (t *tile) isIntersection() bool {
	sum := 0
	for _, value := range t.neighbours {
		if value {
			sum++
		}
	}
	return sum > 2
}

func (t *tile) isNeighbourOut(direction int) bool {
	switch direction {
	case north:
		if t.area.isOut(t.x, t.y-1) {
			return true
		}
	case south:
		if t.area.isOut(t.x, t.y+1) {
			return true
		}
	case west:
		if t.area.isOut(t.x-1, t.y) {
			return true
		}
	case east:
		if t.area.isOut(t.x+1, t.y) {
			return true
		}
	}
	return false
}

func (t *tile) getNeighbours() (north, west, south, east *tile) {
	north = t.area.get(t.x, t.y-1)
	south = t.area.get(t.x, t.y+1)
	west = t.area.get(t.x-1, t.y)
	east = t.area.get(t.x+1, t.y)
	return
}

func (t *tile) getNeighbour(direction int) *tile {
	if t.isNeighbourOut(direction) {
		return nil
	}
	switch direction {
	case north:
		return t.area.get(t.x, t.y-1)
	case south:
		return t.area.get(t.x, t.y+1)
	case west:
		return t.area.get(t.x-1, t.y)
	case east:
		return t.area.get(t.x+1, t.y)
	default:
		return nil
	}
}

func (t *tile) isAccessibleFrom(directions ...int) bool {
	for _, direction := range directions {
		neighbour := t.getNeighbour(direction)
		if neighbour == nil {
			return false
		}
		if !neighbour.isTrackTo(oppositDirection(direction)) {
			return false
		}
	}
	return true
}
