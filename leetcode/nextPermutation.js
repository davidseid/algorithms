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
    for (let j = nums.length - 1; j > i; j--) {
      let secondNum = nums[j];
      if (secondNum > firstNum) {
        nums[j] = firstNum;
        nums[i] = secondNum;
        let start = i + 1;

        while(start < nums.length) {
          let minIndex = start;
          for (let k = start; k < nums.length; k++) {
            if (nums[k] < nums[minIndex]) {
              minIndex = k;
            }
          }
          if (minIndex !== start) {
            let min = nums[minIndex];
            nums[minIndex] = nums[start];
            nums[start] = min;
          }
          start++;
        }

        return nums;
      }
    }
  }
  nums.sort((a, b) => a - b);
  return nums;
}

console.log(nextPermutation([9,1,9,6,3,6,8,4,6,6,7,2,7,4,1,0,0])); 