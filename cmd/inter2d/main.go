package main

import (
	"fmt"
	"github.com/edwardbrowncross/naturalneighbour/delaunay"
	"github.com/edwardbrowncross/naturalneighbour/interpolation"
	"math/rand"
)

func main() {
	// Create 1000 random points of source data.
	dataPoints := make([]*delaunay.Point, 1000)
	for i := 0; i < 1000; i++ {
		dataPoints[i] = interpolation.NewPoint(rand.Float64(), rand.Float64(), rand.Float64())
	}
	// Create an interpolator that will use the source data.
	interpolator, err := interpolation.New(dataPoints)
	if err != nil {
		panic(err)
	}

	result, err := interpolator.Interpolate(0.5, 0.5)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}
