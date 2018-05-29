const removeElement = (nums, val) => {
  let i = 0;
  while (i < nums.length) {
    while (nums[i] === val) {
      nums.splice(i, 1);
    }
    i++;
  }
  return nums.length;
}

let test = [3, 2, 2, 3];
console.log(removeElement(test, 3));