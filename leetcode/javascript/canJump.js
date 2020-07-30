const canJump = (nums, index = 0) => {
  let lowestValidIndex = nums.length - 1;

  for (let i = nums.length - 1; i >= 0; i--) {
    if (i + nums[i] >= lowestValidIndex) {
      lowestValidIndex = i;
    }
  }

  if (lowestValidIndex === 0) return true;
  return false;
}

console.log(canJump([3, 2, 1, 1, 4]));