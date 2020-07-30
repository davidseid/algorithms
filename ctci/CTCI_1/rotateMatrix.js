let matrix = [
  [1, 2, 3],
  [4, 5, 6],
  [7, 8, 9]
];


const rotateMatrix = (matrix) => {
  let m = matrix.length - 1;

  for (let l = 0; l < matrix.length / 2; l++) {

    for (let i = 0; i < matrix.length - 1 - l; i++) {


      let next = matrix[l][i + l];
      let saved = matrix[i + l][m - l];
      matrix[i + l][m - l] = next;
      next = saved;
      
      saved = matrix[m - l][m - i - l];
      matrix[m - l][m - i - l] = next;
      next = saved;
      
      saved = matrix[m - i - l][l];
      matrix[m - i - l][l] = next;
      next = saved;

      matrix[l][i + l] = next;
    }
  }
}





























// const rotateMatrix = (m) => {
//   if (m.length === 0 || m.length !== m[0].length) return false;
//   let n = matrix.length;

//   for (let layer = 0; layer < n / 2; layer++) {
//     let first = layer;
//     let last = n - 1 - layer;

//     for (let i = first; i < last; i++) {
//       let offset = i - first;
//       let top = m[first][i];
//       console.log(i);

//       m[first][i] = m[last - offset][first];

//       m[last - offset][first] = m[last][last - offset];

//       m[last][last - offset] = m[i][last];

//       m[i][last] = top;
//     }
//   }
//   return true;
// }

console.log(matrix);

rotateMatrix(matrix);

console.log(matrix);