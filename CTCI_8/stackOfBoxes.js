const findTallestStack = (availBoxes, currStack = [], currHeight = 0, maxHeight = 0) => {
  
  if (availBoxes.length === 0) {
      if (currHeight > maxHeight) {
          maxHeight = currHeight;
      }
  }

  for (let i = 0; i < availBoxes.length; i++) {
      let nextStack = currStack.slice();
      let nextHeight = currHeight;
      let chosenBox = availBoxes[i];
      nextStack.push(chosenBox);
      nextHeight += chosenBox.h;
      let minH = chosenBox.h;
      let minW = chosenBox.w;
      let minL = chosenBox.l;

      let remainingBoxes = [];

      for (let j = 0; j < availBoxes.length; j++) {
          let boxToCheck = availBoxes[j];
          if (boxToCheck.l > minL && boxToCheck.w > minW && boxToCheck.h > minH) {
              remainingBoxes.push(boxToCheck);
          } 
      }

      maxHeight = findTallestStack(remainingBoxes.slice(), nextStack, nextHeight, maxHeight);
  }

  return maxHeight;

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