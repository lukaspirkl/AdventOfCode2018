package main

import (
	"math"
	"sort"

	"github.com/beefsack/go-astar"
)

const (
	startHitPoints = 200
)

type creature struct {
	kind        rune
	tile        *tile
	attackPower int
	hitPoints   int
	dead        bool
}

func newCreature(tile *tile, kind rune, attackPower int) *creature {
	return &creature{
		tile:        tile,
		kind:        kind,
		attackPower: attackPower,
		hitPoints:   startHitPoints,
	}
}

func (c *creature) String() string {
	return string(c.kind)
}

func (c *creature) turn() {
	if c.dead {
		return
	}

	//TODO: Improve this quick fix for the rule I missed at the first sight: attack and stop or move and also attack
	enemyTiles := c.tile.getNeighbours(tilesWithCreature(otherCreaturesThan(c.kind)))
	if len(enemyTiles) >= 1 {
		c.attack(enemyTiles)
		return
	}

	spaces := c.tile.getNeighbours(tilesWithoutCreature)
	if len(spaces) >= 1 {
		c.move(c.calculateNextMove())
	}

	enemyTiles = c.tile.getNeighbours(tilesWithCreature(otherCreaturesThan(c.kind)))
	if len(enemyTiles) >= 1 {
		c.attack(enemyTiles)
		return
	}
}

func (c *creature) move(target *tile) bool {
	if target == nil {
		return false
	}
	c.tile.creature = nil
	target.creature = c
	c.tile = target
	return true
}

func (c *creature) calculateNextMove() *tile {
	tilesNextToEnemy := []*tile{}
	for _, enemy := range c.tile.area.getCreatures(otherCreaturesThan(c.kind)) {
		tilesNextToEnemy = append(tilesNextToEnemy, enemy.tile.getNeighbours(tilesWithoutCreature)...)
	}

	sort.Slice(tilesNextToEnemy, func(i, j int) bool {
		if tilesNextToEnemy[i].y < tilesNextToEnemy[j].y {
			return true
		} else if tilesNextToEnemy[i].y > tilesNextToEnemy[j].y {
			return false
		} else {
			return tilesNextToEnemy[i].x < tilesNextToEnemy[j].x
		}
	})

	shortestDistance := math.MaxFloat64
	var targetTile *tile
	for _, t := range tilesNextToEnemy {
		path, distance, found := astar.Path(c.tile, t)
		if found && shortestDistance > distance {
			shortestDistance = distance
			targetTile = path[0].(*tile)
		}
	}
	if targetTile == nil {
		return nil
	}

	shortestDistance = math.MaxFloat64
	var nextTile *tile
	for _, t := range c.tile.getNeighbours(tilesWithoutCreature) {
		path, distance, found := astar.Path(targetTile, t)
		if found && shortestDistance > distance {
			shortestDistance = distance
			nextTile = path[0].(*tile)
		}
	}
	return nextTile
}

func (c *creature) attack(targetOnTiles []*tile) {
	weakestIndex := 0
	weakestHitPoints := startHitPoints
	for i := 0; i < len(targetOnTiles); i++ {
		if weakestHitPoints > targetOnTiles[i].creature.hitPoints {
			weakestIndex = i
			weakestHitPoints = targetOnTiles[i].creature.hitPoints
		}
	}
	target := targetOnTiles[weakestIndex].creature
	//fmt.Printf("%s(%d,%d)[%dHP] attacks %s(%d,%d)[%dHP]\n", string(c.kind), c.tile.x, c.tile.y, c.hitPoints, string(target.kind), target.tile.x, target.tile.y, target.hitPoints)
	target.hitPoints -= c.attackPower
	if target.hitPoints <= 0 {
		target.dead = true
		target.tile.creature = nil
		//fmt.Printf("%s(%d,%d)[%dHP] dead\n", string(target.kind), target.tile.x, target.tile.y, target.hitPoints)
	} else {
		//fmt.Printf("%s(%d,%d)[%dHP] hurts\n", string(target.kind), target.tile.x, target.tile.y, target.hitPoints)
	}
}
