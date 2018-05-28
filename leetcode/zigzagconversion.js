var convert = function(s, numRows) {
  let grid = [];

  for (let i = 0; i < numRows; i++) {
    let row = [];
    row.length = s.length;
    grid.push(row);
  }
  
  let currentIndex = 0;
  let column = 0;
  while (currentIndex < s.length) {
    for (let i = 0; i < grid.length; i++) {
      grid[i][column] = s[currentIndex];
      currentIndex++;
    }
    let currentRow = numRows - 2;
    column++;
    for (let i = currentRow; i > 0; i--) {
      grid[currentRow][column] = s[currentIndex];
      currentRow--;
      column++;
      currentIndex++;
    }
  }


  let result = '';

  for (let i = 0; i < grid.length; i++) {
    for (let j = 0; j < grid[0].length; j++) {
      if (grid[i][j] !== undefined) {
        result += grid[i][j];
      }
    }
  }
  return result;
};