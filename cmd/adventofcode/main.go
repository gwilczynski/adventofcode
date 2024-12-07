package main

import (
	"fmt"

	"github.com/gwilczynski/adventofcode/pkg/bridgerepair"
	"github.com/gwilczynski/adventofcode/pkg/reader"
)

func main() {
	bridgeRepair()
}

func bridgeRepair() {
	fmt.Println("--- Day 7: Bridge Repair ---")

	data, err := reader.ReadData("./cmd/data/bridge_repair.txt")
	if err != nil {
		panic(err)
	}

	r := bridgerepair.Call(data)
	fmt.Println("result: ", r)
}
