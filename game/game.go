package game

import (
	"AdventOfCode2020/days"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"log"
	"time"
)

type Game struct {
	currentDay  int
	givenDay    int
	nextDay     int
	isSetup     bool
	startTime   int64
	inputLines  []string
	outputLines string
	inputFile   string
	stepper     func(*ebiten.Image, int64) ([]string, bool)
}

func Init(day int, inputFile string) {
	g := Game{currentDay: day, givenDay: day, inputFile: inputFile}
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Advent of Code 2020")

	if err := ebiten.RunGame(&g); err != nil {
		log.Fatal(err)
	}
}

func (g *Game) Update() error {
	if g.currentDay == -1 {
		if inpututil.IsKeyJustPressed(ebiten.KeyEnter) || inpututil.IsKeyJustPressed(ebiten.KeyNumpadEnter) {
			g.currentDay = g.nextDay
		} else {
			return nil
		}
	}
	day := days.Get(g.currentDay)

	if !g.isSetup {
		if g.givenDay > 0 && g.inputFile != "" {
			g.inputLines = days.GetInput(g.inputFile)
		} else {
			g.inputLines = days.GetInputLines(g.currentDay)
		}

		if day.VisualSetup != nil {
			day.VisualSetup(g.inputLines)
		}
		g.startTime = time.Now().UnixNano()
		g.stepper = day.VisualStep
		g.isSetup = true
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	if g.currentDay == -1 {
		ebitenutil.DebugPrint(screen, g.outputLines+fmt.Sprintf("\nPress Enter to execute day #%d...", g.nextDay))
		return
	}
	day := days.Get(g.currentDay)

	if day.Number < days.NumberDays() {
		var stepDone bool
		var output []string

		if day.VisualStep != nil {
			output, stepDone = day.VisualStep(screen, time.Now().UnixNano()-g.startTime)
		} else {
			output, stepDone = day.CLI(g.inputLines), true
		}

		lines := ""
		for _, line := range output {
			lines += line + "\n"
			fmt.Println(line)
		}

		if day.VisualStep == nil {
			ebitenutil.DebugPrint(screen, lines)
		}

		if stepDone {
			g.outputLines = lines
			if g.givenDay > 0 && g.currentDay == g.givenDay {
				return
			}

			if g.currentDay == days.NumberDays()-1 {
				g.currentDay = -1
				g.isSetup = false
				g.nextDay = 0
			}

			if g.givenDay == -1 {
				g.nextDay = g.currentDay + 1
				g.isSetup = false
				g.currentDay = -1
			}
		}

	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}
