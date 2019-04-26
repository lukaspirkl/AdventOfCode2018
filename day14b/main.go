package main

import (
	"fmt"
	"strings"
)

type scoreboard struct {
	scores []int
	elf1   int
	elf2   int
}

func newScoreboard(scores ...int) *scoreboard {
	if len(scores) < 2 {
		panic("at least two scores should be defined")
	}
	return &scoreboard{
		scores: scores,
		elf1:   0,
		elf2:   1,
	}
}

func (s *scoreboard) next() {
	nextPosition := func(position int) int {
		return (position + s.scores[position] + 1) % len(s.scores)
	}

	nextScore := s.scores[s.elf1] + s.scores[s.elf2]
	tens := (nextScore / 10) % 10
	if tens > 0 {
		s.scores = append(s.scores, tens)
	}
	s.scores = append(s.scores, nextScore%10)

	s.elf1 = nextPosition(s.elf1)
	s.elf2 = nextPosition(s.elf2)
}

func (s *scoreboard) String() string {
	str := ""
	for i, score := range s.scores {
		if s.elf1 == i {
			str = fmt.Sprintf("%s (%d)", str, score)
		} else if s.elf2 == i {
			str = fmt.Sprintf("%s [%d]", str, score)
		} else {
			str = fmt.Sprintf("%s  %d ", str, score)
		}
	}
	return str
}

func (s *scoreboard) indexOf(pattern string) int {
	fmt.Println("start making string")
	str := make([]rune, len(s.scores))
	for i, score := range s.scores {
		str[i] = ([]rune(fmt.Sprintf("%d", score)))[0]
	}
	fmt.Println("start search")
	index := strings.Index(string(str), pattern)
	if index == -1 {
		return -1
	}
	return index
}

func (s *scoreboard) cycle(pattern string) {
	for i := 0; i < 21000000; i++ {
		s.next()
		//fmt.Println(len(s.scores))
	}
	fmt.Println("-----")
	fmt.Println(s.indexOf(pattern))
}

func main() {
	newScoreboard(3, 7).cycle("440231")
}
