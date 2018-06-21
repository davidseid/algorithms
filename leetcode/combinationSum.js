
/*
  input: set of candidate numbers (without dupes), target num (only positive nums)
  output: find all unique combinations where candidates sum to target
  constraints: no duplicate combinations in solution set
  edge cases: candidates is an empty array, nothing in solution set, 

  2, 3, 6, 7 

            2
    2  3  6  7
  2 3 6 7
  2 3 6 7


*/

const combinationSum = (candidates, target, combination = [], sum = 0) => {
  let result = [];

  for (let i = 0; i < candidates.length; i++) {
    let nextCombo = combination.slice();
    nextCombo.push(candidates[i]);
    let nextSum = sum + candidates[i];
    let nextCandidates = candidates.slice(i);

    if (nextSum === target) {
      result.push(nextCombo);
    } else if (nextSum < target) {
      result = result.concat(combinationSum(nextCandidates, target, nextCombo, nextSum));
    }
  }

  return result;


}

const testCandidates = [2, 3, 6, 7];

console.log(combinationSum(testCandidates, 7));