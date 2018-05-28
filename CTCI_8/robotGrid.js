const robotGrid = (grid) => {
  if (grid === null || grid.length === 0) return null;
  path = new Array();
  failedPoints = new Object();

  if (findPath(grid, grid.length - 1, grid[0].length - 1, path, failedPoints)) {
    return path;
  }

  return null;
}



const findPath = (grid, r, c, path, failedPoints) => {
  // if the row is off the grid or the value doesn't exist return false;
  if (c < 0 || r < 0 || !grid[r][c]) {
    return false;
  }

  // save coordinates as a point
  let point = JSON.stringify([r, c]);
  
  // if the point is in our failures, return false
  if (failedPoints[point]) return false;

  // find out if the point is at the origin
  let isAtOrigin = (r === 0) && (c === 0);

  // if at the origin, or if the path up or left is valid, add the path to the path
  if (isAtOrigin || findPath(grid, r, c - 1, path, failedPoints) || 
      findPath(grid, r - 1, c, path, failedPoints)) {
    path.concat(point)
    return true;
  }

  failedPoints[point];
  return false;
}



const testGrid = new Array(4);
for (let i = 0; i < 4; i ++) {
  testGrid[i] = new Array(4).fill(1);
}

testGrid[2][0] = 0;
testGrid[0][3] = 0;
testGrid[1][2] = 0;

console.log(robotGrid(testGrid));