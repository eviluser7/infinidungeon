package main

// UpdateMap06 checks collisions on the last map
func (p *Player) UpdateMap06(g *Game) {
	if g.scene == "ENDSTAGE" {
		if p.y >= 207 && g.atLevel == 0 {
			p.y = 120
			g.atLevel = 1
		}

		if p.y >= 207 && g.atLevel == 1 {
			p.y = 120
			g.atLevel = 2
		}

		if p.y >= 207 && g.atLevel == 2 {
			p.y = 120
			g.atLevel = 3
		}
	}
}
