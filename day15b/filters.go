package main

type creatureFilter func(*creature) bool

func allCreatures(_ *creature) bool {
	return true
}

func onlyCreatures(kind rune) creatureFilter {
	return func(c *creature) bool {
		return c.kind == kind
	}
}

func otherCreaturesThan(kind rune) creatureFilter {
	return func(c *creature) bool {
		return c.kind != kind
	}
}

type tileFilter func(*tile) bool

func allTiles(_ *tile) bool {
	return true
}

func tilesWithoutCreature(c *tile) bool {
	return c.creature == nil
}

func tilesWithCreature(filter creatureFilter) tileFilter {
	return func(t *tile) bool {
		return t.creature != nil && filter(t.creature)
	}
}
