package main

type GameServerSessions map[string]*GameServerSession

func (gsss GameServerSessions) Broadcast(e event) {
	for _, gss := range gsss {
		gss.emit(e)
	}
}

func (gsss GameServerSessions) Close() {
	for _, gss := range gsss {
		gss.Close()
	}
}

func (gsss GameServerSessions) Publish(scene *GameScene) {
	for _, gss := range gsss {
		gss.emit(&PlayerStateEvent{
			State: scene.Show(gss.Username),
		})
	}
}
