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


const isValidSudoku = (board) => {

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

