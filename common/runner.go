package common

import (
	"fmt"
	"log"
	"os"

	"golang.org/x/exp/constraints"
)

type Result interface {
	constraints.Ordered | constraints.Integer | constraints.Float
}

type RunnableParts[T Result] interface {
	Part1(data string) T
	Part2(data string) T
}

type Parts struct{}

type Day[T Result] struct {
	Name string
	RunnableParts[T]
}

func Main[T Result](day Day[T]) {
	if len(os.Args) <= 1 {
		fmt.Printf("Provide input file. Ex.: go run ./day%s ./input.txt 1", day.Name)
		return
	}

	if len(os.Args) <= 2 {
		fmt.Printf("Choose either part 1 or 2. Ex.: go run ./day%s ./input.txt 1", day.Name)
		return
	}

	data, err := LoadData(os.Args[1])

	if err != nil {
		log.Fatalf("Could not load data. %v", err)
	}

	switch os.Args[2] {
	case "1":
		fmt.Println(day.Part1(data))
	case "2":
		fmt.Println(day.Part2(data))
	default:
		fmt.Printf("Choose either part 1 or 2. Ex.: go run ./day%s ./input.txt 1", day.Name)
	}
}
