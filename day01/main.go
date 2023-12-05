package main

import (
	"fmt"
	"log"
	"os"

	"kizivat.eu/aoc2023/common"
)

func Part1(data string) uint32 {
	return calcCalibrationValues(data, findFirstDigit, findLastDigit)
}

func Part2(data string) uint32 {
	return calcCalibrationValues(data, findAnyFirstDigit, findAnyLastDigit)
}

func main() {

	if len(os.Args) <= 1 {
		fmt.Println("Provide input file. Ex.: go run ./day01 ./input.txt 1")
		return
	}

	if len(os.Args) <= 2 {
		fmt.Println("Choose either part 1 or 2. Ex.: go run ./day01 ./input.txt 1")
		return
	}

	data, err := common.LoadData(os.Args[1])

	if err != nil {
		log.Fatalf("Could not load data. %v", err)
	}

	switch os.Args[2] {
	case "1":
		fmt.Println(Part1(data))
	case "2":
		fmt.Println(Part2(data))
	default:
		fmt.Println("Choose either part 1 or 2. Ex.: go run ./day01 ./input.txt 1")
	}

}
