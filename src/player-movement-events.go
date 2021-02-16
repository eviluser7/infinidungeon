package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

// Update player state
func (p *Player) Update(g *Game) {

	// If player is not moving, increase idle animation
	if !p.isMoving {
		p.idleNumber++
	}

	// Update player position & increase animation number
	if ebiten.IsKeyPressed(ebiten.KeyW) && p.y > 92 && !p.talkedToEvil {
		if g.scene != "STAGE5" {
			p.y -= 2.0
		} else if g.scene == "STAGE5" {
			p.y--
		}
	}
	p.isMoving = true
	p.walkNumber++

	if ebiten.IsKeyPressed(ebiten.KeyA) && p.x > 28 && !p.talkedToEvil {
		if g.scene != "STAGE5" {
			p.x -= 2.0
		} else if g.scene == "STAGE5" {
			p.x--
		}
	}
	p.isMoving = true
	p.walkNumber++

	if ebiten.IsKeyPressed(ebiten.KeyS) && p.y < 207 && !p.talkedToEvil {
		if g.scene != "STAGE5" {
			p.y += 2.0
		} else if g.scene == "STAGE5" {
			p.y++
		}
	}
	p.isMoving = true
	p.walkNumber++

	if ebiten.IsKeyPressed(ebiten.KeyD) && p.x < 292 && !p.talkedToEvil {
		if g.scene != "STAGE5" {
			p.x += 2.0
		} else if g.scene == "STAGE5" {
			p.x++
		}
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

	if p.isMoving {
		p.movedOnce = true
	}
}
