package game

import (
	"encoding/json"
	"errors"
	"log"
	"tic-tac-toe/internal/common"
)

type GameConnectionManager struct {
	// @TODO: replace by multiple listeners on Client?
	players map[*common.Client]IPlayer
	game    *GameModule
}

func (manager *GameConnectionManager) OnNewConnection(client *common.Client) {
	player := NewWebsocketPlayer(client)

	manager.players[client] = player
	manager.game.OnPlayerJoin(player)
}

func (manager *GameConnectionManager) OnDisconnected(client *common.Client) {
	manager.game.OnPlayerLeave(manager.players[client])
}

func (manager *GameConnectionManager) OnMessage(client *common.Client, message []byte) {
	var action TypedAction

	err := json.Unmarshal(message, &action)

	if err != nil {
		log.Println(err)
		return
	}

	if action.Type == PlayerReadyType {
		manager.game.OnPlayerReady(manager.players[client])
		return
	}

	payload, err := manager.getPayload(action.Type, message)

	if err != nil {
		log.Println(err)
	} else {
		manager.players[client].Receiver() <- PlayerAction{
			Type:    action.Type,
			Payload: payload,
		}
	}
}

func (manage *GameConnectionManager) getPayload(actionType string, action []byte) (any, error) {
	if actionType == PlayerMovedType {
		var payload PayloadAction[PlayerMovedPayload]

		err := json.Unmarshal(action, &payload)

		if err != nil {
			return nil, err
		}

		return payload.Payload, nil
	}

	return nil, errors.New("invalid action type")
}

func NewGameConnectionManager(game *GameModule) *GameConnectionManager {
	return &GameConnectionManager{players: make(map[*common.Client]IPlayer), game: game}
}

type WebsocketPlayer struct {
	client   *common.Client
	receiver chan PlayerAction
}

func (player *WebsocketPlayer) Send(action *GameNotificationAction) {
	message, err := json.Marshal(action)

	if err != nil {
		log.Fatal(err.Error())
	}

	player.client.Send([]byte(message))
}

func (player *WebsocketPlayer) Receiver() chan PlayerAction {
	return player.receiver
}

func NewWebsocketPlayer(client *common.Client) IPlayer {
	return &WebsocketPlayer{client: client, receiver: make(chan PlayerAction)}
}

type TypedAction struct {
	Type string `json:"type"`
}

type PayloadAction[T any] struct {
	Payload T `json:"payload"`
}
