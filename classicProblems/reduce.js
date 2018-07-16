

const reduce = (array, iteratee, accumulator) => {
  let start;

  if (accumulator) {
    start = 0;
  } else {
    start = 1;
    accumulator = array[0];
  }
  
  for (let i = 1; i < array.length; i++) {
    accumulator = iteratee(accumulator, array[i]);
  }

  return accumulator;
}

const sum = (a, b) => {
  return a + b;
}

let test = [1, 3, 4, 7, -4];

console.log(reduce(test, sum, -11));