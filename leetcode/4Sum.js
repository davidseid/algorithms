const fourSum = (nums, target) => {
  let sortedNums = nums.sort((a, b) => a - b);
  let storage = {};
  for (let i = 0; i < sortedNums.length - 3; i++) {
    let val1 = sortedNums[i];
    for (let j = i + 1; j < sortedNums.length - 2; j++) {
      let val2 = sortedNums[j];
      for (let k = j + 1; k < sortedNums.length - 1; k++) {
        let val3 = sortedNums[k];
        for (let l = k + 1; l < sortedNums.length; l++) {
          let val4 = sortedNums[l];
          if (val1 + val2 + val3 + val4 === target) {
            if (!storage[JSON.stringify([val1, val2, val3, val4])]) {
              storage[JSON.stringify([val1, val2, val3, val4])] = true;
            }
          }
        }
      }
    }
  }
  let result = Object.keys(storage);
  result = result.map(set => JSON.parse(set));
  return result;
}

console.log(fourSum([1, 0, -1, 0, -2, 2], 0));