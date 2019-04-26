package main

import (
	"strings"
	"testing"
)

func Test_detailed(t *testing.T) {
	a := newArea(strings.NewReader(strings.TrimSpace(`
#######
#.G...#
#...EG#
#.#.#G#
#..G#E#
#.....#
#######
	`)))
	expected := strings.TrimSpace(`
#######
#G....#   G(200)
#.G...#   G(131)
#.#.#G#   G(59)
#...#.#
#....G#   G(200)
#######
	`)
	assertRound(a.run(nil), 47, t)
	assertSum(a.sumOfHitPoints(allCreatures), 590, t)
	assertArea(a, expected, t)
}

func Test_1(t *testing.T) {
	a := newArea(strings.NewReader(strings.TrimSpace(`
#######
#G..#E#
#E#E.E#
#G.##.#
#...#E#
#...E.#
#######
	`)))
	expected := strings.TrimSpace(`
#######
#...#E#   E(200)
#E#...#   E(197)
#.E##.#   E(185)
#E..#E#   E(200) E(200)
#.....#
#######
	`)
	assertRound(a.run(nil), 37, t)
	assertSum(a.sumOfHitPoints(allCreatures), 982, t)
	assertArea(a, expected, t)
}

func Test_2(t *testing.T) {
	a := newArea(strings.NewReader(strings.TrimSpace(`
#######
#E..EG#
#.#G.E#
#E.##E#
#G..#.#
#..E#.#
#######
	`)))
	expected := strings.TrimSpace(`
#######
#.E.E.#   E(164) E(197)
#.#E..#   E(200)
#E.##.#   E(98)
#.E.#.#   E(200)
#...#.#
#######
	`)
	assertRound(a.run(nil), 46, t)
	assertSum(a.sumOfHitPoints(allCreatures), 859, t)
	assertArea(a, expected, t)
}

func Test_3(t *testing.T) {
	a := newArea(strings.NewReader(strings.TrimSpace(`
#######
#E.G#.#
#.#G..#
#G.#.G#
#G..#.#
#...E.#
#######
	`)))
	expected := strings.TrimSpace(`
#######
#G.G#.#   G(200) G(98)
#.#G..#   G(200)
#..#..#
#...#G#   G(95)
#...G.#   G(200)
#######
	`)
	assertRound(a.run(nil), 35, t)
	assertSum(a.sumOfHitPoints(allCreatures), 793, t)
	assertArea(a, expected, t)
}

func Test_4(t *testing.T) {
	a := newArea(strings.NewReader(strings.TrimSpace(`
#######
#.E...#
#.#..G#
#.###.#
#E#G#G#
#...#G#
#######
	`)))
	expected := strings.TrimSpace(`
#######
#.....#
#.#G..#   G(200)
#.###.#
#.#.#.#
#G.G#G#   G(98) G(38) G(200)
#######
	`)
	assertRound(a.run(nil), 54, t)
	assertSum(a.sumOfHitPoints(allCreatures), 536, t)
	assertArea(a, expected, t)
}

func Test_5(t *testing.T) {
	a := newArea(strings.NewReader(strings.TrimSpace(`
#########
#G......#
#.E.#...#
#..##..G#
#...##..#
#...#...#
#.G...G.#
#.....G.#
#########
	`)))
	expected := strings.TrimSpace(`
#########
#.G.....#   G(137)
#G.G#...#   G(200) G(200)
#.G##...#   G(200)
#...##..#
#.G.#...#   G(200)
#.......#
#.......#
#########
	`)
	assertRound(a.run(nil), 20, t)
	assertSum(a.sumOfHitPoints(allCreatures), 937, t)
	assertArea(a, expected, t)
}

func assertRound(i, expected int, t *testing.T) {
	if i != expected {
		t.Errorf("Round %d does not match with expected %d", i, expected)
	}
}

func assertSum(i, expected int, t *testing.T) {
	if i != expected {
		t.Errorf("Sum %d does not match with expected %d", i, expected)
	}
}

func assertArea(a area, s string, t *testing.T) {
	if strings.Replace(a.String(), " ", "", -1) != strings.Replace(s, " ", "", -1) {
		t.Errorf("%v\ndoes not match with expected:\n%s", a.String(), s)
	}
}
