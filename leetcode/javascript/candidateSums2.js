

const combinationSum2 = (candidates, target, combo = [], sum = 0, validCombos = {}) => {
  let result = [];

  for (let i = 0; i < candidates.length; i++) {
    let nextCombo = combo.slice();
    nextCombo.push(candidates[i]);

    let nextSum = nextCombo.reduce((a, b) => a + b);

    if (nextSum === target) {
      nextCombo = nextCombo.sort((a, b) => a - b);
      let key = JSON.stringify(nextCombo);
      if (!validCombos[key]) {
        validCombos[key] = true;
        result.push(nextCombo);
      }
    } else if (nextSum < target) {
      nextCandidates = candidates.slice(i + 1);
      result = result.concat(combinationSum2(nextCandidates, target, nextCombo, nextSum, validCombos));
    }
  }

  return result;
}

let testCandidates = [10,1,2,7,6,1,5];

console.log(combinationSum2(testCandidates, 8));