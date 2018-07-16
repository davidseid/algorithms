

class Tree {
  constructor(val) {
    this.val = val;
    this.children = [];
  }
}

myTree = new Tree(1);
myTree.children.push(new Tree(2), new Tree(3));
myTree.children[0].children.push(new Tree(4));
myTree.children[1].children.push(new Tree(5), new Tree(6));

const depthSearch = (root) => {
  console.log(root.val);
  for (let i = 0; i < root.children.length; i++) {
    let child = root.children[i];
    depthSearch(child);
  }
}


const breadthSearch = (root, queue = []) => {
  queue.push(root);
  
  while (queue.length > 0) {
    let node = queue.shift();
    console.log(node.val);
    for (let i = 0; i < node.children.length; i++) {
      queue.push(node.children[i]);
    }
  }
}

let test = [0, 2, 5, 7, 19, 22, 33];

const binarySearch = (nums, target, start = 0, end = nums.length - 1) => {
  let midIndex = Math.floor(end + start / 2);
  let result = -1;

  if (target === nums[midIndex]) return midIndex;
  if (start === end) return -1;

  console.log(midIndex)
  if (target > nums[midIndex]) {
    result = binarySearch(nums, target, midIndex + 1, end);
  } else if (target < nums[midIndex]) {
    result = binarySearch(nums, target, start, midIndex);
  }

  return result;
}

console.log(binarySearch(test, 19));