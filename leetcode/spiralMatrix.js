



const spiralMatrix = (matrix) => {
  let total = matrix.length * matrix[0].length;

  const result = [];
  let top = 0;
  let bottom = matrix.length - 1;
  let left = 0;
  let right = matrix[0].length - 1;

  while (result.length < total) {

    for (let i = 0 + left; i <= right; i++) {
      result.push(matrix[top][i]);
    }
    top++;

    if (result.length === total) return result;

    for (let i = top; i <= bottom; i++) {
      result.push(matrix[i][right]);
    }
    right--;

    if (result.length === total) return result;

    for (let i = right; i >= left; i--) {
      result.push(matrix[bottom][i]);
    }

    bottom--;

    if (result.length === total) return result;

    for (let i = bottom; i >= top; i--) {
      result.push(matrix[i][left]);
    }
    
    left++;
  }

  return result;
  

}

let testMatrix = [
  [ 1, 2, 3 ],
  [ 4, 5, 6 ],
  [ 7, 8, 9 ]
 ]

console.log(spiralMatrix(testMatrix));