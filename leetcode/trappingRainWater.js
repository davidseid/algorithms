


const trap = (height) => {

  let totalWater = 0;
  let inPool = false;
  let poolLevel;

  for (let i = 0; i < height.length - 1; i++) {
    let currHeight = height[i];
    
    if (!inPool) {
      let nextHeight = height[i + 1];
      if (nextHeight < currHeight) {
        inPool = true;
        poolLevel = currHeight;
      }
    } else if (inPool) {

      if (currHeight >= poolLevel) {
        inPool = false;
      } else {
        totalWater += (poolLevel - currHeight);
      }
    }
  }

  if (inPool) {
    
  }

  return totalWater;
  
};

let testInput = [0,1,0,2,1,0,1,3,2,1,2,1];

console.log(trap(testInput));