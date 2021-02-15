package main

// UpdateMap01 checks for collisions on 1st stage
func (p *Player) UpdateMap01(g *Game) {
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
		} else if p.x >= 290 && g.atLevel == 22 {
			p.x = 45
			p.y = 178
			g.atLevel = 23
		} else if p.x <= 30 && g.atLevel == 23 {
			p.x = 264
			p.y = 178
			g.atLevel = 22
		}

		if p.y <= 92 && g.atLevel == 21 {
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

		if p.y <= 92 && g.atLevel == 23 {
			p.x = 160
			p.y = 160
			g.atLevel = 11
		} else if p.y >= 207 && g.atLevel == 11 {
			p.x = 160
			p.y = 120
			g.atLevel = 23
		}

		if p.y >= 207 && g.atLevel == 18 && g.player.enabledShrine1 {
			p.x = 160
			p.y = 120
			g.atLevel = 6
			g.scene = "STAGE2"
		}
	}
}
