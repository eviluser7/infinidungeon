package main

import (
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

// Game properties
type Game struct {
	inited           bool
	op               ebiten.DrawImageOptions
	player           *Player
	scene            string
	atLevel          int
	activateTimer    bool
	sceneTimer       int
	situation        string
	achievementTimer int
}

// Update the game state
func (g *Game) Update() error {
	if !g.inited {
		g.init()
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		os.Exit(0)
		//fmt.Println(g.player.x, g.player.y)
		//fmt.Println(g.atLevel)
	}

	// Grab shrines
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		if g.scene == "STAGE1" && g.atLevel == 10 &&
			g.player.x >= 160 && g.player.x <= 236 &&
			g.player.y >= 95 && g.player.y <= 173 && !g.player.enabledShrine1 {
			g.player.enabledShrine1 = true
			enableShrine.Rewind()
			enableShrine.Play()
		}

		if g.scene == "STAGE2" && g.atLevel == 22 &&
			g.player.x >= 74 && g.player.x <= 148 &&
			g.player.y >= 95 && g.player.y <= 173 && !g.player.enabledShrine2 {
			g.player.enabledShrine2 = true
			enableShrine.Rewind()
			enableShrine.Play()
		}

		if g.scene == "STAGE3" && g.atLevel == 20 &&
			g.player.x >= 128 && g.player.x <= 206 &&
			g.player.y >= 95 && g.player.y <= 173 && !g.player.enabledShrine3 {
			g.player.enabledShrine3 = true
			enableShrine.Rewind()
			enableShrine.Play()
		}

		if g.scene == "STAGE5" && g.atLevel == 43 &&
			g.player.x >= 128 && g.player.x <= 206 &&
			g.player.y >= 95 && g.player.y <= 173 && !g.player.enabledLastShrine {
			g.player.enabledLastShrine = true
			enableShrine.Rewind()
			enableShrine.Play()
		}

		if g.scene == "ENDSTAGE" && g.atLevel == 3 &&
			g.player.x >= 130 && g.player.x <= 192 &&
			g.player.y >= 181 && g.player.y <= 207 && !g.player.talkedToEvil {
			surprise.Rewind()
			surprise.Play()
			g.player.talkedToEvil = true
			g.activateTimer = true
		}
	}

	if g.activateTimer {
		//fmt.Println(g.sceneTimer)
		g.sceneTimer++
	}

	// Achievements
	if g.player.enabledShrine1 && !g.player.gotAchievement1 {
		g.player.gotAchievement = true
		g.player.gotAchievement1 = true
	}

	if g.scene == "STAGE2" && !g.player.gotAchievement2 {
		g.player.gotAchievement = true
		g.player.gotAchievement2 = true
	}

	if g.player.enabledShrine2 && !g.player.gotAchievement3 {
		g.player.gotAchievement = true
		g.player.gotAchievement3 = true
	}

	if g.scene == "STAGE3" && !g.player.gotAchievement4 {
		g.player.gotAchievement = true
		g.player.gotAchievement4 = true
	}

	if g.player.enabledShrine3 && !g.player.gotAchievement5 {
		g.player.gotAchievement = true
		g.player.gotAchievement5 = true
	}

	if g.scene == "STAGE4" && !g.player.gotAchievement6 {
		g.player.gotAchievement = true
		g.player.gotAchievement6 = true
	}

	if g.scene == "STAGE5" && !g.player.gotAchievement7 {
		g.player.gotAchievement = true
		g.player.gotAchievement7 = true
	}

	if g.player.enabledLastShrine && !g.player.gotAchievement8 {
		g.player.gotAchievement = true
		g.player.gotAchievement8 = true
	}

	if g.scene == "ENDSTAGE" && !g.player.gotAchievement9 {
		g.player.gotAchievement = true
		g.player.gotAchievement9 = true
	}

	// Achievement timer events
	if g.player.gotAchievement {
		g.achievementTimer++
	}
	if g.achievementTimer >= 240 {
		g.player.gotAchievement = false
		g.achievementTimer = 0
	}

	// Update player
	if g.scene != "menu" && g.scene != "transition" {
		g.player.Update(g)
	}

	switch g.scene {
	case "STAGE1":
		g.player.UpdateMap01(g)
	case "STAGE2":
		g.player.UpdateMap02(g)
	case "STAGE3":
		g.player.UpdateMap03(g)
	case "STAGE4":
		g.player.UpdateMap04(g)
	case "STAGE5":
		g.player.UpdateMap05(g)
	case "ENDSTAGE":
		g.player.UpdateMap06(g)
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		if g.scene == "menu" {
			punches.SetVolume(0.0)
			punches.Rewind()
			g.player.talkedToEvil = false
			g.sceneTimer = 0
			g.situation = "level1"
			g.scene = "transition"
			g.atLevel = 0
			g.player.y = 120
			g.player.x = 160
		}
	}

	if g.sceneTimer == 120 && g.scene == "ENDSTAGE" {
		g.situation = "end"
		g.scene = "transition"
		g.activateTimer = false
	}

	return nil
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

	if g.scene == "STAGE1" || g.scene == "STAGE2" || g.scene == "STAGE3" {
		screen.DrawImage(blur, &g.op)
	} else if g.scene == "STAGE4" {
		screen.DrawImage(greenBlur, &g.op)
	} else if g.scene == "STAGE5" {
		screen.DrawImage(moreBlur, &g.op)
	}

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

	// Draw shrine blurs
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

	if g.atLevel == 20 && g.scene == "STAGE3" && !g.player.enabledShrine3 {
		g.op.GeoM.Reset()
		g.op.GeoM.Translate(0.0, 0.0)
		screen.DrawImage(shrineBlur3, &g.op)

		op := &ebiten.DrawImageOptions{}
		op.GeoM.Reset()
		op.GeoM.Translate(155.0, 140.0)
		screen.DrawImage(greenShrine, op)
	}

	if g.atLevel == 43 && g.scene == "STAGE5" && !g.player.enabledLastShrine {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Reset()
		op.GeoM.Translate(0.0, 0.0)
		screen.DrawImage(whiteShrine, op)
	}

	// Draw controls
	if !g.player.movedOnce && g.scene != "menu" {
		g.op.GeoM.Reset()
		g.op.GeoM.Translate(float64(g.player.x), float64(g.player.y))
		screen.DrawImage(wasd, &g.op)
	}

	// Draw player
	if g.scene != "menu" && g.scene != "transition" {
		g.player.Draw(screen)
	}

	// Draw Mr. Evil
	if g.scene == "ENDSTAGE" {
		if g.atLevel == 3 && !g.player.talkedToEvil {
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Reset()
			op.GeoM.Translate(125.0, 203.0)
			screen.DrawImage(evilStanding1, op)
		} else if g.atLevel == 3 && g.player.talkedToEvil {
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Reset()
			op.GeoM.Translate(125.0, 203.0)
			screen.DrawImage(evilStanding2, op)

			op.GeoM.Reset()
			op.GeoM.Translate(170.0, 203.0)
			screen.DrawImage(exclamation, op)
		}
	}

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

	if g.scene == "STAGE3" && g.atLevel == 20 &&
		g.player.x >= 128 && g.player.x <= 206 &&
		g.player.y >= 95 && g.player.y <= 173 && !g.player.enabledShrine3 {
		g.op.GeoM.Reset()
		g.op.GeoM.Translate(140.0, 120.0)
		screen.DrawImage(spaceBar, &g.op)
	}

	if g.scene == "STAGE5" && g.atLevel == 43 &&
		g.player.x >= 128 && g.player.x <= 206 &&
		g.player.y >= 95 && g.player.y <= 173 && !g.player.enabledLastShrine {
		g.op.GeoM.Reset()
		g.op.GeoM.Translate(140.0, 120.0)
		screen.DrawImage(spaceBar, &g.op)
	}

	if g.scene == "ENDSTAGE" && g.atLevel == 3 &&
		g.player.x >= 130 && g.player.x <= 192 &&
		g.player.y >= 181 && g.player.y <= 207 && !g.player.talkedToEvil {
		g.op.GeoM.Reset()
		g.op.GeoM.Translate(135.0, 190.0)
		screen.DrawImage(spaceBar, &g.op)
	}

	// Menu
	if g.scene == "menu" {
		g.op.GeoM.Reset()
		g.op.GeoM.Translate(0.0, 0.0)
		screen.DrawImage(backgroundMenu, &g.op)
		screen.DrawImage(menu, &g.op)
		screen.DrawImage(menuCredits, &g.op)
	}

	// Show achievements
	if g.scene != "menu" && g.scene != "transition" &&
		g.player.gotAchievement && g.achievementTimer <= 239 {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Reset()
		op.GeoM.Translate(100.0, 200.0)
		screen.DrawImage(achBar, op)

		if g.player.gotAchievement1 && !g.player.gotAchievement2 {
			screen.DrawImage(ach1, op)
		} else if g.player.gotAchievement2 && !g.player.gotAchievement3 {
			screen.DrawImage(ach2, op)
		} else if g.player.gotAchievement3 && !g.player.gotAchievement4 {
			screen.DrawImage(ach3, op)
		} else if g.player.gotAchievement4 && !g.player.gotAchievement5 {
			screen.DrawImage(ach4, op)
		} else if g.player.gotAchievement5 && !g.player.gotAchievement6 {
			screen.DrawImage(ach5, op)
		} else if g.player.gotAchievement6 && !g.player.gotAchievement7 {
			screen.DrawImage(ach6, op)
		} else if g.player.gotAchievement7 && !g.player.gotAchievement8 {
			screen.DrawImage(ach7, op)
		} else if g.player.gotAchievement8 && !g.player.gotAchievement9 {
			screen.DrawImage(ach8, op)
		} else if g.player.gotAchievement9 {
			screen.DrawImage(ach9, op)
		}
	}

	sceneTransitions(g, screen)
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
