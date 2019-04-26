package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"unicode"
)

const maxInt = int(^uint(0) >> 1)

func main() {
	min := maxInt
	for r := 'a'; r <= 'z'; r++ {
		data := readFile()
		data = strings.Replace(data, string(r), "", -1)
		data = strings.Replace(data, string(unicode.ToUpper(r)), "", -1)
		length := react(data)
		if min > length {
			min = length
		}
	}
	fmt.Println(min)
}

func react(data string) int {
	count := 1
	for count > 0 {
		data, count = removeMatch(data)
	}
	return len(data)
}

func removeMatch(data string) (string, int) {
	runes := []rune(data)
	newRunes := []rune{}
	counter := 0
	for i := 0; i < len(runes)-1; i++ {
		if isMatch(runes[i], runes[i+1]) {
			i++
			counter++
		} else {
			newRunes = append(newRunes, runes[i])
		}
		if i == len(runes)-2 {
			newRunes = append(newRunes, runes[i+1])
		}
	}
	return string(newRunes), counter
}

func isMatch(a, b rune) bool {
	if unicode.ToLower(a) == unicode.ToLower(b) {
		if (unicode.IsLower(a) && unicode.IsUpper(b)) || (unicode.IsUpper(a) && unicode.IsLower(b)) {
			return true
		}
	}
	return false
}

func readFile() string {
	b, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	return string(b)
}
