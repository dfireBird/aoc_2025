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

type Dir string

const (
	LEFT  Dir = "L"
	RIGHT Dir = "R"
)

type Rotation struct {
	dir   Dir
	times int
}

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

	numZero := 0
	curPos := 50
	for _, rotation := range rotations {
		curPos = rotateWhole(curPos, rotation)
		if curPos == 0 {
			numZero += 1
		}
	}

	return result(numZero), nil
}

func Part2(input string, prev_result result) (result, error) {
	rotations := parse(input)

	numZero := 0
	curPos := 50
	for _, rotation := range rotations {
		newPos, zeroClicks := rotateSingle(curPos, rotation)
		numZero += zeroClicks
		curPos = newPos
	}

	return result(numZero), nil
}

func rotateWhole(curPos int, rot Rotation) int {
	if rot.dir == RIGHT {
		return (curPos + rot.times) % 100
	} else {
		newPos := (curPos - rot.times) % 100
		if newPos < 0 {
			newPos = 100 + newPos
		}

		return newPos
	}
}

func rotateSingle(curPos int, rot Rotation) (int, int) {
	clickZero := 0
	for times := rot.times; times > 0; times -- {
		if rot.dir == RIGHT {
			curPos = (curPos + 1) % 100
		} else {
			curPos = curPos - 1
			if curPos < 0 {
				curPos = 100 + curPos
			}
		}

		if curPos == 0 {
			clickZero ++
		}
	}

	return curPos, clickZero
}

func parse(input string) []Rotation {
	res := make([]Rotation, 0)
	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}
		times := internal.ToInt(line[1:])
		if line[0] == 'L' {
			res = append(res, Rotation{LEFT, times})
		} else {
			res = append(res, Rotation{RIGHT, times})
		}
	}

	return res
}
