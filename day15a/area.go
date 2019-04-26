package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

type area [][]*tile

func newArea(rs io.ReadSeeker) area {
	size := func(reader io.Reader) (int, int) {
		scanner := bufio.NewScanner(reader)
		maxX := 0
		maxY := 0
		for scanner.Scan() {
			lenX := len([]rune(scanner.Text()))
			if maxX < lenX {
				maxX = lenX
			}
			maxY++
		}
		return maxX, maxY
	}
	makeEmpty := func(maxX, maxY int) area {
		area := make([][]*tile, maxY)
		for y := range area {
			area[y] = make([]*tile, maxX)
			for x := range area[y] {
				area[y][x] = newTile(x, y, area)
			}
		}
		return area
	}
	parse := func(reader io.Reader, a area) {
		scanner := bufio.NewScanner(reader)
		for y := 0; scanner.Scan(); y++ {
			line := scanner.Text()
			for x, r := range []rune(line) {
				a.parseRunes(x, y, r)
			}
		}
	}

	rs.Seek(0, 0)
	maxX, maxY := size(rs)
	area := makeEmpty(maxX, maxY)
	rs.Seek(0, 0)
	parse(rs, area)
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
	if r == 'G' || r == 'E' {
		tile.visual = '.'
		tile.creature = newCreature(tile, r)
	} else {
		tile.visual = r
	}
}

func (a area) turn() bool {
	creatures := a.getCreatures(allCreatures)
	for _, c := range creatures {
		gnomes := a.sumOfHitPoints(onlyCreatures('G'))
		elfs := a.sumOfHitPoints(onlyCreatures('E'))
		if (gnomes <= 0 && elfs > 0) || (elfs <= 0 && gnomes > 0) {
			return false
		}
		c.turn()
	}
	return true
}

func (a area) getCreatures(filter creatureFilter) []*creature {
	creatures := []*creature{}
	for _, row := range a {
		for _, tile := range row {
			if tile.creature != nil && filter(tile.creature) {
				creatures = append(creatures, tile.creature)
			}
		}
	}
	return creatures
}

func (a area) sumOfHitPoints(filter creatureFilter) int {
	sum := 0
	for _, creature := range a.getCreatures(filter) {
		sum += creature.hitPoints
	}
	return sum
}

func (a area) run(callback func(int, area)) int {
	round := 0
	for {
		round++
		if !a.turn() {
			return round - 1
		}
		if callback != nil {
			callback(round, a)
		}
	}
}

func (a area) String() string {
	sb := strings.Builder{}
	for i, row := range a {
		creatures := []*creature{}
		for _, tile := range row {
			if tile.creature == nil {
				sb.WriteRune(tile.visual)
			} else {
				creatures = append(creatures, tile.creature)
				sb.WriteString(tile.creature.String())
			}
		}
		sb.WriteString("  ")
		for _, c := range creatures {
			sb.WriteString(fmt.Sprintf("%s(%d) ", string(c.kind), c.hitPoints))
		}
		if i < len(a)-1 {
			sb.WriteString("\n")
		}
	}
	return sb.String()
}
