const threeSum = (nums) => {
  // get hashtable of all the positive numbers with index
  // for each combination 

  const resultHash = {};
  
  nums = nums.sort((a, b) => {
    return a - b;
  });

  const hash = {};
  for (let i = 0; i < nums.length; i++) {
    if (nums[i] >= 0) {
      hash[nums[i]] = i;
    }
  }

  if (nums[0] === 0 && nums[nums.length - 1] === 0 && nums.length >= 3) return [[0, 0, 0]];

  for (let i = 0; i < nums.length - 2; i++) {
    if (nums[i] > 0) break;
    let val1 = nums[i];
    for (let j = i + 1; j < nums.length - 1; j++) {
      let val2 = nums[j];
      let combo = val1 + val2;
      if (combo > 0) break;
      let target = combo * -1;
      if (hash[target] > j) {
        let set = [val1, val2, target];
        if (!resultHash[JSON.stringify(set)]) {
          resultHash[JSON.stringify(set)] = set;
        }
      }
    }
  }

  return Object.keys(resultHash).map((set) => JSON.parse(set));

}


console.log(threeSum([-1, 0, 1, 2, -1, -4]));