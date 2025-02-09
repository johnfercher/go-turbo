package models

import "fmt"

func KgKey(boost float64) string {
	return fmt.Sprintf("%.2f", boost)
}
