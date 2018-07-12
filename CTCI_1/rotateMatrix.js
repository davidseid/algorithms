
/*
  [
    [1, 2, 3],
    [4, 5, 6],
    [7, 8, 9]
  ]

  5 9 3 1   
  5 6 7 2
  3 8 7 6
  2 4 8 4

  7 4 1
  8 5 2
  9 6 3
*/


const rotateMatrix = (m) => {
  if (m.length === 0 || m.length !== m[0].length) return false;
  let n = matrix.length;

  for (let layer = 0; layer < n / 2; layer++) {
    let first = layer;
    let last = n - 1 - layer;

    for (let i = first; i < last; i++) {
      let offset = i - first;
      let top = m[first][i];
      console.log(i);

      m[first][i] = m[last - offset][first];

      m[last - offset][first] = m[last][last - offset];

      m[last][last - offset] = m[i][last];

      m[i][last] = top;
    }
  }
  return true;
}

let matrix = [
  [1, 2, 3],
  [4, 5, 6],
  [7, 8, 9]
];

console.log(matrix);

rotateMatrix(matrix);

console.log(matrix);