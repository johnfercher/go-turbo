package models

type Range struct {
	Min   float64
	Max   float64
	Score int
}

type TurboSlice struct {
	PSI    string
	Ranges []Range
}
