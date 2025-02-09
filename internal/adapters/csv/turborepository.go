package csv

import (
	"bytes"
	"context"
	"github.com/gocarina/gocsv"
	"github.com/johnfercher/go-turbo/internal/core/models"
	"github.com/johnfercher/go-turbo/internal/sort"
	"os"
	"strconv"
	"strings"
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

	var entries []*TurboPressureDAO
	err = gocsv.Unmarshal(bytes.NewBuffer(b), &entries)
	if err != nil {
		return nil, err
	}

	valids := t.filterValids(entries)
	arr := t.toArray(valids)
	slices := t.getSlices(arr)

	return &models.Turbo{
		Name:   turboFile,
		Slices: slices,
	}, nil
}

func (t *TurboRepository) getSlices(arr []*TurboPressureDAOArray) []*models.TurboSlice {
	var slices []*models.TurboSlice
	for _, slice := range arr {
		psi, _ := strconv.ParseFloat(strings.TrimSpace(slice.PSI), 64)
		turboSlice := &models.TurboSlice{
			Boost: psi,
		}

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
		turboSlice.Ranges = append(turboSlice.Ranges, &models.Range{
			Min:   GetFlowFromBaseRange(slice.Flow[base]),
			Max:   GetFlowFromBaseRange(slice.Flow[base+1]),
			Score: baseScore,
		})

		// Add bottom half
		weightIterator := 1
		for i := base; i > 1; i-- {
			turboSlice.Ranges = append(turboSlice.Ranges, &models.Range{
				Min:   GetFlowFromBaseRange(slice.Flow[i-1]),
				Max:   GetFlowFromBaseRange(slice.Flow[i]),
				Score: baseScore + float64(weightIterator),
			})
			weightIterator++
		}

		// Add top half
		weightIterator = 1
		for i := base; i < len(slice.Flow)-2; i++ {
			turboSlice.Ranges = append(turboSlice.Ranges, &models.Range{
				Min:   GetFlowFromBaseRange(slice.Flow[i+1]),
				Max:   GetFlowFromBaseRange(slice.Flow[i+2]),
				Score: baseScore + float64(weightIterator),
			})
			weightIterator++
		}

		turboSlice.Ranges = sort.Merge(turboSlice.Ranges)

		slices = append(slices, turboSlice)
	}

	return slices
}

func (t *TurboRepository) toArray(valids []*TurboPressureDAO) []*TurboPressureDAOArray {
	var arr []*TurboPressureDAOArray
	for _, entry := range valids {
		arr = append(arr, entry.ToArray())
	}
	return arr
}

func (t *TurboRepository) filterValids(entries []*TurboPressureDAO) []*TurboPressureDAO {
	var valids []*TurboPressureDAO
	for _, entry := range entries {
		if !entry.IsEmpty() {
			valids = append(valids, entry)
		}
	}
	return valids
}
