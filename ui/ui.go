package ui

import (
	"fmt"
	boardpkg "tic-tac-toe/board"
	gamepkg "tic-tac-toe/game"
	playerpkg "tic-tac-toe/player"
)

type UI interface {
	RenderGrid(board [9]boardpkg.Square)
	RequestMove(player playerpkg.Player) int
	AlertWrongInput()
	AlertSquareTaken(square int)
	AnnounceWinner(winner gamepkg.Winner)
	AnnounceDraw()
}

type ConsoleUI struct {
}

func (consoleUI ConsoleUI) RenderGrid(board [9]boardpkg.Square) {
	fmt.Println("-------")
	fmt.Println("|" + consoleUI.getSquare(board[0]) + "|" + consoleUI.getSquare(board[1]) + "|" + consoleUI.getSquare(board[2]) + "|")
	fmt.Println("|" + consoleUI.getSquare(board[3]) + "|" + consoleUI.getSquare(board[4]) + "|" + consoleUI.getSquare(board[5]) + "|")
	fmt.Println("|" + consoleUI.getSquare(board[6]) + "|" + consoleUI.getSquare(board[7]) + "|" + consoleUI.getSquare(board[8]) + "|")
	fmt.Println("-------")
}

func (ConsoleUI) RequestMove(player playerpkg.Player) int {
	var input int

	if player == playerpkg.Player_X {
		fmt.Println("X, what square do you play?")
	} else {
		fmt.Println("O, what square do you play?")
	}

	fmt.Scanln(&input)

	return input
}

func (ConsoleUI) AlertWrongInput() {
	fmt.Println("Input should be between 0 and 8")
}

func (ConsoleUI) AlertSquareTaken(square int) {
	fmt.Println("Square already taken (" + fmt.Sprint(square) + ")")
}

func (ConsoleUI) AnnounceWinner(winner gamepkg.Winner) {
	if winner == gamepkg.Winner_X {
		fmt.Println("Winner : X")
	} else {
		fmt.Println("Winner : O")
	}
}

func (ConsoleUI) AnnounceDraw() {
	fmt.Println("It's a draw!")
}

func (ConsoleUI) getSquare(square boardpkg.Square) string {
	if square == boardpkg.Square_Empty {
		return " "
	}

	if square == boardpkg.Square_X {
		return "X"
	}

	if square == boardpkg.Square_O {
		return "O"
	}

	return ""
}
