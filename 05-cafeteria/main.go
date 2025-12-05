// https://adventofcode.com/2025/day/5

package main

import (
	"fmt"
	"os"
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
	return ""
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
