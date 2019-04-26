package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type data struct {
	index   int
	x       int
	y       int
	width   int
	height  int
	overlap map[*data]struct{}
}

func createNewData() (func(string) *data, error) {
	regex, err := regexp.Compile(`^#(\d*) @ (\d*),(\d*): (\d*)x(\d*)$`)
	if err != nil {
		return nil, err
	}

	atoi := func(s string) int {
		i, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		return i
	}

	return func(inputRow string) *data {
		s := regex.FindStringSubmatch(inputRow)
		return &data{
			index:   atoi(s[1]),
			x:       atoi(s[2]),
			y:       atoi(s[3]),
			width:   atoi(s[4]),
			height:  atoi(s[5]),
			overlap: make(map[*data]struct{}),
		}
	}, nil
}

func main() {

	fileHandle, _ := os.Open("input.txt")
	defer fileHandle.Close()

	newData, err := createNewData()
	if err != nil {
		panic(err)
	}

	fabric := make(map[int]map[int][]*data)
	claims := make([]*data, 0)

	scanner := bufio.NewScanner(fileHandle)
	for scanner.Scan() {
		d := newData(scanner.Text())
		claims = append(claims, d)
		for xPlus := 0; xPlus < d.width; xPlus++ {
			for yPlus := 0; yPlus < d.height; yPlus++ {
				x, presentX := fabric[d.x+xPlus]
				if !presentX {
					x = make(map[int][]*data)
					fabric[d.x+xPlus] = x
				}
				y, presentY := x[d.y+yPlus]
				if !presentY {
					y = make([]*data, 0)
					x[d.y+yPlus] = y
				}
				for _, overlapped := range y {
					overlapped.overlap[d] = struct{}{}
					d.overlap[overlapped] = struct{}{}
				}
				y = append(y, d)
				x[d.y+yPlus] = y
			}
		}
	}

	for _, claim := range claims {
		if len(claim.overlap) == 0 {
			fmt.Println(claim.index)
			return
		}
	}
}
