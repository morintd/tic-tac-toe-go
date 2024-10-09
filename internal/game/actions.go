package game

/* Notifications from the game to players */

type GameNotificationAction struct {
	Type    string `json:"type"`
	Payload any    `json:"payload"`
}

/* Error */

type ErrorNotificationPayload struct {
	Error string `json:"error"`
}

func NewErrorNotificationAction(err string) *GameNotificationAction {
	return &GameNotificationAction{
		Type: "ERROR",
		Payload: ErrorNotificationPayload{
			Error: err,
		},
	}
}

/* Game Start */

type GameStartNotificationPayload struct {
	Player string `json:"player"`
}

func NewGameStartNotificationAction(player string) *GameNotificationAction {
	return &GameNotificationAction{
		Type:    "GAME_START",
		Payload: GameStartNotificationPayload{Player: player},
	}
}

/* Player has moved */

type PlayerMoveNotificationPayload struct {
	Winner string   `json:"winner"`
	Board  []string `json:"board"`
	Next   string   `json:"next"`
}

func NewPlayerMoveNotificationAction(winner string, board []string, next string) *GameNotificationAction {
	return &GameNotificationAction{
		Type: "PLAYER_MOVED",
		Payload: PlayerMoveNotificationPayload{
			Winner: winner,
			Board:  board,
			Next:   next,
		},
	}
}

/* Player has been moved to lobby */

type PlayerSentLobbyNotificationPayload struct {
	Reason string `json:"reason"`
}

func NewPlayerSentLobbyNotificationPayload(reason string) *GameNotificationAction {
	return &GameNotificationAction{
		Type: "PLAYER_SENT_TO_LOBBY",
		Payload: PlayerSentLobbyNotificationPayload{
			Reason: reason,
		},
	}
}

/* Player has been moved to waiting room */

type PlayerSentWaitingNotificationPayload struct {
	Reason string `json:"reason"`
}

func NewPlayerSentWaitingNotificationPayload(reason string) *GameNotificationAction {
	return &GameNotificationAction{
		Type: "PLAYER_SENT_TO_WAITING_ROOM",
		Payload: PlayerSentLobbyNotificationPayload{
			Reason: reason,
		},
	}
}

/* Actions from the players to the game */

type PlayerAction struct {
	Type    string `json:"type"`
	Payload any    `json:"payload"`
}

/* Player moved */

var PlayerMovedType = "PLAYER_MOVE"

type PlayerMovedPayload struct {
	Square int `json:"square"`
}

var PlayerReadyType = "PLAYER_READY"

type PlayerReadyPayload struct {
}
