package main

import "github.com/hajimehoshi/ebiten/v2"

// DrawMaps draws maps
func DrawMaps(g *Game, screen *ebiten.Image) {
	// Set a general rule and then draw all stages
	g.op.GeoM.Reset()
	g.op.GeoM.Translate(0.0, 0.0)

	if g.scene == "STAGE1" {
		switch g.atLevel {
		case 0:
			screen.DrawImage(background0I, &g.op)
		case 1:
			screen.DrawImage(background1I, &g.op)
		case 2:
			screen.DrawImage(background2I, &g.op)
		case 3:
			screen.DrawImage(background3I, &g.op)
		case 4:
			screen.DrawImage(background4I, &g.op)
		case 5:
			screen.DrawImage(background5I, &g.op)
		case 6:
			screen.DrawImage(background6I, &g.op)
		case 7:
			screen.DrawImage(background7I, &g.op)
		case 8:
			screen.DrawImage(background8I, &g.op)
		case 9:
			screen.DrawImage(background9I, &g.op)
		case 10:
			screen.DrawImage(background10I, &g.op)
		case 11:
			screen.DrawImage(background11I, &g.op)
		case 12:
			screen.DrawImage(background12I, &g.op)
		case 13:
			screen.DrawImage(background13I, &g.op)
		case 14:
			screen.DrawImage(background14I, &g.op)
		case 15:
			screen.DrawImage(background15I, &g.op)
		case 16:
			screen.DrawImage(background16I, &g.op)
		case 17:
			screen.DrawImage(background17I, &g.op)
		case 18:
			if !g.player.enabledShrine1 {
				screen.DrawImage(background18preI, &g.op)
			} else {
				screen.DrawImage(background18postI, &g.op)
			}
		case 19:
			screen.DrawImage(background19I, &g.op)
		case 20:
			screen.DrawImage(background20I, &g.op)
		case 21:
			screen.DrawImage(background21I, &g.op)
		case 22:
			screen.DrawImage(background22I, &g.op)
		case 23:
			screen.DrawImage(background23I, &g.op)
		}
	}

	if g.scene == "STAGE2" {
		switch g.atLevel {
		case 0:
			screen.DrawImage(background0II, &g.op)
		case 1:
			screen.DrawImage(background1II, &g.op)
		case 2:
			screen.DrawImage(background2II, &g.op)
		case 3:
			screen.DrawImage(background3II, &g.op)
		case 4:
			screen.DrawImage(background4II, &g.op)
		case 5:
			screen.DrawImage(background5II, &g.op)
		case 6:
			screen.DrawImage(background6II, &g.op)
		case 7:
			screen.DrawImage(background7II, &g.op)
		case 8:
			screen.DrawImage(background8II, &g.op)
		case 9:
			screen.DrawImage(background9II, &g.op)
		case 10:
			screen.DrawImage(background10II, &g.op)
		case 11:
			screen.DrawImage(background11II, &g.op)
		case 12:
			screen.DrawImage(background12II, &g.op)
		case 13:
			screen.DrawImage(background13II, &g.op)
		case 14:
			screen.DrawImage(background14II, &g.op)
		case 15:
			screen.DrawImage(background15II, &g.op)
		case 16:
			screen.DrawImage(background16II, &g.op)
		case 17:
			screen.DrawImage(background17II, &g.op)
		case 18:
			if !g.player.enabledShrine2 {
				screen.DrawImage(background18preII, &g.op)
			} else {
				screen.DrawImage(background18postII, &g.op)
			}
		case 19:
			screen.DrawImage(background19II, &g.op)
		case 20:
			screen.DrawImage(background20II, &g.op)
		case 21:
			screen.DrawImage(background21II, &g.op)
		case 22:
			screen.DrawImage(background22II, &g.op)
		case 23:
			screen.DrawImage(background23II, &g.op)
		}
	}

	if g.scene == "STAGE3" {
		switch g.atLevel {
		case 0:
			screen.DrawImage(backgroundIII, &g.op)
			screen.DrawImage(arrowLeft, &g.op)
		case 1:
			screen.DrawImage(backgroundIII, &g.op)
			screen.DrawImage(arrowLeft, &g.op)
		case 2:
			screen.DrawImage(backgroundIII, &g.op)
			screen.DrawImage(arrowUp, &g.op)
		case 3:
			screen.DrawImage(backgroundIII, &g.op)
			screen.DrawImage(arrowLeft, &g.op)
		case 4:
			screen.DrawImage(backgroundIII, &g.op)
			screen.DrawImage(arrowUp, &g.op)
		case 5:
			screen.DrawImage(backgroundIII, &g.op)
			screen.DrawImage(arrowDown, &g.op)
		case 6:
			screen.DrawImage(backgroundIII, &g.op)
			screen.DrawImage(arrowUp, &g.op)
		case 7:
			screen.DrawImage(backgroundIII, &g.op)
			screen.DrawImage(arrowUp, &g.op)
		case 8:
			screen.DrawImage(backgroundIII, &g.op)
			screen.DrawImage(arrowDown, &g.op)
		case 9:
			screen.DrawImage(backgroundIII, &g.op)
			screen.DrawImage(arrowUp, &g.op)
		case 10:
			screen.DrawImage(backgroundIII, &g.op)
			screen.DrawImage(arrowLeft, &g.op)
		case 11:
			screen.DrawImage(backgroundIII, &g.op)
			screen.DrawImage(arrowUp, &g.op)
		case 12:
			screen.DrawImage(backgroundIII, &g.op)
			screen.DrawImage(arrowUp, &g.op)
		case 13:
			screen.DrawImage(backgroundIII, &g.op)
			screen.DrawImage(arrowLeft, &g.op)
		case 14:
			screen.DrawImage(backgroundIII, &g.op)
			screen.DrawImage(arrowUp, &g.op)
		case 15:
			screen.DrawImage(backgroundIII, &g.op)
			screen.DrawImage(arrowLeft, &g.op)
		case 16:
			screen.DrawImage(backgroundIII, &g.op)
			screen.DrawImage(arrowRD, &g.op)
		case 17:
			screen.DrawImage(backgroundIII, &g.op)
			screen.DrawImage(arrowUp, &g.op)
		case 18:
			screen.DrawImage(backgroundIII, &g.op)
			screen.DrawImage(arrowUp, &g.op)
		case 19:
			screen.DrawImage(backgroundIII, &g.op)
			screen.DrawImage(arrowRight, &g.op)
		case 20:
			screen.DrawImage(backgroundIII, &g.op)
			screen.DrawImage(circle, &g.op)
		case 21:
			screen.DrawImage(backgroundIII, &g.op)
			screen.DrawImage(arrowRight, &g.op)
		case 22:
			if !g.player.enabledShrine3 {
				screen.DrawImage(exitPreIII, &g.op)
			} else {
				screen.DrawImage(exitPostIII, &g.op)
				screen.DrawImage(shrineBlur3, &g.op)
			}
		case 23:
			screen.DrawImage(backgroundIII, &g.op)
			screen.DrawImage(arrowLR, &g.op)
		case 24:
			screen.DrawImage(backgroundIII, &g.op)
			screen.DrawImage(arrowRight, &g.op)
		}
	}

	if g.scene == "STAGE4" {
		screen.DrawImage(stage4, &g.op)
	}
}
