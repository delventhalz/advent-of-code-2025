// https://adventofcode.com/2025/day/7

package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

func appendUnique[E comparable](slc []E, val E) []E {
	if slices.Contains(slc, val) {
		return slc
	}
	return append(slc, val)
}

func sum(numbers []int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

func partOne(input string) string {
	lines := strings.Split(input, "\n")
	beams := []int{strings.Index(lines[0], "S")}
	splitCount := 0

	for _, line := range lines {
		var nextBeams []int

		for _, beam := range beams {
			if line[beam] == '^' {
				splitCount++
				nextBeams = appendUnique(nextBeams, beam - 1)
				nextBeams = appendUnique(nextBeams, beam + 1)
			} else {
				nextBeams = appendUnique(nextBeams, beam)
			}
		}

		beams = nextBeams
	}

	return fmt.Sprint(splitCount)
}

func partTwo(input string) string {
	lines := strings.Split(input, "\n")
	beamsAt := make([]int, len(lines[0]))
	beamsAt[strings.Index(lines[0], "S")] = 1

	for _, line := range lines {
		nextBeamsAt := make([]int, len(beamsAt))

		for col, count := range beamsAt {
			if count > 0 {
				if line[col] == '^' {
					nextBeamsAt[col - 1] += count
					nextBeamsAt[col + 1] += count
				} else {
					nextBeamsAt[col] += count
				}
			}
		}

		beamsAt = nextBeamsAt
	}

	return fmt.Sprint(sum(beamsAt))
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
