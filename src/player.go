package main

import "github.com/hajimehoshi/ebiten/v2"

// Player properties
type Player struct {
	x                 float64
	y                 float64
	w                 int
	h                 int
	isMoving          bool
	idleNumber        int
	walkNumber        int
	enabledShrine1    bool
	enabledShrine2    bool
	enabledShrine3    bool
	enabledLastShrine bool
	movedOnce         bool
}

// Draw player
func (p *Player) Draw(screen *ebiten.Image) {
	s := charIdle1

	switch {
	// Idle
	case !p.isMoving && p.idleNumber >= 0 && p.idleNumber <= 10:
		s = charIdle1
	case !p.isMoving && p.idleNumber >= 11 && p.idleNumber <= 20:
		s = charIdle2
	case !p.isMoving && p.idleNumber >= 21 && p.idleNumber <= 30:
		s = charIdle3

	// Walk
	case p.isMoving && p.walkNumber >= 0 && p.walkNumber <= 20:
		s = charWalk1
	case p.isMoving && p.walkNumber >= 21 && p.walkNumber <= 40:
		s = charWalk2
	case p.isMoving && p.walkNumber >= 41 && p.walkNumber <= 60:
		s = charWalk1
	case p.isMoving && p.walkNumber >= 61 && p.walkNumber <= 80:
		s = charWalk3
	}

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Reset()
	op.GeoM.Translate(-float64(p.w)/2, -float64(p.h)/2)
	op.GeoM.Translate(float64(p.x), float64(p.y))
	screen.DrawImage(s, op)
}
