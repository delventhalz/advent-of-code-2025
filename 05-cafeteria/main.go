// https://adventofcode.com/2025/day/5

package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func parseFreshRanges(rangesString string) [][2]int64 {
	rangeStrings := strings.Split(rangesString, "\n")
	freshRanges := make([][2]int64, len(rangeStrings))

	for i, rngStr := range rangeStrings {
		prts := strings.Split(rngStr, "-")
		start, _ := strconv.ParseInt(prts[0], 10, 64)
		end, _ := strconv.ParseInt(prts[1], 10, 64)
		freshRanges[i] = [2]int64{start, end}
	}

	return freshRanges
}

func isFresh(freshRanges [][2]int64, id int64) bool {
	for _, rng := range freshRanges {
		if id >= rng[0] && id <= rng[1] {
			return true
		}
	}
	return false
}

func findGteFrom(target int64, numbers []int64, from int) (int, int64) {
	if (from >= len(numbers)) {
		return -1, 0
	}

	for i, num := range numbers[from:] {
		if num >= target {
			return from + i, num
		}
	}

	return -1, 0
}

func combineRanges(ranges [][2]int64) [][2]int64 {
	starts := make([]int64, len(ranges))
	ends := make([]int64, len(ranges))
	for i, rng := range ranges {
		starts[i] = rng[0]
		ends[i] = rng[1]
	}

	slices.Sort(starts)
	slices.Sort(ends)

	s, e := 0, 0
	var start, end int64
	start = starts[s]
	var combined [][2]int64

	for e < len(ends) {
		e, end = findGteFrom(start, ends, e)
		combined = append(combined, [2]int64{start, end})

		nextS, _ := findGteFrom(end + 1, starts, s + 1)
		if (nextS == -1) {
			start = end + 1
		} else if (nextS - s > 1) {
			s = nextS - 1
			start = end + 1
		} else {
			s++
			start = starts[s]
		}

		e++
	}

	return combined
}

func partOne(input string) string {
	parts := strings.Split(input, "\n\n")
	freshRanges := parseFreshRanges(parts[0])
	idStrings := strings.Split(parts[1], "\n")
	freshCount := 0

	for _, idStr := range idStrings {
		id, _ := strconv.ParseInt(idStr, 10, 64)
		if isFresh(freshRanges, id) {
			freshCount++
		}
	}

	return fmt.Sprint(freshCount)
}

func partTwo(input string) string {
	parts := strings.Split(input, "\n\n")
	freshRanges := parseFreshRanges(parts[0])
	combined := combineRanges(freshRanges)

	var rangeSum int64
	for _, rng := range combined {
		rangeSum += rng[1] - rng[0] + 1 // inclusive range
	}

	return fmt.Sprint(rangeSum)
}

func main() {
	inputBytes, err := os.ReadFile("input.txt")

	if err != nil {
		fmt.Println("Error reading input.txt:")
		fmt.Println(err)
		os.Exit(1)
	}

	input := strings.Trim(string(inputBytes), "\n")
	fmt.Printf("Part One Solution: %s\n", partOne(input))
	fmt.Printf("Part Two Solution: %s\n", partTwo(input))
}
