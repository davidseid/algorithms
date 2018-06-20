/*
  4 5 6 7 0 1 2


*/



const search = (nums, target) => {
  let startIndex = 0;
  let endIndex = nums.length - 1;
  let midIndex = startIndex + Math.floor((endIndex - startIndex) / 2);

  let pivot;
  if (nums[endIndex] > nums[startIndex]) {
    pivot = -1;
  }

  while (pivot === undefined) {
    if (nums[startIndex + 1] < nums[startIndex]) {
      pivot = startIndex;
    } else if (nums[midIndex + 1] < nums[midIndex]) {
      pivot = midIndex;
    }

  }
}

console.log(search([3, 4, 5, 6, 1, 2], 0));