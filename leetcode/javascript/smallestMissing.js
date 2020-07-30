


const findSmallestMissing = (arr) => {
  let hashedValues = {};

  for (let i = 0; i < arr.length; i++) {
    let val = arr[i];
    hashedValues[val] = true;
  }

  let i = 1;

  while (true) {
    if (!hashedValues[i]) {
      return i;
    }
    i++;
  }
}

let test = [7, 8, 9];

console.log(findSmallestMissing(test));