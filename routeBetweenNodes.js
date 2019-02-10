

// given directed graph
// find wehteher there is a route between two nodes

const isRoute = (graph, start, end) => {
  // use bfs to search from start, once the end is found, return true, otherwise return false;
  let result = false;
  let queue = [];

  if (start === end) return true;

  queue.push(start);

  while (queue.length > 0) {
    let next = queue.shift();
    if (next === end) return true;
  
    graph[next].forEach((key) => {
      queue.push(key);
    });
  }

  return false;
}


const testGraph = {
  1: [2, 5],
  2: [6],
  5: [7, 9],
  6: [],
  9: [],
  7: []
}

console.log(isRoute(testGraph, 5, 6));