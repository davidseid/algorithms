// input: array, number target (sorted array, with duplicates)
// output: array of start and end
// constraints: log n time
// edge cases: no dupes, not exist, 


const searchRange = (nums, target) => {
  // can't do linear search, must do log n
  // do binary search for num, then scan for left and right, return coordinates
  if (nums.length === 1) {
    if (nums[0] === target) return [0, 0];
    if (nums[0] !== target) return [-1, -1];
  }

  let leftIndex = 0;
  let rightIndex = nums.length - 1;

  while (leftIndex <= rightIndex) {
    let midIndex = leftIndex + Math.floor((rightIndex - leftIndex) / 2);
    if (nums[midIndex] === target) {
      let rangeStart = midIndex;
      let rangeEnd = midIndex;

      while (nums[rangeStart - 1] === target) {
        rangeStart--;
      }

      while (nums[rangeEnd + 1] === target) {
        rangeEnd++;
      }

      return [rangeStart, rangeEnd];


    } else if (nums[midIndex] > target) {
      if (rightIndex === midIndex) {
        rightIndex = midIndex - 1;
      } else {

        rightIndex = midIndex;
      }
    } else if (nums[midIndex] < target) {
      if (midIndex === leftIndex) {
        leftIndex = midIndex + 1;
      } else {
        leftIndex = midIndex;
      }
    }
  }
  return [-1, -1];
}

let testNums = [5,7,7,8,8,10]

console.log(searchRange(testNums, 6));