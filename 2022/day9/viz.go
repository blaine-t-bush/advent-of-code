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
	tileSize  = 2
	fontWidth = 5
)

var (
	knotImage  = ebiten.NewImage(tileSize, tileSize)
	pathImage  = ebiten.NewImage(tileSize, tileSize)
	w          = 640
	h          = 640
	pathsImage *ebiten.Image
)

type Game struct {
	inited                 bool
	started                bool
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
		-c.y + h/2,
	}
}

func (g *Game) init() {
	defer func() {
		g.inited = true
	}()

	input := util.ReadInput(inputFile)
	g.commands = parseCommands(input)
	g.resetState()

	// step through all commands to determine the max and min X and Y coords to set screen size
	xCoords := []int{0}
	yCoords := []int{0}
	for {
		g.stepThroughCommands()
		for _, knot := range g.rope.knots {
			xCoords = append(xCoords, knot.x)
			yCoords = append(yCoords, knot.y)
		}

		if g.currentCommandCount >= len(g.commands)-1 {
			break
		}
	}

	w = tileSize * (util.MaxIntsSlice(xCoords) - util.MinIntsSlice(xCoords) + 50)
	h = tileSize * (util.MaxIntsSlice(yCoords) - util.MinIntsSlice(yCoords) + 50)
	ebiten.SetWindowSize(w, h)

	g.resetState()

	g.op = &ebiten.DrawImageOptions{}
	pathsImage = ebiten.NewImage(w, h)
}

func (g *Game) Update() error {
	if !g.inited {
		g.init()
	}

	if g.started {
		g.stepThroughCommands()
	} else {
		if ebiten.IsKeyPressed(ebiten.KeySpace) {
			g.started = true
		}
	}

	return nil
}

func (g *Game) resetState() {
	g.started = false
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
			// 10: {0, 0},
			// 11: {0, 0},
			// 12: {0, 0},
			// 13: {0, 0},
			// 14: {0, 0},
			// 15: {0, 0},
			// 16: {0, 0},
			// 17: {0, 0},
			// 18: {0, 0},
			// 19: {0, 0},
			// 20: {0, 0},
			// 21: {0, 0},
			// 22: {0, 0},
			// 23: {0, 0},
			// 24: {0, 0},
			// 25: {0, 0},
			// 26: {0, 0},
			// 27: {0, 0},
			// 28: {0, 0},
			// 29: {0, 0},
			// 30: {0, 0},
			// 31: {0, 0},
			// 32: {0, 0},
			// 33: {0, 0},
			// 34: {0, 0},
			// 35: {0, 0},
			// 36: {0, 0},
			// 37: {0, 0},
			// 38: {0, 0},
			// 39: {0, 0},
			// 40: {0, 0},
			// 41: {0, 0},
			// 42: {0, 0},
			// 43: {0, 0},
			// 44: {0, 0},
			// 45: {0, 0},
			// 46: {0, 0},
			// 47: {0, 0},
			// 48: {0, 0},
			// 49: {0, 0},
		},
		tailVisitedCoords: []coord{
			{0, 0},
		},
	}
}

func (g *Game) stepThroughCommands() {
	if g.currentCommandCount < len(g.commands) {
		if g.currentCommandProgress < g.currentCommandSteps {
			g.rope.move(g.commands[g.currentCommandCount].dir, 1)
			g.currentCommandProgress++
		}
	}

	if g.currentCommandCount < len(g.commands)-1 && g.currentCommandProgress == g.currentCommandSteps {
		g.currentCommandCount++
		g.currentCommandProgress = 0
		g.currentCommandSteps = g.commands[g.currentCommandCount].steps
	}
}

func (g *Game) Draw(screen *ebiten.Image) {
	if !g.started {
		ebitenutil.DebugPrintAt(screen, "press space to start", w/2-10*fontWidth, h/2-fontWidth)
	} else {
		ebitenutil.DebugPrint(screen, fmt.Sprintf("current command: %d\ncurrent steps:   %d\nx: %d, y: %d", g.currentCommandCount, g.currentCommandProgress, g.rope.knots[0].x, g.rope.knots[0].y))

		// draw traversed path
		lastVisited := g.rope.tailVisitedCoords[len(g.rope.tailVisitedCoords)-1]
		resized := coord{lastVisited.x * tileSize, lastVisited.y * tileSize}
		translated := resized.translateToScreenCenter()
		g.op.GeoM.Reset()
		g.op.GeoM.Translate(float64(translated.x), float64(translated.y))
		g.op.ColorM.Scale(1, 1, 1, 1)
		pathImage.Fill(color.NRGBA{255, 0, 0, 50})
		pathsImage.DrawImage(pathImage, g.op)
		g.op.GeoM.Reset()
		screen.DrawImage(pathsImage, g.op)

		// draw knots
		for i, knot := range g.rope.knots {
			resized := coord{knot.x * tileSize, knot.y * tileSize}
			translated := resized.translateToScreenCenter()
			g.op.GeoM.Reset()
			g.op.GeoM.Translate(float64(translated.x), float64(translated.y))
			g.op.ColorM.Scale(1, 1, 1, 1)
			transparency := 220*(g.rope.len()-i)/g.rope.len() + 35
			knotImage.Fill(color.NRGBA{255, 255, 255, uint8(transparency)})
			screen.DrawImage(knotImage, g.op)
		}
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
