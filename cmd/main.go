package main

import (
	"github.com/saifsabir97/pcap-analyser/internal/app"
	"github.com/saifsabir97/pcap-analyser/pkg/csv"
)

func main() {
	analyser, err := app.NewClient("data/test_capture.pcap")
	if err != nil {
		panic("unable to create analyser client")
	}
	sessionDetailsMatrix := analyser.Run()
	csv.CreateCSV(sessionDetailsMatrix, "data/results.csv")
}
