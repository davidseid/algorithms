
/*
Given a positive integer n, generate a square matrix filled with elements from 1 to n2 in spiral order.

Example:

Input: 3
Output:
[
    [ 1, 2, 3 ],
    [ 8, 9, 4 ],
    [ 7, 6, 5 ]
]

// create the empty matrix
// start on row 0 increment cols
// start on last column, increment rows
// start on last row, decrement colums
// start on first col, decrement rows
// when get to a number, switch out
*/

const generateMatrix = (n) => {

    const matrix = new Array(n).fill(0).map((e) => new Array(n).fill(0));

    let offset = 0;
    let num = 1;

    while (num <= n * n) {
        for (let i = 0 + offset; i < n - offset; i++) {
            matrix[offset][i] = num;
            num++;
        }
        for (let i = 1 + offset; i < n - offset; i++) {
            matrix[i][n - 1 - offset] = num;
            num++;
        }
        for (let i = n - 2 - offset; i >= 0 + offset; i--) {
            matrix[n - 1 - offset][i] = num;
            num++;
        }
        for (let i = n - 2 - offset; i >= 1 + offset; i--) {
            matrix[i][offset]  = num;
            num++;
        }
        offset++;
    }

    return matrix;
};

console.log(generateMatrix(5));