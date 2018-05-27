const findSubsets = (set) => {
  //using combinatorics
  // get the number of sets we expect
  let allSubsets = [];
  let numSets = 1 << set.length;
  for (let i = 0; i < numSets; i++) {
    let subset = [];
    let index = 0;
    for (let j = i; j > 0; j >>= 1) {
      if ((j & 1) === 1) {
        subset.push(set[index]);
      }
      index++;
    }
    allSubsets.push(subset);
  }
  return allSubsets;
}

const testSet = [1, 3, 4, 5];
console.log(findSubsets(testSet));