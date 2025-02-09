package sort_test

import (
	"fmt"
	"github.com/johnfercher/go-turbo/internal/core/sort"
	"github.com/johnfercher/go-turbo/internal/testutils"
	"testing"
)

func TestMerge_WhenDescIsFalse_ShouldOrderAsc(t *testing.T) {
	// Arrange
	unsortedArr := testutils.GenerateRandomRanges(5)

	// Act
	sorted := sort.Merge(unsortedArr)

	fmt.Println(sorted)
}

/*
func TestMergeInt_WhenDescIsTrue_ShouldOrderDesc(t *testing.T) {
	// Arrange
	unsortedArr := generate.UnsortedIntArray(10, 100)

	valuesMap := make(map[int]bool)
	for _, value := range unsortedArr {
		valuesMap[value] = true
	}

	// Act
	sortedArr := sort.MergeInt(unsortedArr, true)

	// Assert
	assert.Equal(t, len(unsortedArr), len(sortedArr))
	for i := 0; i < len(sortedArr)-1; i++ {
		assert.LessOrEqual(t, sortedArr[i+1], sortedArr[i])
		assert.True(t, valuesMap[sortedArr[i]])
	}
}*/
