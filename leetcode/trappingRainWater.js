
/*
  0 < 1 so pool starts and we know we get 1
  2 -> pool ends
  1 -> less than 2 so pool starts
*/

const trap = (height) => {

  let water = 0;

  let start;
  let end;

  for (let i = 0; i < height.length - 1; i++) {
    let currHeight = height[i];
    let nextHeight = height[i + 1];

    if (!start && nextHeight < currHeight) {
      start = currHeight;
      end = nextHeight;
    } 

    if (start && currHeight >= start) {
      console.log('end pool');
      start = 
    } else if (start && currHeight >= end) {
      end = currHeight;
    }
  }

  
};

/*
  102 21013 3212
*/

let testInput = [0,1,0,2,1,0,1,3,2,1,2,1];

console.log(trap(testInput));