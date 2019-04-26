package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
)

const (
	stepTakes    int = 60
	workersCount int = 5
)

type requirements map[string]map[string]struct{}

func main() {
	workers := make([]int, workersCount)
	inProgress := make([]string, workersCount)

	reqs := buildRequirements()
	for time := 0; ; time++ {
		ready := findReady(reqs)
		if len(ready) == 0 && areWorkersDone(workers) {
			fmt.Printf("TOTAL TIME: %v\n", time)
			break
		}
		ready = readyNow(ready, inProgress)
		for _, r := range ready {
			freeWorker := getFreeWorkerIndex(workers)
			if freeWorker != -1 {
				inProgress[freeWorker] = r
				workers[freeWorker] = getTime(r)
			}
		}

		//fmt.Println(workers)
		//fmt.Println(inProgress)

		subtractSecond(workers, inProgress, reqs)
	}

}

func readyNow(ready []string, inProgress []string) []string {
	readyNow := []string{}
	for _, r := range ready {
		isInProgress := false
		for _, i := range inProgress {
			if r == i {
				isInProgress = true
				break
			}
		}
		if !isInProgress {
			readyNow = append(readyNow, r)
		}
	}
	return readyNow
}

func areWorkersDone(workers []int) bool {
	for _, value := range workers {
		if value != 0 {
			return false
		}
	}
	return true
}

func subtractSecond(workers []int, inProgress []string, reqs requirements) {
	for i := range workers {
		if workers[i] > 0 {
			workers[i]--
			if workers[i] == 0 {
				remove(reqs, inProgress[i])
				inProgress[i] = ""
			}
		}
	}
}

func getFreeWorkerIndex(workers []int) int {
	for index, value := range workers {
		if value <= 0 {
			return index
		}
	}
	return -1
}

func getTime(key string) int {
	return stepTakes + int([]rune(key)[0]-'A') + 1
}

func remove(reqs requirements, key string) {
	req, present := reqs[key]
	if !present {
		return
	}
	delete(reqs, key)
	for dep := range req {
		if _, p := reqs[dep]; !p {
			reqs[dep] = make(map[string]struct{})
		}
	}
}

func findReady(reqs requirements) []string {
	hasDeps := make(map[string]struct{})
	for _, deps := range reqs {
		for dep := range deps {
			hasDeps[dep] = struct{}{}
		}
	}
	ready := []string{}
	for req := range reqs {
		if _, present := hasDeps[req]; !present {
			ready = append(ready, req)
		}
	}
	sort.Strings(ready)
	return ready
}

func buildRequirements() requirements {
	fileHandle, _ := os.Open("input.txt")
	defer fileHandle.Close()

	regex, err := regexp.Compile(`^Step ([A-Z]) must be finished before step ([A-Z]) can begin\.$`)
	if err != nil {
		panic(err)
	}

	requirements := make(requirements)

	scanner := bufio.NewScanner(fileHandle)
	for scanner.Scan() {
		submatch := regex.FindStringSubmatch(scanner.Text())
		r, present := requirements[submatch[1]]
		if !present {
			r = make(map[string]struct{})
			requirements[submatch[1]] = r
		}
		r[submatch[2]] = struct{}{}
	}
	return requirements
}
