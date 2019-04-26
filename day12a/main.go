package main

import (
	"fmt"
	"time"
)

type pots map[int]bool

func newPots(s string, zeroIndex int) pots {
	m := make(pots)
	for i, r := range []rune(s) {
		if r == '#' {
			m[i-zeroIndex] = true
		} else {
			m[i-zeroIndex] = false
		}
	}
	return m
}

func (p pots) minIndex() int {
	min := len(p)
	for key := range p {
		if min > key {
			min = key
		}
	}
	return min
}

func (p pots) String() string {
	min := p.minIndex()
	str := ""
	for i := min; i < len(p)+min; i++ {
		if p[i] {
			str += "#"
		} else {
			str += "."
		}
	}
	return str
}

func (p pots) trim() {
	min := p.minIndex()
	firstFull := min
	for i := min; i < len(p)+min; i++ {
		value, _ := p[i]
		if value {
			firstFull = i
			break
		}
	}
	lastFull := len(p) + min
	for i := len(p) + min; i > min; i-- {
		value, _ := p[i]
		if value {
			lastFull = i
			break
		}
	}
	for key := range p {
		if key < firstFull || key > lastFull {
			delete(p, key)
		}
	}
}

func (p pots) next(rules []pots) pots {
	next := make(pots)
	min := p.minIndex()
	for i := min - 5; i < len(p)+min+5; i++ {
		for _, rule := range rules {
			if p.isMatch(i, rule) {
				next[i] = true
				break
			} else {
				next[i] = false
			}
		}
	}
	next.trim()
	return next
}

func (p pots) isMatch(i int, rule pots) bool {
	match := false
	for j := -2; j <= 2; j++ {
		potValue, present := p[i+j]
		if !present {
			potValue = false
		}
		ruleValue, _ := rule[j]
		if potValue == ruleValue {
			match = true
		} else {
			match = false
			break
		}
	}
	return match
}

func (p pots) sum() int {
	sum := 0
	for key, value := range p {
		if value {
			sum += key
		}
	}
	return sum
}

func main() {
	start := time.Now()
	runWithRealData(20)
	fmt.Println(time.Since(start))
}

func runWithRealData(iterations int) {
	state := newPots("#.##.##.##.##.......###..####..#....#...#.##...##.#.####...#..##..###...##.#..#.##.#.#.#.#..####..#", 0)
	rules := []pots{
		newPots("##..#", 2),
		newPots("##...", 2),
		newPots("###.#", 2),
		newPots("..##.", 2),
		newPots(".##.#", 2),
		newPots("#..#.", 2),
		newPots(".##..", 2),
		newPots("###..", 2),
		newPots(".###.", 2),
		newPots("#####", 2),
		newPots("...#.", 2),
		newPots(".#...", 2),
		newPots("#.#.#", 2),
		newPots(".#.##", 2),
		newPots("..#.#", 2),
		newPots("#.#..", 2),
	}
	makeItSo(state, rules, iterations)
}

func runWithTestData() {
	state := newPots("#..#.#..##......###...###", 0)
	rules := []pots{
		newPots("...##", 2),
		newPots("..#..", 2),
		newPots(".#...", 2),
		newPots(".#.#.", 2),
		newPots(".#.##", 2),
		newPots(".##..", 2),
		newPots(".####", 2),
		newPots("#.#.#", 2),
		newPots("#.###", 2),
		newPots("##.#.", 2),
		newPots("##.##", 2),
		newPots("###..", 2),
		newPots("###.#", 2),
		newPots("####.", 2),
	}
	makeItSo(state, rules, 20)
}

func makeItSo(state pots, rules []pots, iterations int) {
	fmt.Printf("%3d. %v (%d)\n", 0, state, state.sum())
	for i := 1; i <= iterations; i++ {
		state = state.next(rules)
		fmt.Printf("%3d. %v (%d)\n", i, state, state.sum())
	}
	fmt.Printf("sum is: %d\n", state.sum())
}
