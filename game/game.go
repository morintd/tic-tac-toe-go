package game

import (
	boardpkg "tic-tac-toe/board"
	playerpkg "tic-tac-toe/player"
)

type Winner int

const (
	Winner_None Winner = iota
	Winner_X    Winner = iota
	Winner_O    Winner = iota
)

type Game struct {
	board  boardpkg.Board
	winner Winner
	isDraw bool
	turn   playerpkg.Player
}

func NewGame() Game {
	return Game{board: boardpkg.NewBoard(), winner: Winner_None, isDraw: false, turn: playerpkg.Player_X}
}

func (game *Game) GetBoard() *boardpkg.Board {
	return &game.board
}

func (game *Game) GetWinner() Winner {
	return game.winner
}

func (game *Game) GetIsDraw() bool {
	return game.isDraw
}

func (game *Game) GetTurn() playerpkg.Player {
	return game.turn
}

func (game *Game) calculateWinner() Winner {
	squares := game.board.GetSquares()

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

		if squares[line[0]] == squares[line[1]] && squares[line[0]] == squares[line[2]] && squares[line[0]] != boardpkg.Square_Empty {
			square := squares[line[0]]

			if square == boardpkg.Square_X {
				return Winner_X
			} else {
				return Winner_O
			}
		}
	}

	return Winner_None
}

func (game *Game) calculateIsDraw() bool {
	squares := game.board.GetSquares()

	for i := 0; i < len(squares); i++ {
		if squares[i] == boardpkg.Square_Empty {
			return false
		}
	}

	return true
}

func (game *Game) Play(index int) {
	if game.turn == playerpkg.Player_X {
		game.board.SetSquare(boardpkg.Square_X, index)
		game.turn = playerpkg.Player_O
	} else {
		game.board.SetSquare(boardpkg.Square_O, index)
		game.turn = playerpkg.Player_X
	}

	game.winner = game.calculateWinner()
	game.isDraw = game.calculateIsDraw()
}
