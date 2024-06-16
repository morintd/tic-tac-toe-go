package board

type Square int

const (
	Square_Empty Square = iota
	Square_X     Square = iota
	Square_O     Square = iota
)

type Board struct {
	squares [9]Square
}

func (board *Board) SetSquare(square Square, index int) {
	board.squares[index] = square
}

func (board Board) GetSquares() [9]Square {
	return board.squares
}

func (board Board) IsSquareEmpty(index int) bool {
	return board.squares[index] == Square_Empty
}

func NewBoard() Board {
	return Board{squares: [9]Square{Square_Empty, Square_Empty, Square_Empty, Square_Empty, Square_Empty, Square_Empty, Square_Empty, Square_Empty, Square_Empty}}
}
