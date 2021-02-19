package main

import (
	"fmt"
	"image/color"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
)

// Game properties
type Game struct {
	inited                   bool
	op                       ebiten.DrawImageOptions
	player                   *Player
	scene                    string
	atLevel                  int
	activateTimer            bool
	sceneTimer               int
	situation                string
	achievementTimer         int
	achievementPage          int
	dialogueTimer            int
	dialogueShowing          bool
	whatDialogue             int
	amountOfDialoguesInScene int
	mouseX                   int
	mouseY                   int
}

// Update the game state
func (g *Game) Update() error {
	x, y := ebiten.CursorPosition()
	g.mouseX = x
	g.mouseY = y

	if !g.inited {
		g.init()
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		os.Exit(0)
	}

	// Grab shrines
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		if g.scene == "STAGE1" && g.atLevel == 10 &&
			g.player.x >= 160 && g.player.x <= 236 &&
			g.player.y >= 95 && g.player.y <= 173 && !g.player.enabledShrine1 {
			g.whatDialogue = 0
			g.dialogueShowing = true
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
	if g.scene != "menu" && g.scene != "transition" && g.scene != "achievement" {
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
			g.atLevel = 6
			g.player.y = 120
			g.player.x = 160
		}

		if g.dialogueShowing && g.scene != "menu" &&
			g.scene != "transition" {
			g.whatDialogue++
		}
	}

	if !g.dialogueShowing && g.scene != "menu" && g.scene != "transition" &&
		g.scene != "achievement" && g.whatDialogue < g.amountOfDialoguesInScene {
		g.dialogueShowing = true
	}

	if g.sceneTimer == 120 && g.scene == "ENDSTAGE" {
		g.situation = "end"
		g.scene = "transition"
		g.activateTimer = false
	}

	// Change achievement screen page
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) && g.scene == "achievement" {
		if x >= 20 && x <= 60 &&
			y >= 120 && y <= 160 &&
			g.achievementPage > 1 {
			g.achievementPage--
		}

		if x >= 270 && x <= 310 &&
			y >= 120 && y <= 160 &&
			g.achievementPage < 3 {
			g.achievementPage++
		}

		if x >= 20 && x <= 60 &&
			y >= 10 && y <= 50 {
			g.scene = "menu"
		}
	}

	// Get into achievement screen
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) &&
		x >= 10 && x <= 50 &&
		y >= 180 && y <= 220 {
		g.scene = "achievement"
	}

	// Change number of dialogues per scene
	if g.scene == "STAGE1" && !g.player.enabledShrine1 {
		g.amountOfDialoguesInScene = 5
	} else if g.player.enabledShrine1 && g.scene == "STAGE1" {
		g.amountOfDialoguesInScene = 2
	}

	if g.scene == "STAGE2" {
		g.amountOfDialoguesInScene = 7
	}

	if g.scene == "STAGE3" {
		g.amountOfDialoguesInScene = 5
	}

	if g.scene == "STAGE4" {
		g.amountOfDialoguesInScene = 3
	}

	if g.scene == "STAGE5" {
		g.amountOfDialoguesInScene = 7
	}

	if g.scene == "ENDSTAGE" {
		g.amountOfDialoguesInScene = 0
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
	if !g.player.movedOnce && g.scene != "menu" && g.scene != "achievement" {
		g.op.GeoM.Reset()
		g.op.GeoM.Translate(float64(g.player.x), float64(g.player.y))
		screen.DrawImage(wasd, &g.op)
	}

	// Draw player
	if g.scene != "menu" && g.scene != "transition" && g.scene != "achievement" {
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

		g.op.GeoM.Reset()
		g.op.GeoM.Translate(10.0, 180.0)
		screen.DrawImage(achievementButton, &g.op)
	}

	// Show achievements
	if g.scene != "menu" && g.scene != "transition" &&
		g.player.gotAchievement && g.achievementTimer <= 239 {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Reset()
		op.GeoM.Translate(100.0, 10.0)
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

	// Achievements screen
	if g.scene == "achievement" {
		g.op.GeoM.Reset()
		g.op.GeoM.Translate(0.0, 0.0)
		screen.DrawImage(backgroundMenu, &g.op)

		op1 := &ebiten.DrawImageOptions{}
		op2 := &ebiten.DrawImageOptions{}
		op3 := &ebiten.DrawImageOptions{}
		text.Draw(screen, "Achievements", pixeledFont, 110, 20, color.White)
		text.Draw(screen, fmt.Sprint(g.achievementPage), pixeledFont, 155, 230, color.White)

		op1.GeoM.Reset()
		op1.GeoM.Translate(100.0, 50.0)

		op2.GeoM.Reset()
		op2.GeoM.Translate(100.0, 115.0)

		op3.GeoM.Reset()
		op3.GeoM.Translate(100.0, 180.0)

		// Draw achievements if player has them
		if g.player.gotAchievement1 && g.achievementPage == 1 {
			screen.DrawImage(achBar2, op1)

			op1.GeoM.Reset()
			op1.GeoM.Translate(100.0, 45.0) // Achievement in question
			screen.DrawImage(ach1, op1)
		}

		if g.player.gotAchievement2 && g.achievementPage == 1 {
			screen.DrawImage(achBar2, op2)

			op2.GeoM.Reset()
			op2.GeoM.Translate(100.0, 110.0) // Achievement in question
			screen.DrawImage(ach2, op2)
		}

		if g.player.gotAchievement3 && g.achievementPage == 1 {
			screen.DrawImage(achBar2, op3)

			op3.GeoM.Reset()
			op3.GeoM.Translate(100.0, 175.0)
			screen.DrawImage(ach3, op3)
		}

		if g.player.gotAchievement4 && g.achievementPage == 2 {
			screen.DrawImage(achBar2, op1)

			op1.GeoM.Reset()
			op1.GeoM.Translate(100.0, 45.0)
			screen.DrawImage(ach4, op1)
		}

		if g.player.gotAchievement5 && g.achievementPage == 2 {
			screen.DrawImage(achBar2, op2)

			op2.GeoM.Reset()
			op2.GeoM.Translate(100.0, 110.0)
			screen.DrawImage(ach5, op2)
		}

		if g.player.gotAchievement6 && g.achievementPage == 2 {
			screen.DrawImage(achBar2, op3)

			op3.GeoM.Reset()
			op3.GeoM.Translate(100.0, 175.0)
			screen.DrawImage(ach6, op3)
		}

		if g.player.gotAchievement7 && g.achievementPage == 3 {
			screen.DrawImage(achBar2, op1)

			op1.GeoM.Reset()
			op1.GeoM.Translate(100.0, 45.0)
			screen.DrawImage(ach7, op1)
		}

		if g.player.gotAchievement8 && g.achievementPage == 3 {
			screen.DrawImage(achBar2, op2)

			op2.GeoM.Reset()
			op2.GeoM.Translate(100.0, 110.0)
			screen.DrawImage(ach8, op2)
		}

		if g.player.gotAchievement9 && g.achievementPage == 3 {
			screen.DrawImage(achBar2, op3)

			op3.GeoM.Reset()
			op3.GeoM.Translate(100.0, 175.0)
			screen.DrawImage(ach9, op3)
		}

		// Draw locked achievements
		if !g.player.gotAchievement1 && g.achievementPage == 1 {
			screen.DrawImage(achBarLocked, op1)
		}

		if !g.player.gotAchievement2 && g.achievementPage == 1 {
			screen.DrawImage(achBarLocked, op2)
		}

		if !g.player.gotAchievement3 && g.achievementPage == 1 {
			screen.DrawImage(achBarLocked, op3)
		}

		if !g.player.gotAchievement4 && g.achievementPage == 2 {
			screen.DrawImage(achBarLocked, op1)
		}

		if !g.player.gotAchievement5 && g.achievementPage == 2 {
			screen.DrawImage(achBarLocked, op2)
		}

		if !g.player.gotAchievement6 && g.achievementPage == 2 {
			screen.DrawImage(achBarLocked, op3)
		}

		if !g.player.gotAchievement7 && g.achievementPage == 3 {
			screen.DrawImage(achBarLocked, op1)
		}

		if !g.player.gotAchievement8 && g.achievementPage == 3 {
			screen.DrawImage(achBarLocked, op2)
		}

		if !g.player.gotAchievement9 && g.achievementPage == 3 {
			screen.DrawImage(achBarLocked, op3)
		}

		g.op.GeoM.Reset()
		g.op.GeoM.Translate(270.0, 120.0)
		if g.achievementPage < 3 {
			screen.DrawImage(rightArrow, &g.op)
		}

		g.op.GeoM.Reset()
		g.op.GeoM.Translate(20.0, 120.0)
		if g.achievementPage > 1 {
			screen.DrawImage(leftArrow, &g.op)
		}

		g.op.GeoM.Reset()
		g.op.GeoM.Translate(20.0, 10.0)
		screen.DrawImage(leftArrow, &g.op)
	}

	// Show dialogue box when asked for
	if g.dialogueShowing {
		g.op.GeoM.Reset()
		g.op.GeoM.Translate(0.0, 0.0)
		screen.DrawImage(dialogueBox, &g.op)

		g.op.GeoM.Reset()
		g.op.GeoM.Translate(250.0, 180.0)
		screen.DrawImage(enterKey, &g.op)

		// Show certain dialogues
		if g.scene == "STAGE1" {
			if g.amountOfDialoguesInScene == 5 {
				if g.whatDialogue == 0 {
					text.Draw(screen, "Hello there", pixeledFont, 80, 220, color.White)
				}

				if g.whatDialogue == 1 {
					text.Draw(screen, "If you don't know, my name is Mr. Evil", pixeledFontSmall, 80, 220, color.White)
				}

				if g.whatDialogue == 2 {
					text.Draw(screen, "Welcome to my dungeon", pixeledFontSmall, 80, 220, color.White)
				}

				if g.whatDialogue == 3 {
					text.Draw(screen, "To escape this room, find the shrine", pixeledFontSmall, 80, 220, color.White)
				}

				if g.whatDialogue == 4 {
					text.Draw(screen, "Now where could THAT be? Right?", pixeledFontSmall, 80, 220, color.White)
				}

				if g.whatDialogue == 5 {
					text.Draw(screen, "I wish you luck, subject.", pixeledFontSmall, 80, 220, color.White)
				}
			}

			if g.amountOfDialoguesInScene == 2 {
				if g.whatDialogue == 0 {
					text.Draw(screen, "Looks like you found what you needed!", pixeledFontSmall, 80, 220, color.White)
				}

				if g.whatDialogue == 1 {
					text.Draw(screen, "Er... well done", pixeledFontSmall, 80, 220, color.White)
				}

				if g.whatDialogue == 2 {
					text.Draw(screen, "Now find the exit and move on.", pixeledFontSmall, 80, 220, color.White)
				}
			}
		}

		if g.scene == "STAGE2" {
			if g.whatDialogue == 0 {
				text.Draw(screen, "Good job.", pixeledFontSmall, 80, 220, color.White)
			}

			if g.whatDialogue == 1 {
				text.Draw(screen, "But this is only the beginning.", pixeledFontSmall, 80, 220, color.White)
			}

			if g.whatDialogue == 2 {
				text.Draw(screen, "Here's the trick:", pixeledFontSmall, 80, 220, color.White)
			}

			if g.whatDialogue == 3 {
				text.Draw(screen, `"Watch your surroundings ...`, pixeledFontSmall, 80, 220, color.White)
			}

			if g.whatDialogue == 4 {
				text.Draw(screen, `... for everything looks ...`, pixeledFontSmall, 80, 220, color.White)
			}

			if g.whatDialogue == 5 {
				text.Draw(screen, `... all so similar."`, pixeledFontSmall, 80, 220, color.White)
			}

			if g.whatDialogue == 6 {
				text.Draw(screen, "You know what to do.", pixeledFontSmall, 80, 220, color.White)
			}

			if g.whatDialogue == 7 {
				text.Draw(screen, "Good luck.", pixeledFontSmall, 80, 220, color.White)
			}
		}

		if g.scene == "STAGE3" {
			if g.whatDialogue == 0 {
				text.Draw(screen, "Smarter than I thought...", pixeledFontSmall, 80, 220, color.White)
			}

			if g.whatDialogue == 1 {
				text.Draw(screen, "Let's see if you get over this.", pixeledFontSmall, 80, 220, color.White)
			}

			if g.whatDialogue == 2 {
				text.Draw(screen, "The arrows...", pixeledFontSmall, 80, 220, color.White)
			}

			if g.whatDialogue == 3 {
				text.Draw(screen, "They lead to the shrine.", pixeledFontSmall, 80, 220, color.White)
			}

			if g.whatDialogue == 4 {
				text.Draw(screen, "... or perhaps not.", pixeledFontSmall, 80, 220, color.White)
			}

			if g.whatDialogue == 5 {
				text.Draw(screen, "Find it out.", pixeledFontSmall, 80, 220, color.White)
			}
		}

		if g.scene == "STAGE4" {
			if g.whatDialogue == 0 {
				text.Draw(screen, "You made it this far", pixeledFontSmall, 80, 220, color.White)
			}

			if g.whatDialogue == 1 {
				text.Draw(screen, "So I gave you a break room", pixeledFontSmall, 80, 220, color.White)
			}

			if g.whatDialogue == 2 {
				text.Draw(screen, "Stay as much as you need to stay", pixeledFontSmall, 80, 220, color.White)
			}

			if g.whatDialogue == 3 {
				text.Draw(screen, "And prepare for the next riddle.", pixeledFontSmall, 80, 220, color.White)
			}
		}

		if g.scene == "STAGE5" {
			if g.whatDialogue == 0 {
				text.Draw(screen, "This is the hardest riddle.", pixeledFontSmall, 80, 220, color.White)
			}

			if g.whatDialogue == 1 {
				text.Draw(screen, "A room to get locked at.", pixeledFontSmall, 80, 220, color.White)
			}

			if g.whatDialogue == 2 {
				text.Draw(screen, "Fake walls.", pixeledFontSmall, 80, 220, color.White)
			}

			if g.whatDialogue == 3 {
				text.Draw(screen, "Arrows lead nowhere.", pixeledFontSmall, 80, 220, color.White)
			}

			if g.whatDialogue == 4 {
				text.Draw(screen, "It's all pretty dark.", pixeledFontSmall, 80, 220, color.White)
			}

			if g.whatDialogue == 5 {
				text.Draw(screen, "And you are really...", pixeledFontSmall, 80, 220, color.White)
			}

			if g.whatDialogue == 6 {
				text.Draw(screen, "... really slow.", pixeledFontSmall, 80, 220, color.White)
			}

			if g.whatDialogue == 7 {
				text.Draw(screen, "Let's see you beat this one.", pixeledFontSmall, 80, 220, color.White)
			}
		}
	}

	// Check if the number of dialogues per level exceeds the dialogue at the moment
	if g.whatDialogue > g.amountOfDialoguesInScene {
		g.dialogueShowing = false
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
