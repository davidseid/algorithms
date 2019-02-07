console.time('findPath');

const buildGrid = (r, c, blocks) => {
    const grid = new Array(r).fill(0).map(r => new Array(c).fill(0));
    
    for (let i = 0; i < blocks.length; i++) {
        let block = blocks[i];
        grid[block[0]][block[1]] = 'blocked';
    }
    console.log(grid);
    return grid;
};

const findPath = (grid, r = 0, c = 0, memo = {}) => {

    let point = JSON.stringify([r, c]);
    if (memo[point]) return false;

    grid[r][c] = 1;
    if (r === grid.length - 1 && c === grid.length - 1) {
        return true;
    }

    let leftPath;
    let rightPath;

    if (openRight(grid, r, c)) {
        rightPath = findPath(grid, r, c + 1);
    }

    if (openDown(grid, r, c)) {
        leftPath = findPath(grid, r + 1, c);
    }

    if (!rightPath && !leftPath) {
        grid[r][c] = 0;
        memo[point] = true;
        return false;
    }
    return grid;
};

const openDown = (grid, r, c) => {

    if (grid[r + 1] && grid[r + 1][c] !== 'blocked') {
        return true;
    }
    return false;
};

const openRight = (grid, r, c) => {
    if (grid[r][c + 1] !== 1 && grid[r][c + 1] !== 'blocked') {
        return true;
    }
    return false;
};

const testGrid = buildGrid(4, 4, [[0, 2], [1, 1], [2, 2]]);

console.log(findPath(testGrid));
console.timeEnd('findPath');