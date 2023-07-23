package main

type GameScene struct {
	Timer
	stats map[string]int
	GameSceneState
}

type GameSceneState interface {
	State(username string) any
}

type Playing struct {
	GS *GameScene
}

func (p *Playing) State(username string) any {
	opponent := ""
	for un := range p.GS.stats {
		if un != username {
			opponent = un
		}
	}
	return map[string]interface{}{
		"you":      p.GS.stats[username],
		"opponent": p.GS.stats[opponent],
	}
}

type Finished struct {
	Winner string
}

func (p *Finished) State(username string) any {
	if p.Winner == "" {
		return "tie"
	}
	if p.Winner == username {
		return "you win"
	}
	return "you loose"
}

func (gs *GameScene) Tap(username string) {
	gs.stats[username] += 1
}

func (gs *GameScene) ChooseWinner() {
	x := -1
	for username, stats := range gs.stats {
		if stats > x {
			gs.GameSceneState = &Finished{Winner: username}
			return
		}
	}
}

func (gs *GameScene) Loose(username string) {
	for un := range gs.stats {
		if un != username {
			gs.GameSceneState = &Finished{Winner: un}
			return
		}
	}
}

func NewGameScene(t Timer) *GameScene {
	gs := &GameScene{
		stats: make(map[string]int, 2),
		Timer: t,
	}
	gs.GameSceneState = &Playing{gs}
	return gs
}
