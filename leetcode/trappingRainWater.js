
/*
  0 < 1 so pool starts and we know we get 1
  2 -> pool ends
  1 -> less than 2 so pool starts
*/

const findPools = (height) => {
  let pools = [];

  let start;
  let end;
  let level;
  let inPool = false;

  for (let i = 0; i < height.length - 1; i++) {
    let curr = height[i];
    let next = height[i + 1];

    if (!inPool && next < curr) {
      inPool = true;
      start = i;
      level = next;
    } else if (inPool && curr >= level) {
      inPool = false;
      end = i;
      level = curr;
      pools.push([start, end]);
    }
  }

  return pools;


}

const trap = (height) => {

  const pools = findPools(height);
  console.log(pools);

  
};

let testInput = [0,1,0,2,1,0,1,3,2,1,2,1];

console.log(trap(testInput));