package main

import (
	"bufio"
	"fmt"
	"gocv.io/x/gocv"
	"image"
	"image/color"
	"os"
	"strconv"
	"strings"
	"time"
)

const ESC = 27
const Click = 1

var stateMachine = GetStateMachine()

var borderA = Point{}
var borderB = Point{}
var borderCount = 0

var xPoints []Point
var yPoints []Point

var exit = false

func main() {
	go CLI()

	window := gocv.NewWindow("GOCV - gopher")
	img := gocv.IMRead("assets/images/kinugawa18g.png", gocv.IMReadAnyColor)

	window.ResizeWindow(img.Cols(), img.Rows())
	window.SetMouseHandler(MouseHandler, nil)

	for !exit {
		if stateMachine.GetType() != GetBorders {
			gocv.Rectangle(&img, image.Rect(borderA.X, borderA.Y, borderB.X, borderB.Y), color.RGBA{
				R: 255,
			}, 1)
		}

		if stateMachine.GetType() != GetBorders && stateMachine.GetType() != GetScales {
			for _, point := range xPoints {
				gocv.Circle(&img, image.Point{
					X: point.X,
					Y: point.Y,
				}, 1, color.RGBA{
					B: 255,
				}, 1)
			}
			for _, point := range yPoints {
				gocv.Circle(&img, image.Point{
					X: point.X,
					Y: point.Y,
				}, 1, color.RGBA{
					B: 255,
				}, 1)
			}
		}

		window.IMShow(img)
		key := window.WaitKey(1)
		if key == ESC {
			exit = true
		}
	}
}

func CLI() {
	var minFlow int
	var maxFlow int
	var flowChunks int
	var minPressure int
	var maxPressure int
	var pressureChunks int

	for !exit {
		if stateMachine.GetType() != GetBorders {
			reader := bufio.NewReader(os.Stdin)
			fmt.Println("Min Flow:")
			minFlowString, _ := reader.ReadString('\n')
			minFlow, _ = strconv.Atoi(strings.TrimSpace(minFlowString))
			fmt.Println("Max Flow:")
			maxFlowString, _ := reader.ReadString('\n')
			maxFlow, _ = strconv.Atoi(strings.TrimSpace(maxFlowString))
			fmt.Println("Flow Chunks:")
			flowChunksString, _ := reader.ReadString('\n')
			flowChunks, _ = strconv.Atoi(strings.TrimSpace(flowChunksString))
			fmt.Println("Min Pressure:")
			minPressureString, _ := reader.ReadString('\n')
			minPressure, _ = strconv.Atoi(strings.TrimSpace(minPressureString))
			fmt.Println("Max Pressure:")
			maxPressureString, _ := reader.ReadString('\n')
			maxPressure, _ = strconv.Atoi(strings.TrimSpace(maxPressureString))
			fmt.Println("Pressure Chunks:")
			pressureChunksString, _ := reader.ReadString('\n')
			pressureChunks, _ = strconv.Atoi(strings.TrimSpace(pressureChunksString))

			fmt.Println("Forward state")

			flowIterator := (maxFlow - minFlow) / flowChunks
			for i := minFlow; i <= maxFlow; i += flowIterator {
				xPoints = append(xPoints, Point{
					X: i,
					Y: borderA.Y,
				})
			}

			pressureIterator := (maxPressure - minPressure) / pressureChunks
			for i := minPressure; i <= maxPressure; i += pressureIterator {
				yPoints = append(yPoints, Point{
					X: borderA.X,
					Y: i,
				})
			}

			stateMachine = stateMachine.GetNext()
		}

		time.Sleep(100 * time.Millisecond)
	}
}

func MouseHandler(event int, x int, y int, flags int, userdata interface{}) {
	if event == Click {
		switch stateMachine.GetType() {
		case GetBorders:
			if borderCount == 0 {
				fmt.Printf("get border A, click, %d, %d, %d\n", x, y, flags)
				borderA.X = x
				borderA.Y = y
				borderCount++
			} else {
				fmt.Printf("get border B, click, %d, %d, %d\n", x, y, flags)
				borderB.X = x
				borderB.Y = y
				fmt.Println("Forward state")
				stateMachine = stateMachine.GetNext()
			}
		case GetScales:
			fmt.Printf("get scale y, click, %d, %d, %d\n", x, y, flags)
		case GetScaleX:
			fmt.Printf("get scale x, click, %d, %d, %d\n", x, y, flags)
		}

	}
}
