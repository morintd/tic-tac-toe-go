package controller

import (
	gamepkg "tic-tac-toe/game"
	uipkg "tic-tac-toe/ui"
)

type Controller struct {
	ui uipkg.UI
}

func NewController(ui uipkg.UI) Controller {
	return Controller{ui}
}

func (controller *Controller) Start() {
	game := gamepkg.NewGame()
	board := game.GetBoard()

	controller.ui.RenderGrid(board.GetSquares())

	for game.GetWinner() == gamepkg.Winner_None && !game.GetIsDraw() {
		input := controller.ui.RequestMove(game.GetTurn())

		if input > 8 {
			controller.ui.AlertWrongInput()
			continue
		}

		if !board.IsSquareEmpty(input) {
			controller.ui.AlertSquareTaken(input)
			continue
		}

		game.Play(input)

		controller.ui.RenderGrid(board.GetSquares())
	}

	if game.GetWinner() != gamepkg.Winner_None {
		controller.ui.AnnounceWinner(game.GetWinner())
	} else {
		controller.ui.AnnounceDraw()
	}
}
