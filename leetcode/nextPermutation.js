// if next lexicographical permutation not possible
// rearrange in lowest possible order
// in place replacement and only constant extra memory

// 1 2 3 -> 1 3 2
// 3 2 1 -> 1 2 3
// 1 1 5 -> 1 5 1

// starting from the end, if the number is bigger than any to the left, shift it
// if none, sort by ascending

const nextPermutation = (nums) => {
  for (let i = nums.length - 1; i >= 0; i--) {
    let firstNum = nums[i];
    for (let j = i; j >= 0; j--) {
      let secondNum = nums[j];
      if (firstNum > secondNum) {
        nums[j] = firstNum;
        for (let k = j; k < nums.length - 1; k++) {
          let tmp = nums[k + 1];
          nums[k + 1] = secondNum;
          secondNum = tmp;
        }
        return nums;
      }
    }
  }
  nums.sort((a, b) => a - b);
  return nums;
}

console.log(nextPermutation([1, 2, 3]));