const findPath = (grid, r = grid.length - 1, c = grid[0].length - 1, path = [], failedPoints = {}) => {

}



const testGrid = new Array(4);
for (let i = 0; i < 4; i ++) {
  testGrid[i] = new Array(4).fill(null);
}

testGrid[2][0] = 'blocked';
testGrid[0][3] = 'blocked';
testGrid[1][2] = 'blocked';

console.log(findPath(testGrid));