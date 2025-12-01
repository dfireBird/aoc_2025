package main

import (
	_ "embed"
	"fmt"
	"log"
	"strings"

	"github.com/dfirebird/aoc_2025/internal"
)

//go:embed input.txt
var input string

type result int

func main() {
	Result1, err := Part1(input)
	if err == nil {
		fmt.Println("Part 1 Result", Result1)
		Result2, err := Part2(input, Result1)
		if err == nil {
			fmt.Println("Part 2 Result", Result2)
		}
	} else {
		log.Fatalf("Part 1 returned an error %v", err)
	}
}

func Part1(input string) (result, error) {
	rotations := parse(input)

	ans := 0
	curPos := 50
	for _, rotation := range rotations {
		curPos = internal.Mod(curPos+rotation, 100)
		if curPos == 0 {
			ans += 1
		}
	}

	return result(ans), nil
}

func Part2(input string, prevResult result) (result, error) {
	rotations := parse(input)

	ans := 0
	curPos := 50
	for _, rotation := range rotations {
		ans += internal.Abs(rotation / 100)
		rotation = rotation % 100

		newPos := internal.Mod(curPos+rotation, 100)
		if curPos != 0 &&
			(rotation < 0 && newPos > curPos ||
				rotation > 0 && newPos < curPos ||
				newPos == 0) {
			ans += 1
		}

		curPos = newPos
	}

	return result(ans), nil
}

func parse(input string) []int {
	res := make([]int, 0)
	for line := range strings.SplitSeq(strings.TrimSpace(input), "\n") {
		times := internal.ToInt(line[1:])
		if line[0] == 'L' {
			res = append(res, -1*times)
		} else {
			res = append(res, times)
		}
	}

	return res
}
