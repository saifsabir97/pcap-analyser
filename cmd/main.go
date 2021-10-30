package main

import (
	"hawk/internal/app"
	"hawk/pkg/csv"
)

func main() {
	analyser, err := app.NewClient("data/test_capture.pcap")
	if err != nil {
		panic("unable to create hawk analyser")
	}
	sessionDetailsMatrix := analyser.Run()
	csv.CreateCSV(sessionDetailsMatrix, "results.csv")
}
