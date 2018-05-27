const findSubsets = (set) => {
  // if set.length === [] return [[]]
  // if set.length ==== 1 return [[], [i]]
  // store each in a cache
  // to findSubsets (n), take each of hte previous sets and add the latest value
  let memo = {};
  memo[JSON.stringify([])] = [[]];
  memo[JSON.stringify(set.slice(0, 1))] = [[], [set[0]]];

  for (let i = 2; i <= set.length; i++) {
    let lastSubsets = memo[JSON.stringify(set.slice(0, i - 1))];
    let subsets = lastSubsets.slice();
    lastSubsets.forEach((subset) => {
      let copy = subset.slice();
      copy.push(set[i - 1]);
      subsets.push(copy);
    });
    memo[JSON.stringify(set.slice(0, i))] = subsets;
  }

  return memo[JSON.stringify(set.slice())];
}

const testSet = [1, 3, 4, 5];
findSubsets(testSet);