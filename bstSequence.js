// given a bst generated from an array
// find all possible arrays that could have generated it

const getSequences = (node) => {
  if (node === null) return [];;

  let weaved = weave(getSequences(node.left), getSequences(node.right), node.val);

  return weaved;
};

const weave = (leftArrays, rightArrays, prefix) => {

  let weaved = [];

  for (let i = 0; i < leftArrays.length; i++) {
    let leftArray = leftArrays[i];

    for (let j = 0; j < rightArrays.length; j++) {
      let rightArray = rightArrays[j];

      let result = [prefix];

      if (leftArray.length === 0 || rightArray.length === 0) {
        // add rest of left
        // add rest of right
        result = result.concat(leftArray);
        result = result.concat(rightArray);
      }
    }
  }
};

//steps
// getSequences recursively combines the arrays of the left and right nodes
// base case is the bottom where left and right are null, so you just add a prefix to nothing

// weave together step
// take the current node as a prefix
// for each array on left
  // for ach array on right
    // get all the combinations while preserving order and append the prefix to each, then return them;



class LinkedList {
  constructor() {
    this.head = null;
    this.tail = null;
  }

  removeHead() {
    let oldHead = this.head;

    if (this.head === null) return null;

    if (this.head === this.tail) {
      this.tail = null;
    }
    this.head = this.head.next;
    return oldHead;
  }

  addToTail(val) {
    let newTail = new Node(val);

    if (this.tail) {
      this.tail.next = newTail;
    } else {
      this.head = newTail;
    }
    this.tail = newTail;
  }
}

class Node {
  constructor(val) {
    this.val = val;
    this.next = null;
  }
}


class BinarySearchTree {
  constructor(val) {
    this.val = val;
    this.left = null;
    this.right = null;
  }
}

const root = new BinarySearchTree(5);
root.left = new BinarySearchTree(3);
root.right = new BinarySearchTree(8);
root.left.left = new BinarySearchTree(2);
root.left.right = new BinarySearchTree(4);
root.right.left = new BinarySearchTree(7);
root.right.right = new BinarySearchTree(9);

