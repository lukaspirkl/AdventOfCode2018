package main

import "fmt"

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

func (s *scoreboard) cycle(after int) {
	for len(s.scores) < after+10 {
		s.next()
		//fmt.Println(s)
	}
	for _, value := range s.scores[after : after+10] {
		fmt.Print(value)
	}
	fmt.Println()
}

func main() {
	newScoreboard(3, 7).cycle(440231)
}
