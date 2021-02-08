package main

import (
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

// Update player state
func (p *Player) Update(g *Game) {

	// If player is not moving, increase idle animation
	if !p.isMoving {
		p.idleNumber++
	}

	// Update player position & increase animation number
	if ebiten.IsKeyPressed(ebiten.KeyW) && p.y > 92 {
		p.y -= 2
	}
	p.isMoving = true
	p.walkNumber++

	if ebiten.IsKeyPressed(ebiten.KeyA) && p.x > 28 {
		p.x -= 2
	}
	p.isMoving = true
	p.walkNumber++

	if ebiten.IsKeyPressed(ebiten.KeyS) && p.y < 207 {
		p.y += 2
	}
	p.isMoving = true
	p.walkNumber++

	if ebiten.IsKeyPressed(ebiten.KeyD) && p.x < 292 {
		p.x += 2
	}
	p.isMoving = true
	p.walkNumber++

	// Check two keys at the same time
	if ebiten.IsKeyPressed(ebiten.KeyA) && ebiten.IsKeyPressed(ebiten.KeyD) {
		p.x += 0
		p.isMoving = false
	}

	if ebiten.IsKeyPressed(ebiten.KeyW) && ebiten.IsKeyPressed(ebiten.KeyS) {
		p.y += 0
		p.isMoving = false
	}

	// Check if a key is pressed
	if !ebiten.IsKeyPressed(ebiten.KeyW) &&
		!ebiten.IsKeyPressed(ebiten.KeyA) &&
		!ebiten.IsKeyPressed(ebiten.KeyS) &&
		!ebiten.IsKeyPressed(ebiten.KeyD) {
		p.x += 0
		p.y += 0
		p.isMoving = false
	}

	// If idle animation exceeds 31, return to first frame
	if p.idleNumber >= 31 {
		p.idleNumber = 0
	}

	// If walking animation exceeds 81, return to first frame
	if p.walkNumber >= 81 {
		p.walkNumber = 0
	}
}

// Update the game state
func (g *Game) Update() error {
	if !g.inited {
		g.init()
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		os.Exit(0)
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyM) {
		g.atLevel++
		//fmt.Println(g.atLevel)
	} else if inpututil.IsKeyJustPressed(ebiten.KeyN) {
		g.atLevel--
		//fmt.Println(g.atLevel)
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyG) {
		g.atLevel += 11
	} else if inpututil.IsKeyJustPressed(ebiten.KeyH) {
		g.atLevel -= 11
	}

	// Update player
	g.player.Update(g)

	return nil
}
