package main

type Board struct {
	Squares  []Square
	Diagonal int
}

func (board *Board) AddPuzzle(puzzle Puzzle, index int) {
	board.Squares[index].Puzzle = puzzle
}

func NewBoard(diagonal int) Board {
	var squares []Square

	for row := 0; row < diagonal; row++ {
		for col := 0; col < diagonal; col++ {
			squares = append(squares, Square{Row: row, Col: col})
		}
	}

	return Board{Squares: squares, Diagonal: diagonal}
}
