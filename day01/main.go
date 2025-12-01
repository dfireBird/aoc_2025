package main

import (
	_ "embed"
	"fmt"
	"log"

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
	return 0, &internal.NotImplemented{}
}

func Part2(input string, prev_result result) (result, error) {
	return 0, &internal.NotImplemented{}
}

