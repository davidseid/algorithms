const findTallestStack = (availBoxes, currHeight = 0, maxHeight = 0) => {
  
  if (currHeight > maxHeight) {
      maxHeight = currHeight;
  }
  for (let i = 0; i < availBoxes.length; i++) {
      let remainingBoxes = getRemainingBoxes(availBoxes[i], availBoxes);

      maxHeight = findTallestStack(remainingBoxes, currHeight + availBoxes[i].h, maxHeight);
  }

  return maxHeight;
}

const getRemainingBoxes = (topBox, availableBoxes) => {
  let result = [];
  for (let i = 0; i < availableBoxes.length; i++) {
    if (availableBoxes[i].h > topBox.h && availableBoxes[i].l > topBox.l && availableBoxes[i].w > topBox.w) {
      result.push(availableBoxes[i]);
    }
  }

  return result;
}

class Box {
  constructor(w, h, l) {
    this.w = w;
    this.h = h;
    this.l = l;
  }
}
const boxA = new Box(2, 3, 2);
const boxB = new Box(4, 4, 3);
const boxC = new Box(3, 6, 4);
const boxD = new Box(7, 5, 6);
const testBoxes = [];
testBoxes.push(boxA, boxB, boxC, boxD);

console.log(findTallestStack(testBoxes));