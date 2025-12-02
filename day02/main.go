package main

import (
	_ "embed"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/dfirebird/aoc_2025/internal"
)

//go:embed input.txt
var input string

type result int

type Range struct {
	start int
	end   int
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
	return runSolution(input, getInvalidIdsP1)
}

func Part2(input string, prev_result result) (result, error) {
	return runSolution(input, getInvalidIdsP2)
}

func runSolution(input string, invalidFn func(Range) []int) (result, error) {
    idRanges := parse(input)
	invalidIds := make([]int, 0)
	for _, idRange := range idRanges {
		ids := invalidFn(idRange)
		invalidIds = append(invalidIds, ids...)
	}

	ans := 0
	for _, id := range invalidIds {
		ans += id
	}

	return result(ans), nil
}

func getInvalidIdsP1(idRange Range) []int {
	invalidIds := make([]int, 0)
	for id := idRange.start; id <= idRange.end; id++ {
		idStr := strconv.Itoa(id)
		midIdx := len(idStr) / 2
		if idStr[:midIdx] == idStr[midIdx:] {
			invalidIds = append(invalidIds, id)
		}
	}
	return invalidIds
}

func getInvalidIdsP2(idRange Range) []int{
    invalidIds := make([]int, 0)
	for id := idRange.start; id <= idRange.end; id++ {
		idStr := strconv.Itoa(id)
		for idx := len(idStr); idx > 0; idx -- {
			subStr := idStr[:idx]
			count := strings.Count(idStr, subStr)
			if count >= 2 && len(idStr) == count * len(subStr) {
				invalidIds = append(invalidIds, id)
				break
			}
		}
	}
	return invalidIds
}

func parse(input string) []Range {
	res := make([]Range, 0)
	for val := range strings.SplitSeq(strings.TrimSpace(input), ",") {
		splitVal := strings.Split(val, "-")
		res = append(res, Range{start: internal.ToInt(splitVal[0]), end: internal.ToInt(splitVal[1])})
	}
	return res
}
