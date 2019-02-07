const trap = (height) => {
  let water = 0;
  let leftMax = 0;
  let rightMax = 0;
  let n = height.length;
  let left = 0;
  let right = n - 1;

  while (left < right) {

    if (height[left] < height[right]) {
        if (height[left] > leftMax) {
            leftMax = height[left];
        } else {
            water += leftMax - height[left];
        }
        left++;
    } else {
        if (height[right] > rightMax) {
            rightMax = height[right];
        } else {
            water += rightMax - height[right];
        }
        right--;
    }
  }

  return water;

};


const test = [0,1,0,2,1,0,1,3,2,1,2,1];
console.log(trap(test));
