class BinaryTree {
  constructor(val) {
    this.val = val;
    this.left = null;
    this.right = null;
  }
}

let myTree = new BinaryTree('1');
myTree.left = new BinaryTree('2');
myTree.right = new BinaryTree('3');
myTree.left.left = new BinaryTree('4');
myTree.left.right = new BinaryTree('5');
myTree.right.left = new BinaryTree('6');
myTree.right.right = new BinaryTree('7');

const treePrinter = (root, rights = 0, path = []) => {
  
  for (let j = 0; j < rights; j++) {
    root.val = ' ' + root.val;
  }
  path.push(root.val);
  
  if (!root.left && !root.right) {
    path.forEach((val) => console.log(val));
    console.log('\n');
  } else {
    if (root.left) {
      path.forEach((val, i) => path[i] = ' ' + path[i]);
      treePrinter(root.left, rights, path);
      path.forEach((val, i) => path[i] = path[i].substring(1));
    }
    if (root.right) {
      rights++;
      treePrinter(root.right, rights, path);
      rights--;
    }
  }
  path.pop();
}

console.log(treePrinter(myTree));