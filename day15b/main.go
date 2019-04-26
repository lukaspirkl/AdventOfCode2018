package main

import (
	"fmt"
	"os"
)

func main() {
	fileHandle, _ := os.Open("input.txt")
	defer fileHandle.Close()

	var a area
	win := false
	round := 0
	for !win {
		a = newArea(fileHandle)
		round, win = a.run(nil)
		elfPower++
	}

	sum := a.sumOfHitPoints(allCreatures)
	fmt.Printf("Ended in round %d with total %d HP -> outcome: %d", round, sum, round*sum)
}

func print(r int, a area) {
	fmt.Println(r)
	fmt.Println(a)
}
