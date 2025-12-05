// https://adventofcode.com/2025/day/5

package main

import (
	"fmt"
	"os"
	"slices"
	"sort"
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

func combineRanges(ranges [][2]int64) [][2]int64 {
	ordered := slices.Clone(ranges)
	sort.Slice(ordered, func(i, j int) bool {
		return ordered[i][0] < ordered[j][0]
	})

	var combined [][2]int64
	var last [2]int64

	for _, rng := range ordered {
		if combined == nil || rng[0] > last[1] {
			last = rng
			combined = append(combined, rng)
		} else if (rng[1] > last[1]) {
			last = [2]int64{last[1] + 1, rng[1]}
			combined = append(combined, last)
		}
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
