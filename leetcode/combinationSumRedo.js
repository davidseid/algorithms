


const combinationSum = (candidates, target, combination = [], startIndex = 0, sum = -1) => {
  let solutions = [];

  console.log(combination, sum);

  if (sum === target) {
    solutions.push(combination);
  } else if (sum < target) {
    
    for (let i = startIndex; i < candidates.length; i++) {
      let candidate = candidates[i];

      if (sum < 0) {
        sum = candidate;
      } else {
        sum += candidate;
      }
      let nextCombo = combination.slice();
      nextCombo.push(candidate);
      solutions = solutions.concat(combinationSum(candidates, target, nextCombo, i, sum));
      sum -= candidate;
    }
  }

  return solutions;
}

console.log(combinationSum([2, 3, 6, 7], 7));