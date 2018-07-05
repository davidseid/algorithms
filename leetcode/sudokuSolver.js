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
testBoard.board[0][0] = '5';
testBoard.board[0][1] = '3';
testBoard.board[0][4] = '7';
testBoard.board[1][0] = '6';
testBoard.board[1][3] = '1';
testBoard.board[1][4] = '9';
testBoard.board[1][5] = '5';
testBoard.board[2][1] = '9';
testBoard.board[2][2] = '8';
testBoard.board[2][7] = '6';
testBoard.board[3][0] = '8';
testBoard.board[3][4] = '6';
testBoard.board[3][8] = '3';
testBoard.board[4][0] = '4';
testBoard.board[4][3] = '8';
testBoard.board[4][5] = '3';
testBoard.board[4][8] = '1';
testBoard.board[5][0] = '7';
testBoard.board[5][4] = '2';
testBoard.board[5][8] = '6';
testBoard.board[6][1] = '6';
testBoard.board[6][6] = '2';
testBoard.board[6][7] = '8';
testBoard.board[7][3] = '4';
testBoard.board[7][4] = '1';
testBoard.board[7][5] = '9';
testBoard.board[7][8] = '5';
testBoard.board[8][4] = '8';
testBoard.board[8][7] = '7';
testBoard.board[8][8] = '9';


const buildPossibilities = (board) => {

  for (let i = 0; i < board.length; i++) {
    let row = board[i];
    for (let j = 0; j < row.length; j++) {
      let box = row[j];
      if (box === '.') {
        board[i][j] = { remaining: 9 };
        
        for (let k = 1; k <= 9; k++) {
          board[i][j][k] = true;
        }

      }
    }
  }
}

const calculatePossibilitiesForBox = (board, r, c) => {
  for (let i = 1; i <= 9; i++) {
    i = i.toString();
    if (inRow(board, r, i) || inCol(board, c, i) || inArea(board, r, c, i)) {
      if (board[r][c][i]) {
        board[r][c][i] = false;
        board[r][c].remaining--;
      }
    }
  }
}

const narrowBox = (board, r, c) => {
  let options = Object.keys(board[r][c]);

  for (let i = 0; i < options.length - 1; i++) {
    let option = options[i];
    if (inRow(board, r, option) || inCol(board, c, option) || inArea(board, r, c, option)) {
      if (board[r][c][option]) {
        board[r][c][option] = false;
        board[r][c].remaining--;
      }
    }
  }
}

const narrowPossibilities = (board) => {
  for (let i = 0; i < 9; i++) {
    for (let j = 0; j < 9; j++) {
      if (typeof board[i][j] !== 'string') {
        narrowBox(board, i, j);
      }
    }
  }
}

const calculatePossibilities = (board) => {
  for (let i = 0; i < 9; i++) {
    for (let j = 0; j < 9; j++) {
      if (typeof board[i][j] !== 'string') {
        calculatePossibilitiesForBox(board, i, j);
      }
    }
  }
}

const updateBoard = (board) => {
  let updates = 0;
  for (let i = 0; i < 9; i++) {
    for (let j = 0; j < 9; j++) {
      let box = board[i][j];
      if (box.remaining === 1) {
        let value = Object.keys(box).filter((key) => box[key] === true).pop();
        board[i][j] = value;
        updates++;
      }
    }
  }
  return updates;
}

const inArea = (board, r, c, num) => {
  let rowStart;
  let colStart;

  if (r <= 2) rowStart = 0;
  if (r > 2 && r < 6) rowStart = 3;
  if (r > 5) rowStart = 6;

  if (c <= 2) colStart = 0;
  if (c > 2 && c < 6) colStart = 3;
  if (c > 5) colStart = 6;

  for (let i = rowStart; i < rowStart + 3; i++) {
    for (let j = colStart; j < colStart + 3; j++) {
      if (board[i][j] === num) {
        return true;
      }
    }
  }
  return false;
}

const inRow = (board, r, num) => {
  for (let i = 0; i < 9; i++) {
    if (board[r][i] === num) {
      return true;
    }
  }
  return false;
}

const inCol = (board, c, num) => {
  for (let i = 0; i < 9; i++) {
    if (board[i][c] === num) {
      return true;
    }
  }
  return false;
}

const countSpots = (board) => {
  let count = 0;
  for (let i = 0; i < 9; i++) {
    for (let j = 0; j < 9; j++) {
      if (board[i][j] === '.') {
        count++;
      }
    }
  }
  return count;
}



// optimizations:
//   don't duplicate board
//   make narrow down pattern
// only check remaining empties (store coordinates and use length)

const solveSudoku = (board) => {

  let spotsRemaining = countSpots(board);
  buildPossibilities(board);
  calculatePossibilities(board);
  while (spotsRemaining > 0) {
    let numUpdates = updateBoard(board);
    spotsRemaining -= numUpdates;
    narrowPossibilities(board);
  }
}

console.log(solveSudoku(testBoard.board));