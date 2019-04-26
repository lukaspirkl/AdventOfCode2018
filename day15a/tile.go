package main

import astar "github.com/beefsack/go-astar"

type tile struct {
	x        int
	y        int
	visual   rune
	area     area
	creature *creature
}

func newTile(x, y int, a area) *tile {
	return &tile{
		x:      x,
		y:      y,
		visual: ' ',
		area:   a,
	}
}

func (t *tile) getNeighbour(direction int) *tile {
	switch direction {
	case north:
		return t.area.get(t.x, t.y-1)
	case south:
		return t.area.get(t.x, t.y+1)
	case west:
		return t.area.get(t.x-1, t.y)
	case east:
		return t.area.get(t.x+1, t.y)
	}
	panic("unsupported direction")
}

func (t *tile) getNeighbours(filter tileFilter) []*tile {
	neighbours := []*tile{}
	for i := 0; i < 4; i++ {
		neighbour := t.getNeighbour(i)
		if neighbour != nil && neighbour.visual == '.' && filter(neighbour) {
			neighbours = append(neighbours, neighbour)
		}
	}
	return neighbours
}

func (t *tile) PathNeighbors() []astar.Pather {
	neighbours := t.getNeighbours(tilesWithoutCreature)
	pathers := make([]astar.Pather, len(neighbours))
	for i, neighbour := range neighbours {
		pathers[i] = neighbour
	}
	return pathers
}

func (t *tile) PathNeighborCost(to astar.Pather) float64 {
	return 1
}

func (t *tile) PathEstimatedCost(to astar.Pather) float64 {
	toT := to.(*tile)
	absX := toT.x - t.x
	if absX < 0 {
		absX = -absX
	}
	absY := toT.y - t.y
	if absY < 0 {
		absY = -absY
	}
	return float64(absX + absY)
}

/*
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
*/
