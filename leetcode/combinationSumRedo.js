


const combinationSum = (candidates, target, combination = [], startIndex = 0, sum = 0) => {
  let solutions = [];


  if (sum === target) {
    solutions.push(combination.slice());
  } else if (sum < target) {
    
    for (let i = startIndex; i < candidates.length; i++) {
      let candidate = candidates[i];

      sum += candidate;
      combination.push(candidate);
      solutions = solutions.concat(combinationSum(candidates, target, combination, i, sum));
      sum -= candidate;
      combination.pop();
    }
  }


  return solutions;
}

console.log(combinationSum([2, 3, 6, 7], 7));