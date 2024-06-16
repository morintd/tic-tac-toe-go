package game

import (
	"reflect"
	"testing"
	"tic-tac-toe/board"
)

func TestShouldPlayXThenO(t *testing.T) {
	game := NewGame()

	moves := []struct {
		Move    int
		Squares [9]board.Square
	}{
		{
			Move:    0,
			Squares: [9]board.Square{board.Square_X, board.Square_Empty, board.Square_Empty, board.Square_Empty, board.Square_Empty, board.Square_Empty, board.Square_Empty, board.Square_Empty, board.Square_Empty},
		},
		{
			Move:    2,
			Squares: [9]board.Square{board.Square_X, board.Square_Empty, board.Square_O, board.Square_Empty, board.Square_Empty, board.Square_Empty, board.Square_Empty, board.Square_Empty, board.Square_Empty},
		},
		{
			Move:    1,
			Squares: [9]board.Square{board.Square_X, board.Square_X, board.Square_O, board.Square_Empty, board.Square_Empty, board.Square_Empty, board.Square_Empty, board.Square_Empty, board.Square_Empty},
		},
		{
			Move:    3,
			Squares: [9]board.Square{board.Square_X, board.Square_X, board.Square_O, board.Square_O, board.Square_Empty, board.Square_Empty, board.Square_Empty, board.Square_Empty, board.Square_Empty},
		},
		{
			Move:    5,
			Squares: [9]board.Square{board.Square_X, board.Square_X, board.Square_O, board.Square_O, board.Square_Empty, board.Square_X, board.Square_Empty, board.Square_Empty, board.Square_Empty},
		},
		{
			Move:    4,
			Squares: [9]board.Square{board.Square_X, board.Square_X, board.Square_O, board.Square_O, board.Square_O, board.Square_X, board.Square_Empty, board.Square_Empty, board.Square_Empty},
		},
		{
			Move:    6,
			Squares: [9]board.Square{board.Square_X, board.Square_X, board.Square_O, board.Square_O, board.Square_O, board.Square_X, board.Square_X, board.Square_Empty, board.Square_Empty},
		},
		{
			Move:    7,
			Squares: [9]board.Square{board.Square_X, board.Square_X, board.Square_O, board.Square_O, board.Square_O, board.Square_X, board.Square_X, board.Square_O, board.Square_Empty},
		},
		{
			Move:    8,
			Squares: [9]board.Square{board.Square_X, board.Square_X, board.Square_O, board.Square_O, board.Square_O, board.Square_X, board.Square_X, board.Square_O, board.Square_X},
		},
	}

	for _, step := range moves {
		game.Play(step.Move)
		if !reflect.DeepEqual(game.GetBoard().GetSquares(), step.Squares) {
			t.Fatalf("X and O are not being played in the right order on step %v, %v, %v", step.Move, game.GetBoard().GetSquares(), step.Squares)
		}
	}
}

func TestShouldWin(t *testing.T) {
	game := NewGame()
	moves := []int{0, 8, 1, 7, 2}

	for _, move := range moves {
		game.Play(move)
	}

	if game.GetWinner() != Winner_X {
		t.Fatalf("X should be winner, got %v instead", game.GetWinner())
	}
}

func TestShouldDraw(t *testing.T) {
	game := NewGame()
	moves := []int{0, 2, 1, 3, 5, 4, 6, 7, 8}

	for _, move := range moves {
		game.Play(move)
	}

	if game.GetIsDraw() != true {
		t.Fatalf("Game should be drawn")
	}
}
