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
	if g.scene == "STAGE1" {
		// Set a general rule and then draw all stages
		g.op.GeoM.Reset()
		g.op.GeoM.Translate(0.0, 0.0)

		switch g.atLevel {
		case 0:
			screen.DrawImage(background0, &g.op)
		case 1:
			screen.DrawImage(background1, &g.op)
		case 2:
			screen.DrawImage(background2, &g.op)
		case 3:
			screen.DrawImage(background3, &g.op)
		case 4:
			screen.DrawImage(background4, &g.op)
		case 5:
			screen.DrawImage(background5, &g.op)
		case 6:
			screen.DrawImage(background6, &g.op)
		case 7:
			screen.DrawImage(background7, &g.op)
		case 8:
			screen.DrawImage(background8, &g.op)
		case 9:
			screen.DrawImage(background9, &g.op)
		case 10:
			screen.DrawImage(background10, &g.op)
		case 11:
			screen.DrawImage(background11, &g.op)
		case 12:
			screen.DrawImage(background12, &g.op)
		case 13:
			screen.DrawImage(background13, &g.op)
		case 14:
			screen.DrawImage(background14, &g.op)
		case 15:
			screen.DrawImage(background15, &g.op)
		case 16:
			screen.DrawImage(background16, &g.op)
		case 17:
			screen.DrawImage(background17, &g.op)
		case 18:
			if !g.player.enabledShrine1 {
				screen.DrawImage(background18pre, &g.op)
			} else {
				screen.DrawImage(background18post, &g.op)
			}
		case 19:
			screen.DrawImage(background19, &g.op)
		case 20:
			screen.DrawImage(background20, &g.op)
		case 21:
			screen.DrawImage(background21, &g.op)
		case 22:
			screen.DrawImage(background22, &g.op)
		}

		// Draw blur
		blurW, blurH := blur.Size()
		g.op.GeoM.Reset()
		g.op.GeoM.Translate(-float64(blurW)/2, -float64(blurH)/2) // Anchor
		g.op.GeoM.Translate(float64(g.player.x), float64(g.player.y))
		//screen.DrawImage(blur, &g.op)

		if g.atLevel == 18 && g.player.enabledShrine1 && g.scene == "STAGE1" {
			g.op.GeoM.Reset()
			g.op.GeoM.Translate(0.0, 0.0)
			screen.DrawImage(blueBlur, &g.op)
		}

		// Draw player
		g.player.Draw(screen)
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
