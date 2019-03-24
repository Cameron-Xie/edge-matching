const Board = require('./Board');
const Puzzle = require('./Puzzle');


/**
 * @return {boolean}
 */
function IsValidPuzzle(puzzle, board, index = 0) {
    if (index === 0) {
        return true;
    }

    const currentSquare = board.squares[index];

    if (currentSquare.row !== 0) {
        const topSquare = board.squares[index - board.cols];

        if (topSquare.puzzle.bottom + puzzle.top !== 0) {
            return false;
        }
    }

    if (currentSquare.col !== 0) {
        const leftSquare = board.squares[index - 1];

        if (leftSquare.puzzle.right + puzzle.left !== 0) {
            return false;
        }
    }

    return true;
}

function PlacePuzzle(puzzleList, board, index = 0, archived = [], counter = 0) {
    counter ++;
    console.log(counter + " running ...");

    // no more puzzle
    if (puzzleList.length <= 0) {
        return;
    }

    const puzzle = puzzleList[0];

    if (index === 0) {
        if (puzzle.atFirstPosition >= 4) {
            return;
        }

        puzzle.atFirstPosition++;
    }

    const isValidPuzzle = IsValidPuzzle(puzzle, board, index);

    if (isValidPuzzle && puzzle.rotation < 4 && puzzle.active) {
        board.AddPuzzlePiece(puzzle, index);
        archived.push(puzzleList.shift());
        return PlacePuzzle(puzzleList, board, index + 1, archived, counter);
    }

    if (!isValidPuzzle && puzzle.rotation < 4 && puzzle.active) {
        puzzleList[0].Rotate();
        return PlacePuzzle(puzzleList, board, index, archived, counter);
    }

    if (puzzleList[1] && puzzleList[1].active) {
        puzzleList[0].Reset();
        puzzleList[0].active = false;
        puzzleList.push(puzzleList.shift());
        return PlacePuzzle(puzzleList, board, index, archived, counter);
    }

    if (archived.length > 0) {
        const prePuzzle = archived.pop();
        puzzleList.forEach(puzzle => puzzle.Reset());
        prePuzzle.Rotate();
        puzzleList.unshift(prePuzzle);
        return PlacePuzzle(puzzleList, board, index - 1, archived, counter);
    }

    return;
}

const board = new Board(3)

// red bird: 1
// yellow bird: 2
// blue bird: 3
// green bird: 4
const puzzleList = [
    new Puzzle(-3, -4, 2, -1),
    new Puzzle(2, -1, 3, 2),
    // new Puzzle(-1, -2, 4, 1),
    // new Puzzle(-4, 4, -3, 1),
    // new Puzzle(2, 3, -2, 4),
    // new Puzzle(2, -1, 3, -1),
    // new Puzzle(-4, 1, 2, -3),
    // new Puzzle(-4, -4, -3, 2),
    // new Puzzle(3, 4, 2, -4),
];


PlacePuzzle(puzzleList, board);

console.log(board.squares);