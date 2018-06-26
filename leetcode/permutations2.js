const permuteUnique = (nums, permutation = [], unique = {}) => {
  let result = [];
  
  if (nums.length === 0) {
    
    let serialized = JSON.stringify(permutation);
    if (!unique[serialized]) {
      unique[serialized] = true;
      result.push(permutation);
      return result;
    } else {
      return [];
    }
  }

  for (let i = 0; i < nums.length; i++) {
    let nextPerm = permutation.slice();
    nextPerm.push(nums[i]);

    let nextNums = nums.slice();
    nextNums.splice(i, 1);
    result = result.concat(permuteUnique(nextNums, nextPerm, unique));
  }
  
  return result;

}

console.log(permuteUnique([1, 1, 2]));