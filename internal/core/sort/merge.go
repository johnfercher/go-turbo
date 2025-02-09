package sort

type Sortable[T any] interface {
	GreaterThan(a T) bool
}

func Merge[T Sortable[T]](arr []*T) []*T {
	if len(arr) < 2 {
		return arr
	}

	left := arr[:len(arr)/2]
	right := arr[len(arr)/2:]

	sortedLeft := Merge[T](left)
	sortedRight := Merge[T](right)

	return mergeArrays[T](sortedLeft, sortedRight)
}

func mergeArrays[T Sortable[T]](a []*T, b []*T) []*T {
	var merged []*T

	i := 0
	j := 0
	for i < len(a) && j < len(b) {
		aP := *a[i]
		bP := *b[j]
		if !aP.GreaterThan(bP) {
			merged = append(merged, a[i])
			i++
		} else {
			merged = append(merged, b[j])
			j++
		}
	}

	for ; i < len(a); i++ {
		merged = append(merged, a[i])
	}

	for ; j < len(b); j++ {
		merged = append(merged, b[j])
	}

	return merged
}
