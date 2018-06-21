
const findPivot = (nums) => {
  let startIndex = 0;
  let endIndex = nums.length - 1;

  let pivot;
  if (nums[endIndex] > nums[startIndex]) {
    pivot = -1;
  }

  while (pivot === undefined) {
    let midIndex = startIndex + Math.floor((endIndex - startIndex) / 2);
    if (nums[startIndex + 1] < nums[startIndex]) {
      pivot = startIndex;
    } else if (nums[midIndex + 1] < nums[midIndex]) {
      pivot = midIndex;
    }
    
    if (nums[startIndex] > nums[midIndex]) {
      endIndex = midIndex;
    } else if (nums[midIndex] > nums[endIndex]) {
      startIndex = midIndex;
    }
  }
  return pivot;
}



const search = (nums, target) => {
  const pivot = findPivot(nums);

  if (target <= nums[pivot] && target >= nums[0]) {
    return binSearch(0, pivot, nums, target);
  } else if (target >= nums[pivot + 1] && target <= nums[nums.length - 1]) {
    return binSearch(pivot + 1, nums.length - 1, nums, target);
  } else {
    return -1;
  }
}

const binSearch = (startIndex, endIndex, nums, target) => {
  let found = false;

  while (!found) {
    let midIndex = startIndex + (Math.floor((endIndex - startIndex) / 2));
    if (target === nums[midIndex]) {
      return midIndex;
    } else if (nums[midIndex] < target) {
      startIndex = midIndex;
    } else if (nums[midIndex] > target) {
      endIndex = midIndex;
    }
  }
}


console.log(search([4,5,6,7,0,1,2], 0));