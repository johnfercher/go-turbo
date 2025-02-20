package main

import (
	"context"
	"fmt"
	"github.com/johnfercher/go-turbo/internal/adapters/csv"
	"github.com/johnfercher/go-turbo/internal/core/models"
	"github.com/johnfercher/go-turbo/internal/math"
	"gocv.io/x/gocv"
	"image"
	"image/color"
	"time"
)

const ESC = 27
const SPACE = 32
const Click = 1

var stateMachine = GetStateMachine()

var borderA *models.Point
var borderB *models.Point
var borderCount = 0

var xPoints []models.Point
var yPoints []models.Point

var efficiencyMap = make(map[float64][]models.Point)
var efficiencyPoints []models.Point
var efficiencyStep = 0.0
var maxEfficiency = 0.0
var currentEfficiency = 0.0
var exit = false

var xFlowRange *models.Range
var xPixelRange *models.Range
var chart *models.Chart

func main() {
	/*reader := bufio.NewReader(os.Stdin)

	fmt.Println("Turbo:")
	turboName, _ := reader.ReadString('\n')*/
	turboName := "kinugawa18g"

	go CLI()

	window := gocv.NewWindow("GOCV - gopher")
	img := gocv.IMRead(fmt.Sprintf("assets/images/%s.png", turboName), gocv.IMReadAnyColor)

	window.ResizeWindow(img.Cols(), img.Rows())
	window.SetMouseHandler(MouseHandler, nil)

	for !exit {
		if borderA != nil && borderB != nil {
			gocv.Rectangle(&img, image.Rect(borderA.X, borderA.Y, borderB.X, borderB.Y), color.RGBA{
				R: 255,
			}, 1)
		}

		if xPoints != nil {
			for _, point := range xPoints {
				gocv.Line(&img,
					image.Point{
						X: point.X,
						Y: point.Y,
					},
					image.Point{
						X: point.X,
						Y: point.Y + 5,
					},
					color.RGBA{
						R: 255,
					},
					5)
			}
		}

		if yPoints != nil {
			for _, point := range yPoints {
				gocv.Line(&img,
					image.Point{
						X: point.X,
						Y: point.Y,
					},
					image.Point{
						X: point.X + 5,
						Y: point.Y,
					},
					color.RGBA{
						B: 255,
					},
					5)
			}
		}

		for _, point := range efficiencyPoints {
			gocv.Line(&img,
				image.Point{
					X: point.X,
					Y: point.Y,
				},
				image.Point{
					X: point.X + 5,
					Y: point.Y,
				},
				color.RGBA{
					R: 255,
				},
				5)
		}

		for _, points := range efficiencyMap {
			for _, point := range points {
				gocv.Line(&img,
					image.Point{
						X: point.X,
						Y: point.Y,
					},
					image.Point{
						X: point.X + 5,
						Y: point.Y,
					},
					color.RGBA{
						B: 255,
					},
					5)
			}
		}

		window.IMShow(img)
		key := window.WaitKey(100)
		if key == ESC {
			exit = true
		} else if key == SPACE {
			if stateMachine.GetType() == GetEfficiency && currentEfficiency <= maxEfficiency {
				fmt.Printf("From %.2f to %.2f mapping\n", currentEfficiency, currentEfficiency+efficiencyStep)
				currentEfficiency = currentEfficiency + efficiencyStep
				efficiencyMap[currentEfficiency] = efficiencyPoints
				efficiencyPoints = nil
			} else if stateMachine.GetType() == GetEfficiency {
				exit = true
			}
		}
	}

	chart = models.NewChart(xPixelRange, xFlowRange)
	fmt.Println(chart)

	for key, value := range efficiencyMap {
		fmt.Printf("%.2f", key)
		for _, point := range value {
			fmt.Printf("%s", point.String())
		}
		fmt.Println()
	}

	chart.AddTurbo(efficiencyMap)

	turboRepo := csv.NewTurboRepository()
	err := turboRepo.Save(context.Background(), turboName, chart)
	if err != nil {
		panic(err)
	}
}

