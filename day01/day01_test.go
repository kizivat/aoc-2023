package main

import (
	"testing"

	"kizivat.eu/aoc2023/common"
)

func TestPart1(t *testing.T) {
	data, err := common.LoadData("./in/test1.txt")
	if err != nil {
		t.Fatalf("Could not load data. %v", err)
	}
	actual := Parts{}.Part1(data)

	var expected uint32 = 142

	if actual != expected {
		t.Fatalf("Result was incorrect, expected: %d, actual: %d.", expected, actual)
	}
}

func TestPart2(t *testing.T) {
	data, err := common.LoadData("./in/test2.txt")
	if err != nil {
		t.Fatalf("Could not load data. %v", err)
	}
	actual := Parts{}.Part2(data)

	var expected uint32 = 281

	if actual != expected {
		t.Fatalf("Result was incorrect, expected: %d, actual: %d.", expected, actual)
	}
}
