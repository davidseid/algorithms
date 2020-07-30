'use strict';

/*
- 1s and 0s
- count num islands
*/

const numIslands = (grid) => {
    let numIslands = 0;

    for (let i = 0; i < grid.length; i++) {
        
        for (let j = 0; j < grid[0].length; j++) {
            
            if (grid[i][j] === '1') {
                chartIsland(grid, i, j);
                numIslands++;
            }
        }
    }
    
    return numIslands;
};

// helpers
const chartIsland = (grid, r, c) => {
    grid[r][c] = '2';

    for (let i = -1; i < 2; i += 2) {
        if (grid[r + i] && grid[r + i][c] === '1') {
            chartIsland(grid, r + i, c);
        }
        if (grid[r][c + i] === '1') {
            chartIsland(grid, r, c + i);
        }
    }
};

const testGrid = [["1","1","1","1","0"],["1","1","0","1","0"],["1","1","0","0","0"],["0","0","0","0","0"]];

console.log(numIslands(testGrid));