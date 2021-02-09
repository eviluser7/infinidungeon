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

	//////////////// Switching to other rooms //////////////////////////
	// Level 1
	if g.scene == "STAGE1" {
		if p.y >= 207 && g.atLevel == 0 {
			p.x = 160
			p.y = 120
			g.atLevel = 12
		} else if p.y <= 92 && g.atLevel == 12 {
			p.x = 160
			p.y = 160
			g.atLevel = 0
		}

		if p.x >= 290 && g.atLevel == 12 {
			p.x = 45
			p.y = 178
			g.atLevel = 13
		} else if p.x <= 30 && g.atLevel == 13 {
			p.x = 264
			p.y = 178
			g.atLevel = 12
		} else if p.x >= 290 && g.atLevel == 13 {
			p.x = 45
			p.y = 178
			g.atLevel = 14
		} else if p.x <= 30 && g.atLevel == 14 {
			p.x = 264
			p.y = 178
			g.atLevel = 13
		} else if p.y <= 92 && g.atLevel == 13 {
			p.x = 160
			p.y = 160
			g.atLevel = 1
		} else if p.y >= 207 && g.atLevel == 1 {
			p.x = 160
			p.y = 120
			g.atLevel = 13
		}

		if p.x >= 290 && g.atLevel == 1 {
			p.x = 45
			p.y = 152
			g.atLevel = 2
		} else if p.x <= 30 && g.atLevel == 2 {
			p.x = 264
			p.y = 152
			g.atLevel = 1
		}

		if p.x >= 290 && g.atLevel == 2 {
			p.x = 45
			p.y = 152
			g.atLevel = 3
		} else if p.x <= 30 && g.atLevel == 3 {
			p.x = 264
			p.y = 152
			g.atLevel = 2
		}

		if p.x >= 290 && g.atLevel == 3 {
			p.x = 45
			p.y = 152
			g.atLevel = 4
		} else if p.x <= 30 && g.atLevel == 4 {
			p.x = 264
			p.y = 152
			g.atLevel = 3
		} else if p.y >= 207 && g.atLevel == 3 {
			p.x = 160
			p.y = 120
			g.atLevel = 15
		} else if p.y <= 92 && g.atLevel == 15 {
			p.x = 160
			p.y = 160
			g.atLevel = 3
		} else if p.x >= 290 && g.atLevel == 15 {
			p.x = 45
			p.y = 178
			g.atLevel = 16
		} else if p.x <= 30 && g.atLevel == 16 {
			p.x = 264
			p.y = 178
			g.atLevel = 15
		}

		if p.x >= 290 && g.atLevel == 4 {
			p.x = 45
			p.y = 152
			g.atLevel = 5
		} else if p.x <= 30 && g.atLevel == 5 {
			p.x = 264
			p.y = 152
			g.atLevel = 4
		} else if p.y <= 92 && g.atLevel == 16 {
			p.x = 160
			p.y = 160
			g.atLevel = 4
		} else if p.y >= 207 && g.atLevel == 4 {
			p.x = 160
			p.y = 120
			g.atLevel = 16
		}

		if p.x >= 290 && g.atLevel == 5 {
			p.x = 45
			p.y = 152
			g.atLevel = 6
		} else if p.x <= 30 && g.atLevel == 6 {
			p.x = 264
			p.y = 152
			g.atLevel = 5
		} else if p.y <= 92 && g.atLevel == 17 {
			p.x = 160
			p.y = 160
			g.atLevel = 5
		} else if p.y >= 207 && g.atLevel == 5 {
			p.x = 160
			p.y = 120
			g.atLevel = 17
		} else if p.x >= 290 && g.atLevel == 16 {
			p.x = 45
			p.y = 178
			g.atLevel = 17
		} else if p.x <= 30 && g.atLevel == 17 {
			p.x = 264
			p.y = 178
			g.atLevel = 16
		}

		if p.y >= 207 && g.atLevel == 6 {
			p.x = 160
			p.y = 120
			g.atLevel = 18
		} else if p.y <= 92 && g.atLevel == 18 {
			p.x = 160
			p.y = 160
			g.atLevel = 6
		}

		if p.x >= 290 && g.atLevel == 6 {
			p.x = 45
			p.y = 152
			g.atLevel = 7
		} else if p.x <= 30 && g.atLevel == 7 {
			p.x = 264
			p.y = 152
			g.atLevel = 6
		}

		if p.y >= 207 && g.atLevel == 7 {
			p.x = 160
			p.y = 120
			g.atLevel = 19
		} else if p.y <= 92 && g.atLevel == 19 {
			p.x = 160
			p.y = 160
			g.atLevel = 7
		}

		if p.x >= 290 && g.atLevel == 19 {
			p.x = 45
			p.y = 178
			g.atLevel = 20
		} else if p.x <= 30 && g.atLevel == 20 {
			p.x = 264
			p.y = 178
			g.atLevel = 19
		}

		if p.x >= 290 && g.atLevel == 20 {
			p.x = 45
			p.y = 178
			g.atLevel = 21
		} else if p.x <= 30 && g.atLevel == 21 {
			p.x = 264
			p.y = 178
			g.atLevel = 20
		} else if p.y <= 92 && g.atLevel == 20 {
			p.x = 160
			p.y = 160
			g.atLevel = 8
		} else if p.y >= 207 && g.atLevel == 8 {
			p.x = 160
			p.y = 120
			g.atLevel = 20
		}

		if p.x >= 290 && g.atLevel == 21 {
			p.x = 45
			p.y = 178
			g.atLevel = 22
		} else if p.x <= 30 && g.atLevel == 22 {
			p.x = 264
			p.y = 178
			g.atLevel = 21
		} else if p.y <= 92 && g.atLevel == 21 {
			p.x = 160
			p.y = 160
			g.atLevel = 9
		} else if p.y >= 207 && g.atLevel == 9 {
			p.x = 160
			p.y = 120
			g.atLevel = 21
		}

		if p.x >= 290 && g.atLevel == 9 {
			p.x = 45
			p.y = 152
			g.atLevel = 10
		} else if p.x <= 30 && g.atLevel == 10 {
			p.x = 264
			p.y = 152
			g.atLevel = 9
		}

		if p.y <= 92 && g.atLevel == 22 {
			p.x = 160
			p.y = 160
			g.atLevel = 11
		} else if p.y >= 207 && g.atLevel == 11 {
			p.x = 160
			p.y = 120
			g.atLevel = 22
		}
	}
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

	if inpututil.IsKeyJustPressed(ebiten.KeyM) {
		g.atLevel++
		//fmt.Println(g.atLevel)
	} else if inpututil.IsKeyJustPressed(ebiten.KeyN) {
		g.atLevel--
		//fmt.Println(g.atLevel)
	}

	// Debug stages
	if inpututil.IsKeyJustPressed(ebiten.KeyG) {
		g.atLevel += 11
	} else if inpututil.IsKeyJustPressed(ebiten.KeyH) {
		g.atLevel -= 11
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
	}

	// Update player
	g.player.Update(g)

	return nil
}
