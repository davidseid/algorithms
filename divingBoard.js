const buildDivingBoard = (shorter, longer, k) => {
  // return possible lengths

  const lengths = {};
  const visited = {};

  getAllLengths(k, 0, shorter, longer, lengths, visited);

  return Object.keys(lengths);
  
}

const getAllLengths = (k, total, shorter, longer, lengths, visited) => {
  if (k === 0) {
    lengths[total] = true; 
    return;
  }
  key = k + " " + total;

  if (visited[key]) return;
  getAllLengths(k - 1, total + shorter, shorter, longer, lengths, visited);
  getAllLengths(k - 1, total + longer, shorter, longer, lengths, visited);
  visited[key];
}

console.log(buildDivingBoard(4, 6, 3));