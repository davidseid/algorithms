const canJump = (nums, index = 0) => {

  let result = false;

  let maxJump = nums[index];

  let target = nums.length - 1;

  if (index + maxJump >= target) return true;

  for (let i = 1; i <= maxJump; i++) {
    return result || canJump(nums, index + i);
  }

  return result;
}

console.log(canJump([2, 0, 0]));