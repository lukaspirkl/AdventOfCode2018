package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
)

type requirements map[string]map[string]struct{}

func main() {
	result := ""

	reqs := buildRequirements()

	for {
		ready := findReady(reqs)
		if len(ready) == 0 {
			break
		}
		fmt.Println(ready)
		result += ready[0]
		remove(reqs, ready[0])
	}
	fmt.Println(result)
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
