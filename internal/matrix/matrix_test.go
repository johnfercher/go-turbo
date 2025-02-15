package matrix_test

import (
	"github.com/johnfercher/go-turbo/internal/adapters/csv"
	"github.com/johnfercher/go-turbo/internal/matrix"
	"testing"
)

var turboCsv = `
kg,col1,col2,col3,col4,col5,col6,col7,col8,col9,col10,col11,col12,col13,col14,col15,col16
0.0,,,,,,,,,,,,,,,,
0.2,,,,,,,,,,,,,,,,
0.4,,,1,3,4,5-2,5-2,7,8,10,,,,,,
0.6,,,1,3,4,5-1,5-1,7,8,10,,,,,,
0.8,,,1,3,4,5-1,5-1,7,8,10,,,,,,
1.0,,,1,3,4,5-1,5-1,7,8,10,,,,,,
1.2,,,1,3,4,5-1,5-1,7,8,10,,,,,,
1.4,,,,,,,,,,,,,,,,
1.6,,,,,,,,,,,,,,,,
1.8,,,,,,,,,,,,,,,,
2.0,,,,,,,,,,,,,,,,
`

func TestInitMatrix(t *testing.T) {
	t.Run("", func(t *testing.T) {
		turbo := csv.Load([]byte(turboCsv))

		matrix.Print(turbo)

		m := matrix.InitMatrix(10, 10)

		matrix.PrintCFM(m)

		m = matrix.Val(m, turbo)

		matrix.PrintCFM(m)
	})
}
