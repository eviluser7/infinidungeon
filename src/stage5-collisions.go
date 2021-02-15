package main

// UpdateMap05 checks collisions on stage 5
func (p *Player) UpdateMap05(g *Game) {
	if g.scene == "STAGE5" {
		if p.x >= 290 && g.atLevel == 0 {
			p.x = 45
			g.atLevel = 1
		} else if p.x <= 30 && g.atLevel == 1 {
			p.x = 264
			g.atLevel = 0
		} else if p.y >= 207 && g.atLevel == 0 {
			p.y = 120
			g.atLevel = 7
		} else if p.y <= 92 && g.atLevel == 7 {
			p.y = 160
			g.atLevel = 0
		}

		if p.x >= 290 && g.atLevel == 1 {
			p.x = 45
			g.atLevel = 2
		} else if p.x <= 30 && g.atLevel == 2 {
			p.x = 264
			g.atLevel = 1
		} else if p.y >= 207 && g.atLevel == 2 {
			p.y = 120
			g.atLevel = 9
		} else if p.y <= 92 && g.atLevel == 9 {
			p.y = 160
			g.atLevel = 2
		} else if p.x <= 30 && g.atLevel == 9 {
			p.x = 264
			g.atLevel = 8
		} else if p.x >= 290 && g.atLevel == 8 {
			p.x = 45
			g.atLevel = 9
		}

		if p.x >= 290 && g.atLevel == 9 {
			p.x = 45
			g.atLevel = 10
		} else if p.x <= 30 && g.atLevel == 10 {
			p.x = 264
			g.atLevel = 9
		} else if p.y <= 92 && g.atLevel == 10 {
			p.y = 160
			g.atLevel = 3
		} else if p.y >= 207 && g.atLevel == 3 {
			p.y = 120
			g.atLevel = 10
		}

		if p.x >= 290 && g.atLevel == 3 {
			p.x = 45
			g.atLevel = 4
		} else if p.x <= 30 && g.atLevel == 4 {
			p.x = 264
			g.atLevel = 3
		}

		if p.x >= 290 && g.atLevel == 2 {
			p.x = 45
			g.atLevel = 3
		} else if p.x <= 30 && g.atLevel == 3 {
			p.x = 264
			g.atLevel = 2
		}

		if p.x >= 290 && g.atLevel == 4 {
			p.x = 45
			g.atLevel = 5
		} else if p.x <= 30 && g.atLevel == 5 {
			p.x = 264
			g.atLevel = 4
		} else if p.x >= 290 && g.atLevel == 5 {
			p.x = 45
			g.atLevel = 6
		} else if p.x <= 30 && g.atLevel == 6 {
			p.x = 264
			g.atLevel = 5
		}

		if p.y >= 207 && g.atLevel == 5 {
			p.y = 120
			g.atLevel = 12
		} else if p.y <= 92 && g.atLevel == 12 {
			p.y = 160
			g.atLevel = 5
		}

		if p.y >= 207 && g.atLevel == 6 {
			p.y = 120
			g.atLevel = 13
		} else if p.y <= 92 && g.atLevel == 13 {
			p.y = 160
			g.atLevel = 6
		} else if p.y >= 207 && g.atLevel == 13 {
			p.y = 120
			g.atLevel = 20
		} else if p.y <= 92 && g.atLevel == 20 {
			p.y = 160
			g.atLevel = 13
		}

		if p.x >= 290 && g.atLevel == 10 {
			p.x = 45
			g.atLevel = 11
		} else if p.x <= 30 && g.atLevel == 11 {
			p.x = 264
			g.atLevel = 10
		} else if p.y >= 207 && g.atLevel == 11 {
			p.y = 120
			g.atLevel = 18
		} else if p.y <= 92 && g.atLevel == 18 {
			p.y = 160
			g.atLevel = 11
		}

		if p.x >= 290 && g.atLevel == 11 {
			p.x = 45
			g.atLevel = 12
		} else if p.x <= 30 && g.atLevel == 12 {
			p.x = 264
			g.atLevel = 11
		} else if p.x >= 290 && g.atLevel == 12 {
			p.x = 45
			g.atLevel = 13
		} else if p.x <= 30 && g.atLevel == 13 {
			p.x = 264
			g.atLevel = 12
		}

		if p.x >= 290 && g.atLevel == 14 {
			p.x = 45
			g.atLevel = 15
		} else if p.x <= 30 && g.atLevel == 15 {
			p.x = 264
			g.atLevel = 14
		} else if p.y >= 207 && g.atLevel == 14 {
			p.y = 120
			g.atLevel = 21
		} else if p.y <= 92 && g.atLevel == 21 {
			p.y = 160
			g.atLevel = 14
		}

		if p.x >= 290 && g.atLevel == 15 {
			p.x = 45
			g.atLevel = 16
		} else if p.x <= 30 && g.atLevel == 16 {
			p.x = 264
			g.atLevel = 15
		} else if p.y >= 207 && g.atLevel == 15 {
			p.y = 120
			g.atLevel = 22
		} else if p.y <= 92 && g.atLevel == 22 {
			p.y = 160
			g.atLevel = 15
		}

		if p.x >= 290 && g.atLevel == 16 {
			p.x = 45
			g.atLevel = 17
		} else if p.x <= 30 && g.atLevel == 17 {
			p.x = 264
			g.atLevel = 16
		}

		if p.y <= 92 && g.atLevel == 17 {
			p.y = 160
			g.atLevel = 10
		} else if p.y >= 207 && g.atLevel == 10 {
			p.y = 120
			g.atLevel = 17
		} else if p.y >= 207 && g.atLevel == 17 {
			p.y = 120
			g.atLevel = 24
		} else if p.y <= 92 && g.atLevel == 24 {
			p.y = 160
			g.atLevel = 17
		} else if p.x >= 290 && g.atLevel == 17 {
			p.x = 45
			g.atLevel = 18
		} else if p.x <= 30 && g.atLevel == 18 {
			p.x = 264
			g.atLevel = 17
		}

		if p.x >= 290 && g.atLevel == 18 {
			p.x = 45
			g.atLevel = 19
		} else if p.x <= 30 && g.atLevel == 19 {
			p.x = 264
			g.atLevel = 18
		}

		if p.x >= 290 && g.atLevel == 19 {
			p.x = 45
			g.atLevel = 20
		} else if p.x <= 30 && g.atLevel == 20 {
			p.x = 264
			g.atLevel = 19
		} else if p.y >= 207 && g.atLevel == 20 {
			p.y = 120
			g.atLevel = 27
		} else if p.y <= 92 && g.atLevel == 27 {
			p.y = 160
			g.atLevel = 20
		}

		if p.y >= 207 && g.atLevel == 22 {
			p.y = 120
			g.atLevel = 29
		} else if p.y <= 92 && g.atLevel == 29 {
			p.y = 160
			g.atLevel = 22
		}

		if p.x >= 290 && g.atLevel == 23 {
			p.x = 45
			g.atLevel = 24
		}

		if p.x <= 30 && g.atLevel == 24 {
			p.x = 264
			g.atLevel = 23
		} else if p.x >= 290 && g.atLevel == 24 {
			p.x = 45
			g.atLevel = 25
		} else if p.x <= 30 && g.atLevel == 25 {
			p.x = 264
			g.atLevel = 24
		} else if p.y >= 207 && g.atLevel == 24 {
			p.y = 120
			g.atLevel = 31
		} else if p.y <= 92 && g.atLevel == 31 {
			p.y = 160
			g.atLevel = 24
		}

		if p.x >= 290 && g.atLevel == 25 {
			p.x = 45
			g.atLevel = 26
		} else if p.x <= 30 && g.atLevel == 26 {
			p.x = 264
			g.atLevel = 25
		}

		if p.x >= 290 && g.atLevel == 28 {
			p.x = 45
			g.atLevel = 29
		} else if p.x <= 30 && g.atLevel == 29 {
			p.x = 264
			g.atLevel = 28
		} else if p.y >= 207 && g.atLevel == 28 {
			p.y = 120
			g.atLevel = 35
		} else if p.y <= 92 && g.atLevel == 35 {
			p.y = 160
			g.atLevel = 28
		}

		if p.x >= 290 && g.atLevel == 30 {
			p.x = 45
			g.atLevel = 31
		}

		if p.x <= 30 && g.atLevel == 31 {
			p.x = 264
			g.atLevel = 30
		} else if p.x >= 290 && g.atLevel == 31 {
			p.x = 45
			g.atLevel = 32
		} else if p.x <= 30 && g.atLevel == 32 {
			p.x = 264
			g.atLevel = 31
		} else if p.y >= 207 && g.atLevel == 31 {
			p.y = 120
			g.atLevel = 38
		} else if p.y <= 92 && g.atLevel == 38 {
			p.y = 160
			g.atLevel = 31
		}

		if p.y >= 207 && g.atLevel == 33 {
			p.y = 120
			g.atLevel = 40
		}

		if p.y <= 92 && g.atLevel == 34 {
			p.y = 160
			g.atLevel = 27
		}

		if p.y >= 207 && g.atLevel == 35 {
			p.y = 120
			g.atLevel = 42
		} else if p.y <= 92 && g.atLevel == 42 {
			p.y = 160
			g.atLevel = 35
		}

		if p.x >= 290 && g.atLevel == 35 {
			p.x = 45
			g.atLevel = 36
		} else if p.x <= 30 && g.atLevel == 36 {
			p.x = 264
			g.atLevel = 35
		} else if p.x >= 290 && g.atLevel == 36 {
			p.x = 45
			g.atLevel = 37
		}

		if p.x <= 30 && g.atLevel == 37 {
			p.x = 264
			g.atLevel = 36
		} else if p.y >= 207 && g.atLevel == 37 {
			p.y = 120
			g.atLevel = 44
		} else if p.y <= 92 && g.atLevel == 44 {
			p.y = 160
			g.atLevel = 37
		}

		if p.x <= 30 && g.atLevel == 44 {
			p.x = 264
			g.atLevel = 43
		} else if p.x >= 290 && g.atLevel == 43 {
			p.x = 45
			g.atLevel = 44
		}

		if p.x >= 290 && g.atLevel == 38 {
			p.x = 45
			g.atLevel = 39
		} else if p.x <= 30 && g.atLevel == 39 {
			p.x = 264
			g.atLevel = 38
		}

		if p.y >= 207 && g.atLevel == 38 {
			p.y = 120
			g.atLevel = 45
		} else if p.y <= 92 && g.atLevel == 45 && !g.player.enabledLastShrine {
			p.y = 160
			g.atLevel = 38
		}

		if p.x >= 290 && g.atLevel == 39 {
			p.x = 45
			g.atLevel = 40
		} else if p.x <= 30 && g.atLevel == 40 {
			p.x = 264
			g.atLevel = 39
		} else if p.x >= 290 && g.atLevel == 40 {
			p.x = 45
			g.atLevel = 41
		} else if p.x <= 30 && g.atLevel == 41 {
			p.x = 264
			g.atLevel = 40
		}

		if p.y <= 92 && g.atLevel == 40 {
			p.y = 160
			g.atLevel = 33
		} else if p.y >= 207 && g.atLevel == 40 {
			p.y = 120
			g.atLevel = 47
		}

		if p.y <= 92 && g.atLevel == 41 {
			p.y = 160
			g.atLevel = 34
		}

		if p.x >= 290 && g.atLevel == 46 {
			p.x = 45
			g.atLevel = 47
		} else if p.x <= 30 && g.atLevel == 48 {
			p.x = 264
			g.atLevel = 47
		}

		if p.x <= 30 && g.atLevel == 47 {
			p.x = 264
			g.atLevel = 46
		} else if p.x >= 290 && g.atLevel == 47 {
			p.x = 45
			g.atLevel = 48
		}

		if p.y >= 207 && g.atLevel == 45 && g.player.enabledLastShrine {
			g.scene = "ENDSTAGE"
			p.y = 120
		}
	}
}
