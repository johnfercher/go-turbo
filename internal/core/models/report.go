package models

type Report struct {
	Engine  *Engine
	Turbo   *Turbo
	Boost   float64
	Entries Entries
}

type Entry struct {
	RPM      int
	LbsMin   float64
	Health   float64
	MinHP    float64
	MaxHP    float64
	MinHPE85 float64
	MaxHPE85 float64
}

type Entries []Entry

func NewReport(engine *Engine, turbo *Turbo, boost float64) *Report {
	return &Report{
		Engine: engine,
		Turbo:  turbo,
		Boost:  boost,
	}
}

func (r *Report) Add(rpm int, lbsMin float64, health float64) {
	e := Entry{
		RPM:      rpm,
		LbsMin:   lbsMin,
		Health:   health,
		MinHP:    lbsMin * 0.95,
		MaxHP:    lbsMin * 1.05,
		MinHPE85: lbsMin * 0.95 * 1.15,
		MaxHPE85: lbsMin * 1.05 * 1.15,
	}

	r.Entries = append(r.Entries, e)
}

func (e Entries) GetTop() Entry {
	maxLbs := 0.0
	maxIndex := 0

	for i, entry := range e {
		if entry.LbsMin > maxLbs {
			maxLbs = entry.LbsMin
			maxIndex = i
		}
	}

	return e[maxIndex]
}
