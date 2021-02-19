package main

import (
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth  = 320
	screenHeight = 240
)

func main() {
	rand.Seed(time.Now().UnixNano())
	loadResources()
	loadMaps()
	loadSounds()
	loadFonts()

	game := &Game{
		scene:           "menu",
		atLevel:         0,
		achievementPage: 1,
		dialogueShowing: false,
		whatDialogue:    0,
	}
	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
	ebiten.SetWindowResizable(true)
	ebiten.MaximizeWindow()
	ebiten.SetWindowTitle("Infinidungeon")
	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}
