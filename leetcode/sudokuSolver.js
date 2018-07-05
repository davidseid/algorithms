/*
  input: board (with . in empty spaces)
  output: return nothing, modify board in place
  constraints: none
  edge cases: 9x9, only 1 valid solution

*/

class SudokuBoard {
  constructor() {
    this.board = new Array(9).fill(null).map((row) => new Array(9).fill('.'));
  }
}

let testBoard = new SudokuBoard();
testBoard.board[0][0] = 5;
testBoard.board[0][1] = 3;
testBoard.board[0][4] = 7;
testBoard.board[1][0] = 6;
testBoard.board[1][3] = 1;
testBoard.board[1][4] = 9;
testBoard.board[1][5] = 5;
testBoard.board[2][1] = 9;
testBoard.board[2][2] = 8;
testBoard.board[2][7] = 6;
testBoard.board[3][0] = 8;
testBoard.board[3][4] = 6;
testBoard.board[3][8] = 3;
testBoard.board[4][0] = 4;
testBoard.board[4][3] = 8;
testBoard.board[4][5] = 3;
testBoard.board[4][8] = 1;
testBoard.board[5][0] = 7;
testBoard.board[5][4] = 2;
testBoard.board[5][8] = 6;
testBoard.board[6][1] = 6;
testBoard.board[6][6] = 2;
testBoard.board[6][7] = 8;
testBoard.board[7][3] = 4;
testBoard.board[7][4] = 1;
testBoard.board[7][5] = 9;
testBoard.board[7][8] = 5;
testBoard.board[8][4] = 8;
testBoard.board[8][7] = 7;
testBoard.board[8][8] = 9;





console.log(testBoard);

const solveSudoku = (board) => {

}