package matrix

import (
	"fmt"
)

func IsSurgeOrChoke(s string) bool {
	return IsSurge(s) || IsChoke(s)
}

func IsSurge(s string) bool {
	return s == "S"
}

func IsChoke(s string) bool {
	return s == "C"
}

func Print[T any](matrix [][]T) {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			fmt.Print(matrix[i][j], " ")
		}
		fmt.Println()
	}
}
