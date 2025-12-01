package main

import (
	"testing"

	"github.com/dfirebird/aoc_2025/internal"
)

type TestData internal.TestData[result]

var example = ``

func Test_part1(t *testing.T) {
	tests := []TestData{
		{
			Name:  "example",
			Input: example,
			Want:  11,
		},
	}
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			if got, err := Part1(tt.Input); err != nil || got != tt.Want {
				t.Errorf("part1() = %v, want %v", got, tt.Want)
			}
		})
	}
}

var example2 = ``

func Test_part2(t *testing.T) {
	tests := []TestData{
		{
			Name:  "example",
			Input: example2,
			Want:  31,
		},
	}
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			if got, err := Part2(tt.Input, -1); err != nil || got != tt.Want {
				t.Errorf("part2() = %v, want %v", got, tt.Want)
			}
		})
	}
}

