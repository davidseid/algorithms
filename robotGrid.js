const findPath = (grid, r = 0, c = 0) => {
  
  grid[r][c] = 'path';
  if (r === grid.length - 1 && c === grid.length - 1) {
    return grid;
  }

  nextGrid = JSON.parse(JSON.stringify(grid));
  
  if (nextGrid[r + 1] && nextGrid[r + 1][c] === null) {
    return findPath(nextGrid, r + 1, c);
  }

  if (nextGrid[r][c + 1] === null) {
    return findPath(nextGrid, r, c + 1);
  }

}

const testGrid = new Array(4);
for (let i = 0; i < 4; i ++) {
  testGrid[i] = new Array(4).fill(null);
}

testGrid[2][0] = 'blocked';
testGrid[0][3] = 'blocked';
testGrid[1][2] = 'blocked';

console.log(findPath(testGrid));