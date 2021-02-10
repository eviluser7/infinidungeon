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

	if g.atLevel == 18 && g.player.enabledShrine2 && g.scene == "STAGE2" {
		g.op.GeoM.Reset()
		g.op.GeoM.Translate(0.0, 0.0)
		screen.DrawImage(yellowBlur, &g.op)
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

	if g.atLevel == 22 && g.scene == "STAGE2" && !g.player.enabledShrine2 {
		g.op.GeoM.Reset()
		g.op.GeoM.Translate(0.0, 0.0)
		screen.DrawImage(shrineBlur2, &g.op)

		op := &ebiten.DrawImageOptions{}
		op.GeoM.Reset()
		op.GeoM.Translate(100.0, 130.0)
		screen.DrawImage(yellowShrine, op)
	}

	// Draw controls
	if !g.player.movedOnce && g.scene != "menu" {
		g.op.GeoM.Reset()
		g.op.GeoM.Translate(float64(g.player.x), float64(g.player.y))
		screen.DrawImage(wasd, &g.op)
	}

	// Draw player
	g.player.Draw(screen)

	// Show spacebar
	if g.scene == "STAGE1" && g.atLevel == 10 &&
		g.player.x >= 160 && g.player.x <= 236 &&
		g.player.y >= 95 && g.player.y <= 173 && !g.player.enabledShrine1 {
		g.op.GeoM.Reset()
		g.op.GeoM.Translate(170.0, 110.0)
		screen.DrawImage(spaceBar, &g.op)
	}

	if g.scene == "STAGE2" && g.atLevel == 22 &&
		g.player.x >= 74 && g.player.x <= 148 &&
		g.player.y >= 95 && g.player.y <= 173 && !g.player.enabledShrine2 {
		g.op.GeoM.Reset()
		g.op.GeoM.Translate(85.0, 110.0)
		screen.DrawImage(spaceBar, &g.op)
	}
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
