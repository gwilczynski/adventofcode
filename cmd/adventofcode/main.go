package main

import (
	"fmt"

	"github.com/gwilczynski/adventofcode/pkg/historianhysteria"
	"github.com/gwilczynski/adventofcode/pkg/reader"
)

func main() {
	data, err := reader.ReadData("./data/historianhysteria.txt")
	if err != nil {
		panic(err)
	}

	totalDistance := historianhysteria.TotalDistance(data)
	fmt.Println("total distance:", totalDistance)
}
