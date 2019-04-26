package main

import (
	"fmt"
	"os"
)

func main() {
	fileHandle, _ := os.Open("input.txt")
	defer fileHandle.Close()

	a := newArea(fileHandle)

	//fmt.Println(a)
	// round := a.run(func(r int, a area) {
	// 	fmt.Println(r)
	// 	fmt.Println(a)
	// })

	round := a.run(nil)
	sum := a.sumOfHitPoints(allCreatures)
	fmt.Printf("Ended in round %d with total %d HP -> outcome: %d", round, sum, round*sum)
}
