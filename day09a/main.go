package main

import (
	"fmt"
)

type marble struct {
	clockwise        *marble
	counterClockwise *marble
	value            int
}

func newZeroMarble() *marble {
	new := &marble{value: 0}
	new.clockwise = new
	new.counterClockwise = new
	return new
}

func (m *marble) place(value int) *marble {
	one := m.clockwise
	two := one.clockwise
	new := &marble{
		counterClockwise: one,
		clockwise:        two,
		value:            value,
	}
	one.clockwise = new
	two.counterClockwise = new
	return new
}

func (m *marble) remove7thCounterClockwise() *marble {
	current := m
	for i := 0; i < 7; i++ {
		current = current.counterClockwise
	}
	counterClockwise := current.counterClockwise
	clockwise := current.clockwise
	counterClockwise.clockwise = clockwise
	clockwise.counterClockwise = counterClockwise
	return current
}

func (m *marble) String() string {
	sprint := func(current *marble) string {
		if m == current {
			return fmt.Sprintf("(%2d)", current.value)
		}
		return fmt.Sprintf("%3d ", current.value)
	}

	zeroMarble := m
	for ; zeroMarble.value != 0; zeroMarble = zeroMarble.clockwise {
	}

	result := sprint(zeroMarble)
	for current := zeroMarble.clockwise; current != zeroMarble; current = current.clockwise {
		result += sprint(current)
	}
	return result
}

const (
	playersCount = 459
	lastMarble   = 71790
)

func main() {
	currentMarble := newZeroMarble()
	//fmt.Printf("%v\n", currentMarble)
	player := 0
	scores := make(map[int]int)
	for marble := 1; marble <= lastMarble; marble++ {
		player++
		if player > playersCount {
			player = 1
		}

		if marble%23 == 0 {
			scores[player] += marble
			removed := currentMarble.remove7thCounterClockwise()
			scores[player] += removed.value
			currentMarble = removed.clockwise

		} else {
			currentMarble = currentMarble.place(marble)
		}

		//fmt.Printf("%v\n", currentMarble)
	}

	max := 0
	for _, score := range scores {
		if max < score {
			max = score
		}
	}
	fmt.Println(max)
}
