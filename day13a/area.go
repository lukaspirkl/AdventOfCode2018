package main

import "fmt"

type area [][]*tile

func newArea(maxX, maxY int) area {
	area := make([][]*tile, maxY)
	for y := range area {
		area[y] = make([]*tile, maxX)
		for x := range area[y] {
			area[y][x] = newTile(x, y, area)
		}
	}
	return area
}

func (a area) isOut(x, y int) bool {
	return x < 0 || y < 0 || y >= len(a) || x >= len(a[0])
}

func (a area) get(x, y int) *tile {
	if a.isOut(x, y) {
		return nil
	}
	return a[y][x]
}

func (a area) parseRunes(x, y int, r rune) {
	tile := a.get(x, y)
	switch r {
	case '<':
		tile.visual = '-'
		tile.train = newTrain(west, tile)
	case '>':
		tile.visual = '-'
		tile.train = newTrain(east, tile)
	case '^':
		tile.visual = '|'
		tile.train = newTrain(north, tile)
	case 'v':
		tile.visual = '|'
		tile.train = newTrain(south, tile)
	default:
		tile.visual = r
	}
}

func (a area) connectNeighbours() {
	for _, row := range a {
		for _, tile := range row {
			switch tile.visual {
			case '-':
				tile.allowNeighbours(east, west)
			case '|':
				tile.allowNeighbours(north, south)
			case '+':
				tile.allowNeighbours(east, west, north, south)
			case '/':
				n, w, s, e := tile.getNeighbours()
				if n != nil && w != nil && (n.visual == '|' || n.visual == '+') && (w.visual == '-' || w.visual == '+') {
					tile.allowNeighbours(north, west)
				} else if s != nil && e != nil && (s.visual == '|' || s.visual == '+') && (e.visual == '-' || e.visual == '+') {
					tile.allowNeighbours(south, east)
				} else {
					panic(fmt.Sprintf("imposible curve track '/' on (%d,%d)", tile.x, tile.y))
				}
			case '\\':
				n, w, s, e := tile.getNeighbours()
				if n != nil && e != nil && (n.visual == '|' || n.visual == '+') && (e.visual == '-' || e.visual == '+') {
					tile.allowNeighbours(north, east)
				} else if s != nil && w != nil && (s.visual == '|' || s.visual == '+') && (w.visual == '-' || w.visual == '+') {
					tile.allowNeighbours(south, west)
				} else {
					panic(fmt.Sprintf("imposible curve track '\\' on (%d,%d)", tile.x, tile.y))
				}
			}
		}
	}
}

func (a area) moveTrains() {
	trains := []*train{}
	for _, row := range a {
		for _, tile := range row {
			if tile.train != nil {
				trains = append(trains, tile.train)
			}
		}
	}
	for _, train := range trains {
		train.move()
	}
}

func (a area) print() {
	for _, row := range a {
		for _, tile := range row {
			if tile.train == nil {
				fmt.Print(string(tile.visual))
			} else {
				fmt.Print(tile.train)
			}

		}
		fmt.Println()
	}
}
