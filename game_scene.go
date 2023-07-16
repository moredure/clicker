package main

type GameScene struct {
	deadlineTimer Timer
	stats         map[string]int
	winner        string
}

func (gs *GameScene) State(username string) any {
	if gs.winner != "" {
		if gs.winner == username { // TODO tie
			return "you win"
		}
		return "you loose"
	}
	opponent := ""
	for un := range gs.stats {
		if un != username {
			opponent = un
		}
	}
	return map[string]interface{}{
		"you":      gs.stats[username],
		"opponent": gs.stats[opponent],
	}
}

func (gs *GameScene) Tap(username string) {
	gs.stats[username] += 1
}

func (gs *GameScene) ChooseWinner() {
	if gs.winner != "" {
		return
	}
	x := -1
	for username, stats := range gs.stats {
		if stats > x {
			gs.winner = username
		}
	}
}

func (gs *GameScene) Loose(username string) {
	if gs.winner != "" {
		return
	}
	for un := range gs.stats {
		if un != username {
			gs.winner = un
		}
	}
}
