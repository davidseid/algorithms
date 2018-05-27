const findSubsets = (set, subset = [], subsets = []) => {
  for (let i = 0; i < set.length; i++) {
    let subset1 = subset.slice();
    let subset2 = subset.slice();
    subset2.push(set[i]);
    let nextSet = set.slice(i + 1);
    subsets.push(subset1, subset2);
    findSubsets(nextSet, subset1, subsets);
    findSubsets(nextSet, subset2, subsets);
  }

  let uniq = {};
  subsets.forEach((subset) => {
      uniq[JSON.stringify(subset)] = true;
  });
  let result = Object.keys(uniq);
  return result.map((subset) => JSON.parse(subset));
}

const testSet = [1, 3, 4];
findSubsets(testSet);