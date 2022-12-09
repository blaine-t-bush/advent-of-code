package day9

import (
	"fmt"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"

	util "github.com/blaine-t-bush/advent-of-code/util"
)

const (
	tileSize = 4
	w        = 1080
	h        = 1080
)

var (
	knotImage = ebiten.NewImage(tileSize, tileSize)
)

type Game struct {
	inited                 bool
	op                     *ebiten.DrawImageOptions
	commands               []command
	currentCommandCount    int
	currentCommandSteps    int
	currentCommandProgress int
	rope                   *rope
}

func (c coord) translateToScreenCenter() coord {
	return coord{
		c.x + w/2,
		c.y + h/2,
	}
}

func (g *Game) init() {
	defer func() {
		g.inited = true
	}()

	input := util.ReadInput(inputFile)
	g.commands = parseCommands(input)
	g.currentCommandCount = 0
	g.currentCommandSteps = g.commands[g.currentCommandCount].steps
	g.currentCommandProgress = 0
	g.rope = &rope{
		knots: map[int]coord{
			0: {0, 0},
			1: {0, 0},
			2: {0, 0},
			3: {0, 0},
			4: {0, 0},
			5: {0, 0},
			6: {0, 0},
			7: {0, 0},
			8: {0, 0},
			9: {0, 0},
		},
		tailVisitedCoords: []coord{
			{0, 0},
		},
	}

	g.op = &ebiten.DrawImageOptions{}
}

func (g *Game) Update() error {
	if !g.inited {
		g.init()
	}

	if g.currentCommandCount < len(g.commands) {
		if g.currentCommandProgress < g.currentCommandSteps {
			g.rope.move(g.commands[g.currentCommandCount].dir, 1)
			g.currentCommandProgress++
		}
	}

	if g.currentCommandCount < len(g.commands) {
		if g.currentCommandProgress == g.currentCommandSteps {
			g.currentCommandCount++
			g.currentCommandProgress = 0
			g.currentCommandSteps = g.commands[g.currentCommandCount].steps
		}
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, fmt.Sprintf("current command: %d\ncurrent steps:   %d\nx: %d, y: %d", g.currentCommandCount, g.currentCommandProgress, g.rope.knots[0].x, g.rope.knots[0].y))
	for i, knot := range g.rope.knots {
		resized := coord{knot.x * tileSize, knot.y * tileSize}
		translated := resized.translateToScreenCenter()
		g.op.GeoM.Reset()
		g.op.GeoM.Translate(float64(translated.x), float64(translated.y))
		g.op.ColorM.Scale(1, 1, 1, 1)
		transparency := 230*(g.rope.len()-i)/g.rope.len() + 25
		knotImage.Fill(color.NRGBA{255, 255, 255, uint8(transparency)})
		screen.DrawImage(knotImage, g.op)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return w, h
}

func Viz() {
	ebiten.SetWindowSize(w, h)
	ebiten.SetWindowTitle("Advent of Code 2022 - Day 9")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
