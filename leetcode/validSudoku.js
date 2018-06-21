/*
  input: Array of 9 arrays with slots each
  output: bool whether valid
  constraints: may not be filled 
  edge cases: can have '.' for many slots

  three validity checks, each row has no repeats of any number, only . 1 - 9
    - each column is the same
    - each group of 3x3 is the same

  build check that looks at 9 pieces
  systematically run all columns, all rows, and all 3x3 grids through that box
  return false if any fails, otherwise return true


*/

const isValidUnit = (arr) => {
  const numbers = {};

  for (let i = 0; i < arr.length; i++) {
    if (arr[i] !== '.' && numbers[arr[i]]) {
      return false;
    } else {
      numbers[arr[i]] = true;
    }
  }
  return true;
}


const isValidSudoku = (board) => {
  
  for (let i = 0; i < board.length; i++) {
    let row = board[i];
    if (!isValidUnit(row)) {
      return false;
    }
  }

  for (let i = 0; i < 9; i++) {
    let column = [];
    for (let j = 0; j < 9; j++) {
      column.push(board[j][i]);
    }
    if (!isValidUnit(column)) return false;
  }
}

const testInput = [
  ["5","3",".",".","7",".",".",".","."],
  ["6",".",".","1","9","5",".",".","."],
  [".","9","8",".",".",".",".","6","."],
  ["8",".",".",".","6",".",".",".","3"],
  ["4",".",".","8",".","3",".",".","1"],
  ["7",".",".",".","2",".",".",".","6"],
  [".","6",".",".",".",".","2","8","."],
  [".",".",".","4","1","9",".",".","5"],
  [".",".",".",".","8",".",".","7","9"]
];

