package main

type Puzzle struct {
	Origin        []int
	Edges         []int
	Rotation      int
	AtFirstSquare int
	Active        bool
}

func (puzzle *Puzzle) Left() int {
	return puzzle.Edges[0]
}

func (puzzle *Puzzle) Top() int {
	return puzzle.Edges[1]
}

func (puzzle *Puzzle) Right() int {
	return puzzle.Edges[2]
}

func (puzzle *Puzzle) Bottom() int {
	return puzzle.Edges[3]
}

func (puzzle *Puzzle) Rotate() {
	puzzle.Edges = append(puzzle.Edges[1:], puzzle.Edges[0])
	puzzle.Rotation ++
}

func (puzzle *Puzzle) Reset() {
	copy(puzzle.Edges, puzzle.Origin)
	puzzle.Active = true
	puzzle.Rotation = 0
}

func NewPuzzle(left, top, right, bottom int) Puzzle {
	return Puzzle{
		Origin:        []int{left, top, right, bottom},
		Edges:         []int{left, top, right, bottom},
		Rotation:      0,
		AtFirstSquare: 0,
		Active:        true,
	}
}
