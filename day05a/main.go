package main

import (
	"fmt"
	"io/ioutil"
	"unicode"
)

func main() {
	data := readFile()
	count := 1
	for count > 0 {
		data, count = remove(data)
		fmt.Println(count)
	}
	// fmt.Println("---------")
	// fmt.Println(data)
	fmt.Println("---------")
	fmt.Println(len(data))
}

func remove(data string) (string, int) {
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
