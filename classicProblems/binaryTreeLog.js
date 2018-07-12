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

const treePrinter = (root, path = []) => {
  
  path.push(root.val);
  path.forEach((val) => console.log(val));
  console.log('\n');
  
  if (!root.left && !root.right) {
    path.forEach((val) => console.log(val));
    console.log('\n');
  } else {
    if (root.left) {
      path.forEach((val, i) => path[i] = ' ' + path[i]);
      treePrinter(root.left, path);
      path.forEach((val, i) => path[i] = path[i].substring(1));
    }
    if (root.right) {
      path[path.length - 1] = ' ' + path[path.length - 1];
      treePrinter(root.right, path);
      path[path.length - 1] = path[path.length - 1].substring(1);
    }
  }
  path.pop();
}

console.log(treePrinter(myTree));