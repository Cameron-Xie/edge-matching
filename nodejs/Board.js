const Square = require('./Square');

module.exports = class Board {
    constructor(n) {
        this.cols = n;
        this.squares = Board.CreateBoard(n);
    }

    AddPuzzlePiece(puzzle, index) {
        this.squares[index].puzzle = puzzle;
    }

    static CreateBoard(n) {
        const squares = [];
        for (let row = 0; row < n; row++) {
            for (let col = 0; col < n; col++) {
                squares.push(new Square(row, col));
            }
        }

        return squares;
    }
}
