/*
  4 5 6 7 0 1 2


*/



const search = (nums, target) => {
  
  let pivot;
  let startIndex = 0;
  let endIndex = nums.length - 1;

  while (pivot === undefined && startIndex !== endIndex) {
    let start = nums[startIndex];
    let end = nums[endIndex];
    let midIndex = Math.floor(startIndex + Math.floor((endIndex - startIndex) / 2));
    let mid = nums[midIndex];
    console.log(start, mid, end)
    
    if (start > nums[startIndex + 1]) {
      pivot = startIndex;
    }
    if (mid > nums[midIndex + 1]) {
      pivot = midIndex;
    }

    if (end < mid) {
      startIndex = midIndex + 1;
      endIndex--
    }
    if (mid < start) {
      startIndex++;
      endIndex = midIndex - 1;
    }
  }
}

console.log(search([4, 5, 6, 7, 0, 1], 0));