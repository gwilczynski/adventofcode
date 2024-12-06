package main

import (
	"fmt"

	"github.com/gwilczynski/adventofcode/pkg/ceressearch"
	"github.com/gwilczynski/adventofcode/pkg/guardgallivant"
	"github.com/gwilczynski/adventofcode/pkg/historianhysteria"
	"github.com/gwilczynski/adventofcode/pkg/mullitover"
	"github.com/gwilczynski/adventofcode/pkg/printqueue"
	"github.com/gwilczynski/adventofcode/pkg/reader"
	"github.com/gwilczynski/adventofcode/pkg/rednosedreport"
)

func main() {
	// historianHysteria()
	// redNosedReport()
	// mullIt0ver()
	// ceresSearch()
	// printQueue()
	guardGallivant()
}

func guardGallivant() {
	fmt.Println("Day 6: Guard Gallivant ---")
	data, err := reader.ReadData("./cmd/data/test.txt")
	if err != nil {
		panic(err)
	}

	r := guardgallivant.Call(data)
	fmt.Println("result: ", r)
}

func printQueue() {
	fmt.Println("--- Day 5: Print Queue ---")
	data, err := reader.ReadData("./cmd/data/print_queue.txt")
	if err != nil {
		panic(err)
	}

	r, _ := printqueue.Call(data)
	fmt.Println("result: ", r)

	_, r = printqueue.Call(data)
	fmt.Println("result (part 2): ", r)
}

func ceresSearch() {
	fmt.Println("--- Day 4: Ceres Search ---")
	data, err := reader.ReadData("./cmd/data/ceres_search.txt")
	if err != nil {
		panic(err)
	}

	r := ceressearch.Call(data, false)
	fmt.Println("result: ", r)

	r = ceressearch.Call(data, true)
	fmt.Println("result (part 2): ", r)
}

func mullIt0ver() {
	fmt.Println("--- Day 3: Mull It Over ---")
	data, err := reader.ReadData("./cmd/data/mull_it_over.txt")
	if err != nil {
		panic(err)
	}

	r := mullitover.Call(data, false)
	fmt.Println("the results of the multiplications: ", r)

	r = mullitover.Call(data, true)
	fmt.Println("the results of the multiplications (combined): ", r)
}

func redNosedReport() {
	fmt.Println("--- Day 2: Red-Nosed Reports ---")
	data, err := reader.ReadData("./cmd/data/red_nosed_reports.txt")
	if err != nil {
		panic(err)
	}

	r := rednosedreport.HowManyReportsAreSafe(data, false)
	fmt.Println("safe reports: ", r)

	rd := rednosedreport.HowManyReportsAreSafe(data, true)
	fmt.Println("safe reports with problem dampener: ", rd)
}

func historianHysteria() {
	fmt.Println("--- Day 1: Historian Hysteria ---")
	data, err := reader.ReadData("./cmd/data/historian_hysteria.txt")
	if err != nil {
		panic(err)
	}

	t := historianhysteria.TotalDistance(data)
	fmt.Println("total distance:", t)

	s := historianhysteria.SimilarityScore(data)
	fmt.Println("similarity score:", s)
}
