package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

type pots struct {
	data         []bool
	zeroPosition int
}

func newPots(s string) pots {
	p := pots{
		zeroPosition: 0,
		data:         make([]bool, len([]rune(s))),
	}
	for i, r := range []rune(s) {
		if r == '#' {
			p.data[i] = true
		} else {
			p.data[i] = false
		}
	}
	return p
}

func (p pots) String() string {
	str := ""
	for _, value := range p.data {
		if value {
			str += "#"
		} else {
			str += "."
		}
	}
	return str
}

func (p *pots) next(rules map[int]struct{}) {
	p.ensureFiveFree()
	new := make([]bool, len(p.data))
	for i := 2; i < len(p.data)-2; i++ {
		slice := p.data[i-2 : i+3]
		if _, present := rules[boolSliceToInt(slice)]; present {
			new[i] = true
		}
	}
	p.data = new
}

func (p *pots) ensureFiveFree() {
	freeEnd := 0
	for ; freeEnd < len(p.data); freeEnd++ {
		if p.data[len(p.data)-freeEnd-1] {
			break
		}
	}
	if freeEnd < 5 {
		p.data = append(p.data, make([]bool, 5-freeEnd)...)
	} else if freeEnd > 5 {
		p.data = p.data[:len(p.data)-freeEnd-5]
	}

	freeBegin := 0
	for ; freeBegin < len(p.data); freeBegin++ {
		if p.data[freeBegin] {
			break
		}
	}
	if freeBegin < 5 {
		p.data = append(make([]bool, 5-freeBegin), p.data...)
		p.zeroPosition += 5 - freeBegin
	} else if freeBegin > 5 {
		p.data = p.data[freeBegin-5:]
		p.zeroPosition += 5 - freeBegin
	}
	//fmt.Printf("begin: %v end: %v\n", freeBegin, freeEnd)
}

func (p pots) sum() int {
	sum := 0
	for index, value := range p.data {
		if value {
			sum += index - p.zeroPosition
		}
	}
	return sum
}

func (p pots) sumPositionLess() int {
	sum := 0
	for index, value := range p.data {
		if value {
			sum += index
		}
	}
	return sum
}

func ruleToNum(s string) int {
	i, err := strconv.ParseInt(strings.Replace(strings.Replace(s, "#", "1", -1), ".", "0", -1), 2, 8)
	if err != nil {
		panic(err)
	}
	return int(i)
}

func boolSliceToInt(b []bool) int {
	r := 0
	for i := 0; i < len(b); i++ {
		if b[i] {
			r |= 1 << (uint(len(b) - 1 - i))
		}
	}
	return r
}

func main() {

	//runWithTestData()

	start := time.Now()
	runWithRealData(50000000000)
	fmt.Println(time.Since(start))
}

func runWithRealData(iterations int) {
	state := newPots("#.##.##.##.##.......###..####..#....#...#.##...##.#.####...#..##..###...##.#..#.##.#.#.#.#..####..#")
	rules := make(map[int]struct{})
	rules[ruleToNum("##..#")] = struct{}{}
	rules[ruleToNum("##...")] = struct{}{}
	rules[ruleToNum("###.#")] = struct{}{}
	rules[ruleToNum("..##.")] = struct{}{}
	rules[ruleToNum(".##.#")] = struct{}{}
	rules[ruleToNum("#..#.")] = struct{}{}
	rules[ruleToNum(".##..")] = struct{}{}
	rules[ruleToNum("###..")] = struct{}{}
	rules[ruleToNum(".###.")] = struct{}{}
	rules[ruleToNum("#####")] = struct{}{}
	rules[ruleToNum("...#.")] = struct{}{}
	rules[ruleToNum(".#...")] = struct{}{}
	rules[ruleToNum("#.#.#")] = struct{}{}
	rules[ruleToNum(".#.##")] = struct{}{}
	rules[ruleToNum("..#.#")] = struct{}{}
	rules[ruleToNum("#.#..")] = struct{}{}
	fmt.Println(state)
	makeItSo(state, rules, iterations)
}

func runWithTestData() {
	state := newPots("#..#.#..##......###...###")
	rules := make(map[int]struct{})
	rules[ruleToNum("...##")] = struct{}{}
	rules[ruleToNum("..#..")] = struct{}{}
	rules[ruleToNum(".#...")] = struct{}{}
	rules[ruleToNum(".#.#.")] = struct{}{}
	rules[ruleToNum(".#.##")] = struct{}{}
	rules[ruleToNum(".##..")] = struct{}{}
	rules[ruleToNum(".####")] = struct{}{}
	rules[ruleToNum("#.#.#")] = struct{}{}
	rules[ruleToNum("#.###")] = struct{}{}
	rules[ruleToNum("##.#.")] = struct{}{}
	rules[ruleToNum("##.##")] = struct{}{}
	rules[ruleToNum("###..")] = struct{}{}
	rules[ruleToNum("###.#")] = struct{}{}
	rules[ruleToNum("####.")] = struct{}{}

	makeItSo(state, rules, 20)
}

func makeItSo(state pots, rules map[int]struct{}, iterations int) {
	previousPattern := state.String()
	fmt.Printf("%3d. %v (sum: %d)\n", 0, previousPattern, state.sum())
	for i := 1; i <= iterations; i++ {
		state.next(rules)
		currentPattern := state.String()

		fmt.Printf("%3d. %v (sum: %d) %d\n", i, currentPattern, state.sum(), state.zeroPosition)

		if previousPattern == currentPattern {
			fmt.Println("moving pattern found")
			state.zeroPosition -= iterations - i
			fmt.Printf("sum is: %d\n", state.sum())
			break
		}
		previousPattern = currentPattern
	}
}