func CLI() {
	var minFlow float64
	var maxFlow float64
	var flowStep float64
	var minPressure float64
	var maxPressure float64
	var pressureStep float64

	for !exit {
		if stateMachine.GetType() == GetScales {
			/*reader := bufio.NewReader(os.Stdin)

			fmt.Println("Min Flow:")
			minFlowString, _ := reader.ReadString('\n')
			minFlow, _ = strconv.ParseFloat(strings.TrimSpace(minFlowString), 64)

			fmt.Println("Max Flow:")
			maxFlowString, _ := reader.ReadString('\n')
			maxFlow, _ = strconv.ParseFloat(strings.TrimSpace(maxFlowString), 64)

			fmt.Println("Flow Steps:")
			flowStepString, _ := reader.ReadString('\n')
			flowStep, _ = strconv.ParseFloat(strings.TrimSpace(flowStepString), 64)

			fmt.Println("Min Pressure:")
			minPressureString, _ := reader.ReadString('\n')
			minPressure, _ = strconv.ParseFloat(strings.TrimSpace(minPressureString), 64)

			fmt.Println("Max Pressure:")
			maxPressureString, _ := reader.ReadString('\n')
			maxPressure, _ = strconv.ParseFloat(strings.TrimSpace(maxPressureString), 64)

			fmt.Println("Pressure Step:")
			pressureStepString, _ := reader.ReadString('\n')
			pressureStep, _ = strconv.ParseFloat(strings.TrimSpace(pressureStepString), 64)*/

			minFlow = 0
			maxFlow = 675
			flowStep = 100
			minPressure = 1.0
			maxPressure = 3.4
			pressureStep = 0.2

			xFlowRange = models.NewRange(minFlow, maxFlow)
			fmt.Printf("flowRange: %s\n", xFlowRange)

			currentEfficiency = minPressure
			maxEfficiency = maxPressure
			efficiencyStep = pressureStep

			fmt.Println("Forward state")

			rateFlow := math.GetRate(minFlow, maxFlow, float64(borderA.X), float64(borderB.X))
			ratePressure := math.GetRate(minPressure, maxPressure, float64(borderA.Y), float64(borderB.Y))

			//fmt.Println(minFlow, minPressure, flowStep, rateFlow)
			fmt.Println(minPressure, maxPressure, pressureStep, ratePressure)

			for i := minFlow; i < maxFlow; i += flowStep {
				xPoints = append(xPoints, models.Point{
					X: borderA.X + int(i*rateFlow),
					Y: borderA.Y,
				})
			}

			for i := minPressure; i <= maxPressure+pressureStep; i += pressureStep {
				yPoints = append(yPoints, models.Point{
					X: borderA.X,
					Y: borderA.Y - int((i-1)*ratePressure),
				})
			}

			fmt.Println(xPoints)
			fmt.Println(yPoints)

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
				borderA = &models.Point{
					X: x,
					Y: y,
				}
				borderCount++
			} else {
				fmt.Printf("get border B, click, %d, %d, %d\n", x, y, flags)
				borderB = &models.Point{
					X: x,
					Y: y,
				}
				fmt.Println("Forward state")

				xPixelRange = models.NewRange(float64(borderA.X), float64(borderB.X))
				fmt.Printf("pixelRange: %s\n", xPixelRange)

				stateMachine = stateMachine.GetNext()
			}
		case GetScales:
			fmt.Printf("get scales, click, %d, %d, %d\n", x, y, flags)
		case GetEfficiency:
			fmt.Printf("get efficiency %.2f, click, %d, %d, %d\n", currentEfficiency, x, y, flags)
			point := models.Point{
				X: x,
				Y: y,
			}
			efficiencyPoints = append(efficiencyPoints, point)
		}
	}
}
