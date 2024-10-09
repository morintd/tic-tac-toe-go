package game

type GameModule struct {
	games   []*Game
	lobby   []IPlayer
	waiting []IPlayer
}

func (module *GameModule) OnPlayerJoin(player IPlayer) {
	if len(module.lobby)%2 > 0 {
		game := NewGame(module.lobby[0], player, func(finished *Game) {
			module.OnGameFinished(finished)
		})

		module.lobby = []IPlayer{}
		module.games = append(module.games, game)

		go game.Start()
	} else {
		module.lobby = append(module.lobby, player)
	}
}

// @TODO: link game to player to avoid finding them
func (module *GameModule) OnPlayerLeave(player IPlayer) {
	for i, p := range module.lobby {
		if p == player {
			module.lobby = append(module.lobby[:i], module.lobby[i+1:]...)
			return
		}
	}

	for i, p := range module.waiting {
		if p == player {
			module.waiting = append(module.waiting[:i], module.waiting[i+1:]...)
			return
		}
	}

	for _, game := range module.games {
		if game.PlayerX == player {
			game.Stop()

			game.PlayerO.Send(NewPlayerSentLobbyNotificationPayload("OTHER_PLAYER_DISCONNECTED"))
			module.OnPlayerJoin(game.PlayerO)

			game.Clear()

			return
		}

		if game.PlayerO == player {
			game.Stop()

			game.PlayerX.Send(NewPlayerSentLobbyNotificationPayload("OTHER_PLAYER_DISCONNECTED"))
			module.OnPlayerJoin(game.PlayerX)

			game.Clear()

			return
		}
	}
}

func (module *GameModule) OnPlayerReady(player IPlayer) {
	for i, p := range module.waiting {
		if p == player {
			module.waiting = append(module.waiting[:i], module.waiting[i+1:]...)
			player.Send(NewPlayerSentLobbyNotificationPayload("PLAYER_READY"))
			module.OnPlayerJoin(player)

			return
		}
	}
}

func (module *GameModule) OnGameFinished(finishedGame *Game) {
	for i, game := range module.games {
		if game == finishedGame {
			module.games = append(module.games[:i], module.games[i+1:]...)

			game.PlayerO.Send(NewPlayerSentWaitingNotificationPayload("GAME_FINISHED"))
			game.PlayerX.Send(NewPlayerSentWaitingNotificationPayload("GAME_FINISHED"))

			module.waiting = append(module.waiting, game.PlayerO, game.PlayerX)

			game.Clear()

			return
		}
	}
}

func NewGameModule() *GameModule {
	return &GameModule{games: []*Game{}, lobby: []IPlayer{}, waiting: []IPlayer{}}
}
