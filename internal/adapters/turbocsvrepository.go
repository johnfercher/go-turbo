package adapters

import (
	"bytes"
	"context"
	"fmt"
	"github.com/gocarina/gocsv"
	"os"
)

type TurboPressureCSV struct {
	PSI   string `csv:"psi"`
	COL1  string `csv:"col1"`
	COL2  string `csv:"col2"`
	COL3  string `csv:"col3"`
	COL4  string `csv:"col4"`
	COL5  string `csv:"col5"`
	COL6  string `csv:"col6"`
	COL7  string `csv:"col7"`
	COL8  string `csv:"col8"`
	COL9  string `csv:"col9"`
	COL10 string `csv:"col10"`
	COL11 string `csv:"col11"`
	COL12 string `csv:"col12"`
	COL13 string `csv:"col13"`
	COL14 string `csv:"col14"`
	COL15 string `csv:"col15"`
}

type TurboCsvRepository struct {
}

func NewTurboCsvRepository() *TurboCsvRepository {
	return &TurboCsvRepository{}
}

func (t *TurboCsvRepository) Get(ctx context.Context, turbo string) (string, error) {
	b, err := os.ReadFile("data/turbo/" + turbo + ".csv")
	if err != nil {
		return "", err
	}

	var entries []*TurboPressureCSV
	err = gocsv.Unmarshal(bytes.NewBuffer(b), &entries)
	if err != nil {
		return "", err
	}

	for _, entry := range entries {
		fmt.Println(entry)
	}
	return "", nil
}
