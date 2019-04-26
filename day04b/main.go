package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"time"
)

type action int

const (
	wakesUp     action = 0
	fallsAsleep action = 1
	beginsShift action = 2
)

func (action action) String() string {
	names := [...]string{
		"wakes up",
		"falls asleep",
		"begins shift",
	}

	if action < wakesUp || action > beginsShift {
		return "Unknown"
	}

	return names[action]
}

type row struct {
	time   time.Time
	action action
	id     int
}

func (r row) String() string {
	return fmt.Sprintf("%v %v", r.time, r.action)
}

type sleep struct {
	start    time.Time
	duration time.Duration
}

func main() {
	fileHandle, _ := os.Open("input.txt")
	defer fileHandle.Close()

	rows := parseFile(fileHandle)

	sort.Slice(rows, func(i, j int) bool {
		return rows[i].time.Before(rows[j].time)
	})

	guards := parseGuards(rows)

	maxGuardID := 0
	maxMinute := 0
	maxCount := 0

	for guardID, sleeps := range guards {
		minute, count := findMostFrequentSleepingMinute(sleeps)
		if maxCount < count {
			maxGuardID = guardID
			maxMinute = minute
			maxCount = count
		}
	}

	fmt.Println("-----")
	fmt.Println(maxGuardID * maxMinute)

}

func findMostFrequentSleepingMinute(sleeps []sleep) (int, int) {
	minutes := make(map[int]int)
	for _, sleep := range sleeps {
		if sleep.start.Hour() != 0 {
			continue
		}
		for i := 0; i <= int(sleep.duration.Minutes()); i++ {
			count, _ := minutes[sleep.start.Minute()+i]
			minutes[sleep.start.Minute()+i] = count + 1
		}
	}
	maximumMinute := 0
	maximumCount := 0
	for minute, count := range minutes {
		if maximumCount < count {
			maximumCount = count
			maximumMinute = minute
		}
	}
	return maximumMinute, maximumCount
}

func parseGuards(rows []row) map[int][]sleep {
	guards := make(map[int][]sleep)
	currentGuardID := 0
	currentSleep := sleep{}
	for _, row := range rows {
		if row.action == beginsShift {
			currentGuardID = row.id
			_, present := guards[row.id]
			if !present {
				guards[row.id] = make([]sleep, 0)
			}
			continue
		}
		if row.action == fallsAsleep {
			currentSleep = sleep{start: row.time}
			continue
		}
		if row.action == wakesUp {
			currentSleep.duration = row.time.Sub(currentSleep.start.Add(time.Minute))
			guards[currentGuardID] = append(guards[currentGuardID], currentSleep)
			continue
		}
	}
	return guards
}

func parseFile(fileHandle *os.File) []row {
	beginsShiftRegex, _ := regexp.Compile(`^\[([\d- :]*)\].*#(\d*).*$`)
	wakesUpRegex, _ := regexp.Compile(`^\[([\d- :]*)\] wakes up$`)
	fallsAsleepRegex, _ := regexp.Compile(`^\[([\d- :]*)\] falls asleep$`)

	rows := []row{}
	scanner := bufio.NewScanner(fileHandle)
	for scanner.Scan() {
		text := scanner.Text()

		if submatch := beginsShiftRegex.FindStringSubmatch(text); submatch != nil {
			rows = append(rows, row{
				time:   parseTime(submatch[1]),
				action: beginsShift,
				id:     parseInt(submatch[2]),
			})
			continue
		}
		if submatch := wakesUpRegex.FindStringSubmatch(text); submatch != nil {
			rows = append(rows, row{
				time:   parseTime(submatch[1]),
				action: wakesUp,
			})
			continue
		}
		if submatch := fallsAsleepRegex.FindStringSubmatch(text); submatch != nil {
			rows = append(rows, row{
				time:   parseTime(submatch[1]),
				action: fallsAsleep,
			})
			continue
		}
	}
	return rows
}

func parseTime(text string) time.Time {
	layout := "2006-01-02 15:04"
	t, err := time.Parse(layout, text)
	if err != nil {
		panic(err)
	}
	return t
}

func parseInt(text string) int {
	i, _ := strconv.Atoi(text)
	return i
}
