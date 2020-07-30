class LinkedList {
  constructor() {
    this.head = null;
    this.tail = null;
    this.size = 0;
  }

  removeHead() {
    let oldHead = this.head;
    this.head = this.head.next;
    this.size--;

    if (this.size === 0) this.tail = null;

    return oldHead;
  }

  addToTail(val) {
    let newNode = new LinkedListNode(val);

    if (this.tail) {
      this.tail.next = newNode;
      this.tail = newNode;
    } else {
      this.head = newNode;
      this.tail = newNode;
    }
    this.size++;
  }
}

class LinkedListNode {
  constructor(val) {
    this.val = val;
    this.next = null;
  }
}

class BinaryTree {
  constructor(val) {
    this.val = val;
    this.left = null;
    this.right =null;
  }
}

const generateDepthLists = (root) => {

  let result = {};
  let queue = new LinkedList();
  root.level = 1;

  queue.addToTail(root);

  while (queue.size > 0) {
    let listNode = queue.removeHead();
    let node = listNode.val;

    if (!result[node.level]) {
      result[node.level] = new LinkedList();
    }
    result[node.level].addToTail(node);

    if (node.left) {
      node.left.level = node.level + 1;
      queue.addToTail(node.left);
    }

    if (node.right) {
      node.right.level = node.level + 1;
      queue.addToTail(node.right);
    }
  }

  return result;
};

let testTree = new BinaryTree(5);
testTree.left = new BinaryTree(7);
testTree.right = new BinaryTree(3);
testTree.left.left = new BinaryTree(1);
testTree.left.right = new BinaryTree(4);
testTree.right.left = new BinaryTree(6);
testTree.right.right = new BinaryTree(2);


console.log(generateDepthLists(testTree));