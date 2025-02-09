package csv

import (
	"fmt"
	"reflect"
)

type TurboPressureDAO struct {
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

type TurboPressureDAOArray struct {
	PSI  string
	Flow []string
}

func (t TurboPressureDAO) IsEmpty() bool {
	tType := reflect.ValueOf(t)

	for i := 0; i < tType.NumField(); i++ {
		col := tType.FieldByName(fmt.Sprintf("COL%d", i+1))
		if col.IsValid() {
			s := col.String()
			if s != "" {
				return false
			}
		}
	}

	return true
}

func (t TurboPressureDAO) ToArray() *TurboPressureDAOArray {
	tType := reflect.ValueOf(t)
	arr := &TurboPressureDAOArray{}

	for i := 0; i < tType.NumField(); i++ {
		psi := tType.FieldByName("PSI")
		if psi.IsValid() {
			arr.PSI = psi.String()
		}

		col := tType.FieldByName(fmt.Sprintf("COL%d", i+1))
		if col.IsValid() {
			s := col.String()
			if s != "" {
				arr.Flow = append(arr.Flow, s)
			}
		}
	}

	return arr
}
