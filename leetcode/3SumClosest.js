const threeSumClosest = (nums, target) => {
  
  const sortedNums = nums.sort((a, b) => a - b);

  let highestLowerSum = sortedNums[0] + sortedNums[1] + sortedNums[2];
  let lowestHigherSum = sortedNums[sortedNums.length - 1] + sortedNums[sortedNums.length - 2] + sortedNums[sortedNums.length - 3];

  if (highestLowerSum === lowestHigherSum || target < highestLowerSum) return highestLowerSum;
  if (target > lowestHigherSum) return lowestHigherSum;

  for (let i = 0; i < nums.length - 2; i++) {
    let val1 = sortedNums[i];
    for (let j = i + 1; j < nums.length - 1; j++) {
      let val2 = sortedNums[j];
      for (let k = j + 1; k < sortedNums.length; k++) {
        let val3 = sortedNums[k];
        let sum = val1 + val2 + val3;
        if (sum === target) return sum;
        if (sum > highestLowerSum && sum < target) {
          highestLowerSum = sum;
        }
        if (sum > target && sum < lowestHigherSum) {
          lowestHigherSum = sum;
          break;
        }
      }
    }
  }

  let leftDiff = target - highestLowerSum;
  let rightDiff = lowestHigherSum - target;
  if (leftDiff < rightDiff) return highestLowerSum;
  return lowestHigherSum;
}

console.log(threeSumClosest([-1, 2, 1, -4], 1));