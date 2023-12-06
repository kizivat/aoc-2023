package main

import (
	"kizivat.eu/aoc2023/common"
)

type Parts struct{}

func (_ Parts) Part1(data string) uint32 {
	panic("not implemented")
}

func (_ Parts) Part2(data string) uint32 {
	panic("not implemented")
}

func main() {

	common.Main(common.Day[uint32]{
		RunnableParts: Parts{},
		Name:          "01",
	})

}
