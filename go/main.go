package main

import "fmt"

func IsValidPuzzle(puzzle Puzzle, board Board, index int) bool {
	if index == 0 {
		return true
	}

	currentSquare := board.Squares[index]

	if currentSquare.Row != 0 {
		topSquare := board.Squares[index-board.Diagonal]

		if topSquare.Puzzle.Bottom()+puzzle.Top() != 0 {
			return false
		}
	}

	if currentSquare.Col != 0 {
		leftSquare := board.Squares[index-1]

		if leftSquare.Puzzle.Right()+puzzle.Left() != 0 {
			return false
		}
	}

	return true
}

func PlacePuzzle(puzzles []Puzzle, board Board, index int, archived []Puzzle, counter int) (int, error) {
	counter ++

	if len(puzzles) <= 0 {
		return counter, nil
	}

	puzzle := puzzles[0]

	if index == 0 {
		if puzzle.AtFirstSquare >= 4 {
			return counter, nil
		}

		puzzle.AtFirstSquare++
	}

	isValidPuzzle := IsValidPuzzle(puzzle, board, index)

	if isValidPuzzle && puzzle.Rotation < 4 && puzzle.Active {
		board.AddPuzzle(puzzle, index)
		archived = append(archived, puzzles[0])
		return PlacePuzzle(puzzles[1:], board, index+1, archived, counter)
	}

	if !isValidPuzzle && puzzle.Rotation < 4 && puzzle.Active {
		puzzles[0].Rotate()
		return PlacePuzzle(puzzles, board, index, archived, counter)
	}

	if len(puzzles) > 1 && puzzles[1].Active {
		puzzles[0].Reset()
		puzzles[0].Active = false
		puzzles = append(puzzles[1:], puzzles[0])

		return PlacePuzzle(puzzles, board, index, archived, counter)
	}

	if len(archived) > 0 {
		prePuzzle := archived[len(archived)-1]
		archived = archived[:len(archived)-1]

		for i, _ := range puzzles {
			puzzles[i].Reset()
		}

		prePuzzle.Rotate()
		puzzles = append([]Puzzle{prePuzzle}, puzzles...)
		return PlacePuzzle(puzzles, board, index-1, archived, counter)
	}

	return counter, nil
}

func main() {
	board := NewBoard(3)
	var archived []Puzzle
	puzzles := []Puzzle{
		NewPuzzle(-3, -4, 2, -1),
		NewPuzzle(2, -1, 3, 2),
		NewPuzzle(-1, -2, 4, 1),
		//NewPuzzle(-4, 4, -3, 1),
		//NewPuzzle(2, 3, -2, 4),
		//NewPuzzle(2, -1, 3, -1),
		//NewPuzzle(-4, 1, 2, -3),
		//NewPuzzle(-4, -4, -3, 2),
		//NewPuzzle(3, 4, 2, -4),
	}

	totalSteps, _ := PlacePuzzle(puzzles, board, 0, archived, 0)

	fmt.Println(totalSteps)
	fmt.Println(board.Squares)
}
