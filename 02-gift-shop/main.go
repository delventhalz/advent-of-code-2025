// https://adventofcode.com/2025/day/2

package main

import (
	"fmt"
	"os"
	"strings"
	"strconv"
)

func parseRange(rangeStr string) (int, int) {
	rng := strings.Split(rangeStr, "-")
	start, _ := strconv.Atoi(rng[0])
	end, _ := strconv.Atoi(rng[1])
	return start, end
}

func isRepeat(num int) bool {
	asString := strconv.Itoa(num)
	halfLen := len(asString) / 2
	return asString[0: halfLen] == asString[halfLen:]
}

func isRepeatOf(num, count int) bool {
	asString := strconv.Itoa(num)
	if len(asString) % count != 0 {
		return false
	}

	sectionLen := len(asString) / count
	section := asString[0: sectionLen]

	for i := 1; i < count; i++ {
		nextSection := asString[i * sectionLen: (i + 1) * sectionLen]
		if section != nextSection {
			return false
		}
		section = nextSection
	}

	return true
}

func hasAnyRepeat(num int) bool {
	asString := strconv.Itoa(num)

	for i := 2; i <= len(asString); i++ {
		if isRepeatOf(num, i) {
			return true
		}
	}

	return false
}


func partOne(input string) string {
	rangeStrings := strings.Split(input, ",")
	invalidSum := 0

	for _, rangeStr := range rangeStrings {
		start, end := parseRange(rangeStr)

		for i := start; i <= end; i++ {
			if isRepeat(i) {
				invalidSum += i
			}
		}
	}

	return strconv.Itoa(invalidSum)
}

func partTwo(input string) string {
	rangeStrings := strings.Split(input, ",")
	invalidSum := 0

	for _, rangeStr := range rangeStrings {
		start, end := parseRange(rangeStr)

		for i := start; i <= end; i++ {
			if hasAnyRepeat(i) {
				invalidSum += i
			}
		}
	}

	return strconv.Itoa(invalidSum)
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
