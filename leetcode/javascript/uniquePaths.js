

/* 
    Dynamic Programming problem
    robot top left corner of m x n grid
    robot can move down or right, trying to get to bottom right corner
    how many possible unique paths?

    Approach 1: recursively find all the paths, and increment each time we get to the finish
    Approahc 2: use dynamic programming to work backwards from the finish
*/

const uniquePaths = (m, n) => {

    const grid = [...Array(m)].map(e => Array(n).fill(0));
    
    grid[m - 1][n - 1] = 1;

    for (let i = m - 1; i >= 0; i--) {
        for (let j = n - 1; j >= 0; j--) {
            if (grid[i + 1]) {
                grid[i][j] += grid[i+1][j];
            }
            if (grid[i][j + 1]) {
                grid[i][j] += grid[i][j + 1];
            }
        }
    }

    return grid[0][0];
}

uniquePaths(7, 3);