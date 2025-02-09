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

func (t *TurboRepository) Get(ctx context.Context, turbo string) (*models.Turbo, error) {
	b, err := os.ReadFile("data/turbo/" + turbo + ".csv")
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

	for _, slice := range arr {
		for _, f := range slice.Flow {
			fmt.Println(f)
		}
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
