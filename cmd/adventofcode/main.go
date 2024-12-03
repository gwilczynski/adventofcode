package main

import (
	"fmt"

	"github.com/gwilczynski/adventofcode/pkg/historianhysteria"
	"github.com/gwilczynski/adventofcode/pkg/mullitover"
	"github.com/gwilczynski/adventofcode/pkg/reader"
	"github.com/gwilczynski/adventofcode/pkg/rednosedreport"
)

func main() {
	historianHysteria()
	redNosedReport()
	mullIt0ver()
}

func mullIt0ver() {
	fmt.Println("--- Day 3: Mull It Over ---")
	data, err := reader.ReadData("./data/mull_it_over.txt")
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
	data, err := reader.ReadData("./data/red_nosed_reports.txt")
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
	data, err := reader.ReadData("./data/historian_hysteria.txt")
	if err != nil {
		panic(err)
	}

	t := historianhysteria.TotalDistance(data)
	fmt.Println("total distance:", t)

	s := historianhysteria.SimilarityScore(data)
	fmt.Println("similarity score:", s)
}
