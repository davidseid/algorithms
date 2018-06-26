


const rotate = (matrix) => {
  let queue = [];
  let n = matrix.length;

  for (let i = 0; i < n; i++) {
    queue.push(matrix[0][i]);
  }

  for (let i = 0; i < n; i++) {
    queue.push(matrix[i][n -1]);
    matrix[i][n - 1] = queue.shift();
  }

  for (let i = n - 1; i >= 0; i--) {
    queue.push(matrix[n - 1][i]);
    matrix[n - 1][i] = queue.shift();
  }
  console.log(queue)

  queue.shift();

  for (let i = n - 2; i >= 0; i--) {
    queue.push(matrix[n - 2][i]);
    matrix[n - 2][i] = queue.shift();
  }



}

let test = [
  [1,2,3],
  [4,5,6],
  [7,8,9]
]

rotate(test);
console.log(test);