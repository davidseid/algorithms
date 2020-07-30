

const permute = (nums, permutation = []) => {
  let result = []; 

  if (nums.length === 0) {
    result.push(permutation);
    return result;
  }

  for (let i = 0; i < nums.length; i++) {
    let nextPermutation = permutation.slice();
    nextPermutation.push(nums[i]);

    let nextNums = nums.slice(0, i).concat(nums.slice(i + 1));

    result = result.concat(permute(nextNums, nextPermutation));

  }
  return result;
}

console.log(permute([1, 2, 3]));