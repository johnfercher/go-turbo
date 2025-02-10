package csv

import (
	"bytes"
	"context"
	"encoding/csv"
	"fmt"
	"github.com/johnfercher/go-turbo/internal/core/models"
	"os"
)

type TurboRepository struct {
}

func NewTurboRepository() *TurboRepository {
	return &TurboRepository{}
}

func (t *TurboRepository) Get(ctx context.Context, turboFile string) (*models.Turbo, error) {
	b, err := os.ReadFile("data/turbo/" + turboFile + ".csv")
	if err != nil {
		return nil, err
	}

	reader := csv.NewReader(bytes.NewBuffer(b))
	reader.FieldsPerRecord = -1
	csvData, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	data := t.getDataWithMinPadding(csvData)
	data = t.getDataWithMaxPadding(data)

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
			data[i][j] = fmt.Sprintf("%s-%d", data[i][j], found+int(base))
			found--
		}
		offset := 1
		for j := max + 2; j < len(data[i]); j++ {
			data[i][j] = fmt.Sprintf("%s-%d", data[i][j], offset+int(base))
			offset++
		}
	}

	fmt.Println(data)

	/*matrix := matrix.BuildEmptyTurbo()
	arr := t.toArray(entries)

	for i := 0; i < consts.TurboMaxLines; i++ {
		for j := 0; j < consts.TurboMaxColumns; j++ {
			matrix[i][j] = models.NewTurboScoreFromString(arr[i].Flow[j])
		}
	}

	return &models.Turbo{
		Name:   turboFile,
		Slices: nil,
	}, nil*/
	return nil, nil
}

func (t *TurboRepository) getDataWithMaxPadding(data [][]string) [][]string {
	for i := 0; i < len(data); i++ {
		for j := 0; j < len(data[i]); j++ {
			d := data[i][j]
			if d == "" {
				data[i][j] = "C"
			}
		}
	}

	return data
}

func (t *TurboRepository) getDataWithMinPadding(data [][]string) [][]string {
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

/*func (t *TurboRepository) getSlices(arr []*TurboPressureDAOArray) map[string][]*models.Range {
	slices := make(map[string][]*models.Range)
	for _, slice := range arr {
		var ranges []*models.Range

		// find base line, the better range
		base := 0
		for i, f := range slice.Flow {
			if IsBaseRange(f) {
				base = i
				break
			}
		}

		baseScore := GetScoreFromBaseRange(slice.Flow[base])

		// Add base line, the better range
		ranges = append(ranges, &models.Range{
			Min:    GetFlowFromBaseRange(slice.Flow[base]),
			Max:    GetFlowFromBaseRange(slice.Flow[base+1]),
			Health: baseScore,
			Boost:  baseScore,
		})

		// Add bottom half
		for i := base; i > 0; i-- {
			ranges = append(ranges, &models.Range{
				Min: GetFlowFromBaseRange(slice.Flow[i-1]),
				Max: GetFlowFromBaseRange(slice.Flow[i]),
			})
		}

		// Surge bottom
		ranges = append(ranges, &models.Range{
			Min: 0,
			Max: GetFlowFromBaseRange(slice.Flow[0]),
		})

		// Add top half
		for i := base; i < len(slice.Flow)-2; i++ {
			ranges = append(ranges, &models.Range{
				Min: GetFlowFromBaseRange(slice.Flow[i+1]),
				Max: GetFlowFromBaseRange(slice.Flow[i+2]),
			})
		}

		// Choke top
		ranges = append(ranges, &models.Range{
			Min: GetFlowFromBaseRange(slice.Flow[len(slice.Flow)-1]),
			Max: 10000,
		})

		ranges = sort.Merge(ranges)

		kg, _ := strconv.ParseFloat(strings.TrimSpace(slice.Kg), 64)
		slices[models.KgKey(kg)] = ranges
	}

	return slices
}*/

func (t *TurboRepository) toArray(valids []*TurboPressureDAO) []*TurboPressureDAOArray {
	var arr []*TurboPressureDAOArray
	for _, entry := range valids {
		arr = append(arr, entry.ToArray())
	}
	return arr
}
