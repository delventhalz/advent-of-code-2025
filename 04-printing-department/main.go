// https://adventofcode.com/2025/day/4

package main

import (
	"fmt"
	"os"
	"strings"
)

func countAdjacentRolls(diagram []string, x, y int) int {
	xMax := len(diagram[y]) - 1
	yMax := len(diagram) - 1
	count := 0

	if x > 0 && y > 0 && diagram[y - 1][x - 1] == '@' {
		count++
	}
	if y > 0 && diagram[y - 1][x] == '@' {
		count++
	}
	if x < xMax && y > 0 && diagram[y - 1][x + 1] == '@' {
		count++
	}
	if x < xMax && diagram[y][x + 1] == '@' {
		count++
	}
	if x < xMax && y < yMax && diagram[y + 1][x + 1] == '@' {
		count++
	}
	if y < yMax && diagram[y + 1][x] == '@' {
		count++
	}
	if x > 0 && y < yMax && diagram[y + 1][x - 1] == '@' {
		count++
	}
	if x > 0 && diagram[y][x - 1] == '@' {
		count++
	}

	return count
}

func removeRolls(diagram []string) ([]string, int) {
	updatedDiagram := make([]string, len(diagram))
	removeCount := 0

	for y := 0; y < len(diagram); y++ {
		updatedDiagram[y] = ""

		for x := 0; x < len(diagram[y]); x++ {
			if diagram[y][x] == '@' && countAdjacentRolls(diagram, x, y) < 4 {
				removeCount++
				updatedDiagram[y] += "."
			} else {
				updatedDiagram[y] += string(diagram[y][x])
			}
		}
	}

	return updatedDiagram, removeCount
}

func partOne(input string) string {
	diagram := strings.Split(input, "\n")
	accessibleCount := 0

	for y := 0; y < len(diagram); y++ {
		for x := 0; x < len(diagram[y]); x++ {
			if diagram[y][x] == '@' {
				adjacent := countAdjacentRolls(diagram, x, y)
				if adjacent < 4 {
					accessibleCount++
				}
			}
		}
	}

	return fmt.Sprint(accessibleCount)
}

func partTwo(input string) string {
	diagram := strings.Split(input, "\n")
	removeCount := 0
	lastRemoveCount := -1

	for lastRemoveCount != 0 {
		nextDiagram, nextRemoveCount := removeRolls(diagram)
		diagram = nextDiagram
		removeCount += nextRemoveCount
		lastRemoveCount = nextRemoveCount
	}

	return fmt.Sprint(removeCount)
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
