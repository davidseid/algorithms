

/*
i: array pos and neg integers
output: int, sum of the longest subarray (with at least one num)
constraints: none
edge cases: empty array, 
[-2,1,-3,4,-1,2,1,-5,4]

// brute force
  find all substrings
  for each, update biggest if it is bigger
  
  start: 3
  pointer: 6
  subarray: 4, -1, 2, 1
  sum: 6
  freshsum = 1
  max 6
*/

const maxSubarray = (nums) => {
  let start = 0;
  let pointer = 1;
  let sub = [nums[0]];
  let sum = nums[0];
  let max = sum;

  while (pointer < nums.length) {
    let nextNum = nums[pointer];
    if (sum + nextNum > nextNum) {
      sub.push(nextNum);
      sum += nextNum;
    } else {
      sub = [nextNum];
      start = pointer;
      sum = nextNum;
    }
    if (sum > max) max = sum;
    pointer++;
  }
  return max;
}

let testNums = [-2,1,-3,4,-1,2,1,3,-5,4];

console.log(maxSubarray(testNums));