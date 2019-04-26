package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
)

type opcode func(registers, instruction []int) []int

func rr(registers, instruction []int, op func(int, int) int) []int {
	new := make([]int, len(registers))
	copy(new, registers)
	new[instruction[3]] = op(registers[instruction[1]], registers[instruction[2]])
	return new
}

func rv(registers, instruction []int, op func(int, int) int) []int {
	new := make([]int, len(registers))
	copy(new, registers)
	new[instruction[3]] = op(registers[instruction[1]], instruction[2])
	return new
}

func vr(registers, instruction []int, op func(int, int) int) []int {
	new := make([]int, len(registers))
	copy(new, registers)
	new[instruction[3]] = op(instruction[1], registers[instruction[2]])
	return new
}

func addr(r, i []int) []int {
	return rr(r, i, func(a, b int) int {
		return a + b
	})
}

func addi(r, i []int) []int {
	return rv(r, i, func(a, b int) int {
		return a + b
	})
}

func mulr(r, i []int) []int {
	return rr(r, i, func(a, b int) int {
		return a * b
	})
}

func muli(r, i []int) []int {
	return rv(r, i, func(a, b int) int {
		return a * b
	})
}

func banr(r, i []int) []int {
	return rr(r, i, func(a, b int) int {
		return a & b
	})
}

func bani(r, i []int) []int {
	return rv(r, i, func(a, b int) int {
		return a & b
	})
}

func borr(r, i []int) []int {
	return rr(r, i, func(a, b int) int {
		return a | b
	})
}

func bori(r, i []int) []int {
	return rv(r, i, func(a, b int) int {
		return a | b
	})
}

func setr(r, i []int) []int {
	return rr(r, i, func(a, b int) int {
		return a
	})
}

func seti(r, i []int) []int {
	return vr(r, i, func(a, b int) int {
		return a
	})
}

func gtir(r, i []int) []int {
	return vr(r, i, func(a, b int) int {
		if a > b {
			return 1
		}
		return 0
	})
}

func gtri(r, i []int) []int {
	return rv(r, i, func(a, b int) int {
		if a > b {
			return 1
		}
		return 0
	})
}

func gtrr(r, i []int) []int {
	return rr(r, i, func(a, b int) int {
		if a > b {
			return 1
		}
		return 0
	})
}

func eqir(r, i []int) []int {
	return vr(r, i, func(a, b int) int {
		if a == b {
			return 1
		}
		return 0
	})
}

func eqri(r, i []int) []int {
	return rv(r, i, func(a, b int) int {
		if a == b {
			return 1
		}
		return 0
	})
}

func eqrr(r, i []int) []int {
	return rr(r, i, func(a, b int) int {
		if a == b {
			return 1
		}
		return 0
	})
}

func count(opcodes []opcode, before, instruction, after []int) int {
	count := 0
	for _, op := range opcodes {
		result := op(before, instruction)
		match := true
		for i := range after {
			if result[i] != after[i] {
				match = false
				break
			}
		}
		if match {
			count++
		}
	}
	return count
}

func main() {
	opcodes := []opcode{addr, addi, mulr, muli, banr, bani, borr, bori, setr, seti, gtir, gtri, gtrr, eqir, eqri, eqrr}

	rexp, err := regexp.Compile(`Before: \[(\d*), (\d*), (\d*), (\d*)\]\r?\n(\d*) (\d*) (\d*) (\d*)\r?\nAfter: *\[(\d*), (\d*), (\d*), (\d*)\]`)
	if err != nil {
		panic(err)
	}

	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	result := 0
	for _, m := range rexp.FindAllStringSubmatch(string(input), -1) {
		before := []int{atoi(m[1]), atoi(m[2]), atoi(m[3]), atoi(m[4])}
		instruction := []int{atoi(m[5]), atoi(m[6]), atoi(m[7]), atoi(m[8])}
		after := []int{atoi(m[9]), atoi(m[10]), atoi(m[11]), atoi(m[12])}
		c := count(opcodes, before, instruction, after)
		if c >= 3 {
			result++
		}
	}
	fmt.Println(result)
}

func atoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}
