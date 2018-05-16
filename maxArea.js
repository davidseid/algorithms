/*
i: array of integers
o: maxArea
c: none
e: none

[1, 1, 2, 3]

1 3 3 * min height = 3
1 -> 2
2 -> 2

*/

const maxArea = (heights) => {
  let result = 0;

  for (let i = 0; i < heights.length; i++) {
    let startBar = heights[i];
    for (let j = heights.length -1; j > i; j--) {
      let endBar = heights[j];
      let distance = j - i;
      let minHeight = Math.min(startBar, endBar);
      let area = findArea(distance, minHeight);
      if (area > result) {
        result = area;
      }
    } 
  }
}

const findArea = (distance, minHeight) => {
  return distance * minHeight;
}