package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

// Game properties
type Game struct {
	inited  bool
	op      ebiten.DrawImageOptions
	player  *Player
	scene   string
	atLevel int
}

// Draw renders the screen
func (g *Game) Draw(screen *ebiten.Image) {
	// Draw maps
	DrawMaps(g, screen)

	// Draw blur
	blurW, blurH := blur.Size()
	g.op.GeoM.Reset()
	g.op.GeoM.Translate(-float64(blurW)/2, -float64(blurH)/2) // Anchor
	g.op.GeoM.Translate(float64(g.player.x), float64(g.player.y))
	screen.DrawImage(blur, &g.op)

	if g.atLevel == 18 && g.player.enabledShrine1 && g.scene == "STAGE1" {
		g.op.GeoM.Reset()
		g.op.GeoM.Translate(0.0, 0.0)
		screen.DrawImage(blueBlur, &g.op)
	}

	if g.atLevel == 10 && g.scene == "STAGE1" && !g.player.enabledShrine1 {
		g.op.GeoM.Reset()
		g.op.GeoM.Translate(0.0, 0.0)
		screen.DrawImage(shrineBlur1, &g.op)

		op := &ebiten.DrawImageOptions{}
		op.GeoM.Reset()
		op.GeoM.Translate(185.0, 132.0)
		screen.DrawImage(blueShrine, op)
	}

	// Draw player
	g.player.Draw(screen)
}

// Layout is the window size
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func (g *Game) init() {
	defer func() {
		g.inited = true
	}()

	g.player = &Player{
		x: 160,
		y: 145,
		w: 32,
		h: 32,
	}
}
