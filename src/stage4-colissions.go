package main

// UpdateMap04 checks for collisions on 4th stage
func (p *Player) UpdateMap04(g *Game) {
	if g.scene == "STAGE4" {
		if p.y >= 207 {
			p.y = 120
			g.atLevel = 3
			g.scene = "transition"
			g.situation = "level5"
			g.whatDialogue = 0
		}
	}
}
