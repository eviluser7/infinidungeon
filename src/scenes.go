package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

// sceneTransitions detects what situation is going on and sets a transition
// from screen to screen so it won't feel as empty.
func sceneTransitions(g *Game, screen *ebiten.Image) {
	if g.scene == "transition" {
		if g.situation == "level1" {
			g.op.GeoM.Reset()
			g.op.GeoM.Translate(0.0, 0.0)
			screen.DrawImage(loadingScreen1, &g.op)
			g.activateTimer = true

			if g.sceneTimer == 120 && g.activateTimer {
				g.activateTimer = false
				g.scene = "STAGE1"
				g.sceneTimer = 6
			}
		}

		if g.situation == "level2" {
			g.op.GeoM.Reset()
			g.op.GeoM.Translate(0.0, 0.0)
			screen.DrawImage(loadingScreen2, &g.op)
			g.activateTimer = true

			if g.sceneTimer == 120 && g.activateTimer {
				g.activateTimer = false
				g.scene = "STAGE2"
				g.sceneTimer = 0
			}
		}

		if g.situation == "level3" {
			g.op.GeoM.Reset()
			g.op.GeoM.Translate(0.0, 0.0)
			screen.DrawImage(loadingScreen3, &g.op)
			g.activateTimer = true

			if g.sceneTimer == 120 && g.activateTimer {
				g.activateTimer = false
				g.scene = "STAGE3"
				g.sceneTimer = 0
			}
		}

		if g.situation == "level4" {
			g.op.GeoM.Reset()
			g.op.GeoM.Translate(0.0, 0.0)
			screen.DrawImage(loadingScreen4, &g.op)
			g.activateTimer = true

			if g.sceneTimer == 120 && g.activateTimer {
				g.activateTimer = false
				g.scene = "STAGE4"
				g.sceneTimer = 0
			}
		}

		if g.situation == "level5" {
			g.op.GeoM.Reset()
			g.op.GeoM.Translate(0.0, 0.0)
			screen.DrawImage(loadingScreen5, &g.op)
			g.activateTimer = true

			if g.sceneTimer == 120 && g.activateTimer {
				g.activateTimer = false
				g.scene = "STAGE5"
				g.sceneTimer = 0
			}
		}

		if g.situation == "level6" {
			g.op.GeoM.Reset()
			g.op.GeoM.Translate(0.0, 0.0)
			screen.DrawImage(loadingScreen6, &g.op)
			g.activateTimer = true

			if g.sceneTimer == 120 && g.activateTimer {
				g.activateTimer = false
				g.scene = "ENDSTAGE"
				g.sceneTimer = 0
			}
		}

		if g.situation == "end" {
			punches.SetVolume(1.0)
			punches.Play()
			g.activateTimer = true

			if g.sceneTimer == 240 && g.activateTimer {
				g.activateTimer = false
				g.scene = "menu"
				g.sceneTimer = 0
			}
		}
	}
}
