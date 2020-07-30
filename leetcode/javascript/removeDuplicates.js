const removeDuplicates = (nums) => {
  // return length of modified array

  // get current element, if the next element is the same, delete it, 
  // if different get current element 

  let i = 0;
  while (i < nums.length) {
    let el = nums[i];
    let j = 1;
    while (nums[i + j] === el) {
      nums.splice(i + j, 1);
    }
    i++;
  }
  
  return nums.length;
}

let test = [1, 1, 2];
console.log(removeDuplicates(test));