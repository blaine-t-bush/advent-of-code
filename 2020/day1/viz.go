package day1

import (
	"fmt"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"

	util "github.com/blaine-t-bush/advent-of-code/util"
)

const (
	fontWidth = 4
)

var (
	w = 640
	h = 640
)

type Game struct {
	inited  bool
	started bool
	op      *ebiten.DrawImageOptions
}

func (g *Game) updateWindowSize() {
	w = 640
	h = 640
}

func (g *Game) init() {
	defer func() {
		g.inited = true
	}()

	input := util.ReadInput(inputFile)
	fmt.Println(len(input))

	g.op = &ebiten.DrawImageOptions{}
	g.started = false
	g.updateWindowSize()
}

func (g *Game) Update() error {
	if !g.inited {
		g.init()
	}

	if !g.started {
		if ebiten.IsKeyPressed(ebiten.KeySpace) {
			g.started = true
		}
	} else {
		// do updates
		return nil
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	if !g.started {
		ebitenutil.DebugPrintAt(screen, "press space to start", w/2-len("press space to start")*fontWidth/2, h/2-fontWidth/2)
	} else {
		ebitenutil.DebugPrint(screen, "Hello, world!")
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return w, h
}

func Viz() {
	ebiten.SetWindowSize(w, h)
	ebiten.SetWindowTitle("Advent of Code")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
