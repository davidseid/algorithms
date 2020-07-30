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
  
  for (let rowStart = 0; rowStart < board.length; rowStart += 3) {
    for (let colStart = 0; colStart < board.length; colStart += 3) {
      let square = [];

      for (let i = 0; i < 3; i++) {
        let row = rowStart + i;
        for (let j = 0; j < 3; j++) {
          let col = colStart + j;
          square.push(board[row][col]);
        }
      }
      if (!isValidUnit(square)) return false;
    }
  }

  return true;
}

const testInput = [
  ["8","3",".",".","7",".",".",".","."],
  ["6",".",".","1","9","5",".",".","."],
  [".","9","8",".",".",".",".","6","."],
  ["8",".",".",".","6",".",".",".","3"],
  ["4",".",".","8",".","3",".",".","1"],
  ["7",".",".",".","2",".",".",".","6"],
  [".","6",".",".",".",".","2","8","."],
  [".",".",".","4","1","9",".",".","5"],
  [".",".",".",".","8",".",".","7","9"]
];

console.log(isValidSudoku(testInput));
