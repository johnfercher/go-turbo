package csv

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"github.com/johnfercher/go-turbo/internal/core/models"
)

func Load(b []byte) [][]string {
	reader := csv.NewReader(bytes.NewBuffer(b))
	reader.FieldsPerRecord = -1
	csvData, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	data := addSurgeFirstColumnSurge(csvData)
	data = addChokeColumns(data)
	data = addWeightsOutsideSurgeChoke(data)

	return data
}

func addSurgeFirstColumnSurge(data [][]string) [][]string {
	var minPadded [][]string
	for i := 1; i < len(data); i++ {
		var line []string
		line = append(line, "S")
		for j := 1; j < len(data[i]); j++ {
			line = append(line, data[i][j])
		}
		minPadded = append(minPadded, line)
	}
	return minPadded
}

func addChokeColumns(data [][]string) [][]string {
	for i := 0; i < len(data); i++ {
		allEmpty := true
		for j := 0; j < len(data[i]); j++ {
			if data[i][j] != "" && data[i][j] != "S" {
				allEmpty = false
				break
			}
		}

		if allEmpty {
			for j := 0; j < len(data[i]); j++ {
				d := data[i][j]
				if d == "" {
					data[i][j] = "S"
				}
			}
		} else {
			for j := 0; j < len(data[i]); j++ {
				d := data[i][j]
				if d == "" {
					data[i][j] = "C"
				}
			}
		}

	}

	return data
}

func addWeightsOutsideSurgeChoke(data [][]string) [][]string {
	for i := 0; i < len(data); i++ {
		found := 0
		max := 0
		base := 0.0
		for j := 0; j < len(data[i]); j++ {
			if models.IsBaseRange(data[i][j]) {
				base = models.GetScoreFromBaseRange(data[i][j])
				break
			} else {
				found++
			}
		}
		max = found
		for j := 0; j < max; j++ {
			if data[i][j] != "S" && data[i][j] != "C" {
				data[i][j] = fmt.Sprintf("%s-%d", data[i][j], found+int(base))
			}
			found--
		}
		offset := 1
		for j := max + 2; j < len(data[i]); j++ {
			if data[i][j] != "S" && data[i][j] != "C" {
				data[i][j] = fmt.Sprintf("%s-%d", data[i][j], offset+int(base))
			}
			offset++
		}
	}

	return data
}
