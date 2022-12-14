package day14

import (
	"fmt"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"

	util "github.com/blaine-t-bush/advent-of-code/util"
)

const (
	fontWidth = 4
	tileSize  = 4
)

var (
	w          = 640
	h          = 820
	rockImage  = ebiten.NewImage(tileSize, tileSize)
	sandImage  = ebiten.NewImage(tileSize, tileSize)
	rocksImage *ebiten.Image
	sandsImage *ebiten.Image
)

type Game struct {
	inited           bool
	started          bool
	finished         bool
	op               *ebiten.DrawImageOptions
	rocksAndSand     map[coord]int
	defaultSandCoord coord
	currentSandCoord coord
}

func (c coord) adjustToScreen() coord {
	r := coord{
		c.x * tileSize,
		c.y * tileSize,
	}

	return coord{
		r.x - tileSize*500 + w/2,
		r.y + 50,
	}
}

func (g *Game) updateWindowSize() {
	minX, maxX, minY, maxY := getBounds(g.rocksAndSand)
	fmt.Printf("(%d, %d) (%d, %d)", minX, maxX, minY, maxY)
	// w = (maxX - minX)
	// h = (maxY - minY)
}

func (g *Game) init() {
	defer func() {
		g.inited = true
	}()

	input := util.ReadInput(inputFile)

	g.op = &ebiten.DrawImageOptions{}
	g.started = false
	g.finished = false
	g.rocksAndSand = addFloor(parseRocks(input))
	g.defaultSandCoord = coord{500, 0}
	g.currentSandCoord = coord{500, 0}
	g.updateWindowSize()

	rockImage.Fill(color.NRGBA{66, 66, 66, 255})
	sandImage.Fill(color.NRGBA{240, 221, 96, 255})
	rocksImage = ebiten.NewImage(w, h)
	sandsImage = ebiten.NewImage(w, h)

	for c, t := range g.rocksAndSand {
		translated := c.adjustToScreen()
		g.op.GeoM.Reset()
		g.op.GeoM.Translate(float64(translated.x), float64(translated.y))
		g.op.ColorM.Scale(1, 1, 1, 1)
		if t == terrainRock {
			rocksImage.DrawImage(rockImage, g.op)
		}
	}
}

func (g *Game) Update() error {
	if !g.inited {
		g.init()
	}

	if !g.started {
		if ebiten.IsKeyPressed(ebiten.KeySpace) {
			g.started = true
		}
	} else if !g.finished {
		// do updates
		newSandCoord := moveSandOne(g.currentSandCoord, g.rocksAndSand)

		if newSandCoord == g.defaultSandCoord {
			// sand couldn't move from start position
			g.finished = true
		} else if newSandCoord == g.currentSandCoord {
			// sand is not in start position but has settled to final coord.
			// spawn a new sand at start position.
			g.currentSandCoord = g.defaultSandCoord
			g.rocksAndSand[g.currentSandCoord] = terrainSand

			// since it has settled, we can draw it to the permanent sand image.
			translated := newSandCoord.adjustToScreen()
			g.op.GeoM.Reset()
			g.op.GeoM.Translate(float64(translated.x), float64(translated.y))
			g.op.ColorM.Scale(1, 1, 1, 1)
			sandsImage.DrawImage(sandImage, g.op)
		} else {
			// sand is moving.
			delete(g.rocksAndSand, g.currentSandCoord)
			g.currentSandCoord = newSandCoord
			g.rocksAndSand[g.currentSandCoord] = terrainSand
		}
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	if !g.started {
		ebitenutil.DebugPrintAt(screen, "press space to start", w/2-len("press space to start")*fontWidth/2, h/2-fontWidth/2)
	} else {

		// draw permanent rocks and sand
		g.op.GeoM.Reset()
		g.op.ColorM.Scale(1, 1, 1, 1)
		screen.DrawImage(rocksImage, g.op)
		screen.DrawImage(sandsImage, g.op)

		// draw active sand grain
		translated := g.currentSandCoord.adjustToScreen()
		g.op.GeoM.Reset()
		g.op.GeoM.Translate(float64(translated.x), float64(translated.y))
		g.op.ColorM.Scale(1, 1, 1, 1)
		screen.DrawImage(sandImage, g.op)
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
