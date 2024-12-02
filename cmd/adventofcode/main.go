package main

import (
	"fmt"

	"github.com/gwilczynski/adventofcode/pkg/rednosedreport"

	"github.com/gwilczynski/adventofcode/pkg/historianhysteria"
	"github.com/gwilczynski/adventofcode/pkg/reader"
)

func main() {
	historianHysteria()
	redNosedReport()
}

func redNosedReport() {
	fmt.Println("--- Day 2: Red-Nosed Reports ---")
	data, err := reader.ReadData("./data/red_nosed_reports.txt")
	if err != nil {
		panic(err)
	}

	safeReports := rednosedreport.HowManyReportsAreSafe(data, false)
	fmt.Println("safe reports: ", safeReports)

	safeReportsWithProblemDampener := rednosedreport.HowManyReportsAreSafe(data, true)
	fmt.Println("safe reports with problem dampener: ", safeReportsWithProblemDampener)
}

func historianHysteria() {
	fmt.Println("--- Day 1: Historian Hysteria ---")
	data, err := reader.ReadData("./data/historian_hysteria.txt")
	if err != nil {
		panic(err)
	}

	totalDistance := historianhysteria.TotalDistance(data)
	fmt.Println("total distance:", totalDistance)

	similarityScore := historianhysteria.SimilarityScore(data)
	fmt.Println("similarity score:", similarityScore)
}
