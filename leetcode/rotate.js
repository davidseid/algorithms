
const rotate = (matrix) => {
  let n = matrix.length - 1; 

  let saved = [];
  let putBack = [];
  let level = 0;

  while (level < matrix.length - 1) {
    
    for (let i = 0 + level; i < n; i++) {
      saved.push(matrix[0 + level][i]);
    }
    
    putBack = saved;
    saved = [];
    
    for (let i = 0 + level; i < n; i++) {
      saved.push(matrix[i][n - level]);
      matrix[i][n - level] = putBack.shift();
    }
    
    
    putBack = saved;
    saved = [];
    
    for (let i = n - level; i > 0; i--) {
      saved.push(matrix[n - level][i]);
      matrix[n - level][i] = putBack.shift();
    }
    
    
    
    putBack = saved;
    saved = [];
    
    for (let i = n - level; i > 0; i--) {
      saved.push(matrix[i][0 + level]);
      matrix[i][0 + level] = putBack.shift();
    }
    
    putBack = saved;
    saved = [];
    
    for (let i = 0 + level; i < n; i++) {
      matrix[0 + level][i] = putBack.shift();
    }

    level++;
    n = n - 2;
  }
}
  
let test = [
  [ 5, 1, 9,11],
  [ 2, 4, 8,10],
  [13, 3, 6, 7],
  [15,14,12,16]
]

rotate(test);
console.log(test);