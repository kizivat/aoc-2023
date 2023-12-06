package main

import (
	"kizivat.eu/aoc2023/common"
)

type Parts struct{}

func (_ Parts) Part1(data string) uint32 {
	return calcCalibrationValues(data, findFirstDigit, findLastDigit)
}

func (_ Parts) Part2(data string) uint32 {
	return calcCalibrationValues(data, findAnyFirstDigit, findAnyLastDigit)
}

func main() {

	common.Main(common.Day[uint32]{
		RunnableParts: Parts{},
		Name:          "01",
	})

}
