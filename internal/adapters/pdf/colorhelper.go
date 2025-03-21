package pdf

import "github.com/johnfercher/maroto/v2/pkg/props"

func Color(index int) props.Color {
	switch index {
	case 0:
		return props.Color{
			Red:   0,
			Green: 57,
			Blue:  125,
		}
	case 1:
		return props.Color{
			Red:   217,
			Green: 47,
			Blue:  56,
		}
	case 2:
		return props.Color{
			Red:   76,
			Green: 202,
			Blue:  141,
		}
	case 3:
		return props.Color{
			Red:   199,
			Green: 1,
			Blue:  254,
		}
	case 4:
		return props.Color{
			Red:   253,
			Green: 155,
			Blue:  58,
		}
	case 5:
		return props.Color{
			Red:   0,
			Green: 140,
			Blue:  254,
		}
	default:
		return props.Color{
			Red:   253,
			Green: 115,
			Blue:  179,
		}
	}
}
