package day0

import (
	"fmt"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"

	util "github.com/blaine-t-bush/advent-of-code/util"
)

const (
	w = 1080
	h = 1080
)

type Game struct {
	inited bool
	op     *ebiten.DrawImageOptions
}

func (g *Game) init() {
	defer func() {
		g.inited = true
	}()

	g.op = &ebiten.DrawImageOptions{}
}

func (g *Game) Update() error {
	if !g.inited {
		g.init()
	}

	input := util.ReadInput(inputFile)
	fmt.Println(len(input))

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "Hello, world!")
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
