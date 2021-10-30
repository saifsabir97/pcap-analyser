package csv

import (
	"encoding/csv"
	"log"
	"os"
)

func CreateCSV(data [][]string, fileName string) {
	csvFile, err := os.Create(fileName)
	if err != nil {
		log.Fatalf("failed creating file: %s", err)
		return
	}
	csvWriter := csv.NewWriter(csvFile)

	for _, vulRow := range data {
		_ = csvWriter.Write(vulRow)
	}
	csvWriter.Flush()
	csvFile.Close()
}
