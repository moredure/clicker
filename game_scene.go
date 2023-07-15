package main

type GameScene struct {
	deadlineTimer Timer
	stats         map[string]int
	winner        string
}

func (gs *GameScene) Show(username string) event {
	return nil
}

func (gs *GameScene) Tap(username string) {
	gs.stats[username] += 1
}

func (gs *GameScene) Finish() {
}

func (gs *GameScene) Loose(username string) {
}
