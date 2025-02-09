package csv

import (
	"bytes"
	"context"
	"fmt"
	"github.com/gocarina/gocsv"
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

	var entries []*TurboPressureDAO
	err = gocsv.Unmarshal(bytes.NewBuffer(b), &entries)
	if err != nil {
		return nil, err
	}

	valids := t.filterValids(entries)
	arr := t.toArray(valids)

	turbo := &models.Turbo{}
	for _, slice := range arr {
		turboSlice := &models.TurboSlice{}
		base := 0
		for i, f := range slice.Flow {
			if IsBaseRange(f) {
				base = i
				break
			}
		}

		turboSlice.Ranges = append(turboSlice.Ranges, &models.Range{
			Min:   GetFlowFromBaseRange(slice.Flow[base]),
			Max:   GetFlowFromBaseRange(slice.Flow[base+1]),
			Score: GetScoreFromBaseRange(slice.Flow[base]),
		})

		turbo.Slices = append(turbo.Slices, turboSlice)
	}

	for _, slice := range turbo.Slices {
		fmt.Println(slice)
	}

	return nil, nil
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
