package game

import "log"

type PlayerName string

const (
	PlayerNameX PlayerName = "X"
	PlayerNameO PlayerName = "O"
)

type Game struct {
	PlayerX IPlayer
	PlayerO IPlayer

	Board []string
	Next  PlayerName

	quit       chan int
	onFinished func(game *Game)
}

// @TODO: it's possible receivers contain actions from before the game started, should empty it
func (game *Game) Start() {
	game.PlayerX.Send(NewGameStartNotificationAction(string(PlayerNameX)))
	game.PlayerO.Send(NewGameStartNotificationAction(string(PlayerNameO)))

	for {
		select {
		case action := <-game.PlayerO.Receiver():
			if finished := game.HandlePlayerAction(PlayerNameO, action); finished {
				game.onFinished(game)
				return
			}
		case action := <-game.PlayerX.Receiver():
			if finished := game.HandlePlayerAction(PlayerNameX, action); finished {
				game.onFinished(game)
				return
			}
		case <-game.quit:
			return
		}
	}
}

func (game *Game) HandlePlayerAction(player PlayerName, action PlayerAction) bool {
	if player != game.Next {
		game.getPlayerByName(player).Send(NewErrorNotificationAction("It's not your turn to play yet!"))
		return false
	}

	if action.Type == PlayerMovedType {
		if payload, ok := action.Payload.(PlayerMovedPayload); ok {
			return game.onPlayerMove(player, payload)
		} else {
			log.Println("Wrong payload!")
			log.Println(payload)
		}
	}

	return false
}

func (game *Game) onPlayerMove(player PlayerName, payload PlayerMovedPayload) bool {
	if payload.Square < 0 || payload.Square > 8 || game.Board[payload.Square] != "" {
		game.getPlayerByName(player).Send(NewErrorNotificationAction("You cannot play on this square!"))
		return false
	}

	game.Board[payload.Square] = string(player)
	winner := game.calculateWinner()

	if game.Next == PlayerNameX {
		game.Next = PlayerNameO
	} else {
		game.Next = PlayerNameX
	}

	game.PlayerO.Send(NewPlayerMoveNotificationAction(winner, game.Board, string(game.Next)))
	game.PlayerX.Send(NewPlayerMoveNotificationAction(winner, game.Board, string(game.Next)))

	return winner != ""
}

func (game *Game) getPlayerByName(name PlayerName) IPlayer {
	if name == PlayerNameX {
		return game.PlayerX
	}

	return game.PlayerO
}

func (game *Game) calculateWinner() string {
	lines := [][]int{
		{0, 1, 2},
		{3, 4, 5},
		{6, 7, 8},
		{0, 3, 6},
		{1, 4, 7},
		{2, 5, 8},
		{0, 4, 8},
		{2, 4, 6},
	}

	for i := 0; i < len(lines); i++ {
		line := lines[i]

		if game.Board[line[0]] != "" && game.Board[line[0]] == game.Board[line[1]] && game.Board[line[0]] == game.Board[line[2]] {
			return game.Board[line[0]]
		}
	}

	for i := 0; i < len(game.Board); i++ {
		if game.Board[i] == "" {
			return ""
		}
	}

	return "draw"

}

func (game *Game) Stop() {
	game.quit <- 0
}

func (game *Game) Clear() {
	game.PlayerO = nil
	game.PlayerX = nil
	game.Board = nil
	game.onFinished = nil
}

func NewGame(playerX IPlayer, playerO IPlayer, onFinished func(game *Game)) *Game {
	return &Game{
		PlayerX:    playerX,
		PlayerO:    playerO,
		Board:      []string{"", "", "", "", "", "", "", "", ""},
		Next:       PlayerNameX,
		quit:       make(chan int),
		onFinished: onFinished,
	}
}
