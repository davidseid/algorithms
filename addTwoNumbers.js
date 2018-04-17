// You are given two non-empty linked lists representing two non-negative integers. 
// The digits are stored in reverse order and each of their nodes contain a single digit. 
// Add the two numbers and return it as a linked list.
// You may assume the two numbers do not contain any leading zero, except the number 0 itself.

// traverse each linked list and build an array of strings representing the number
// join it and turn it into a number
// add it up
// build out a linked list from the number as strings

/**
 * Definition for singly-linked list.
 * function ListNode(val) {
 *     this.val = val;
 *     this.next = null;
 * }
 */
/**
 * @param {ListNode} l1
 * @param {ListNode} l2
 * @return {ListNode}
 */
var addTwoNumbers = function(l1, l2) {
  
  let num1 = getNumberFromLinkedList(l1);
  let num2 = getNumberFromLinkedList(l2);

  let total = num1 + num2;

  total = total.toString().split('').reverse();
  
  total = total.map((num) => {
    return Number(num);
  })

  return total; 
};

const LinkedList = function(val) {
  this.val = val;
  this.next = null;
}

const addNode = (linkedList, val) => {
  if (linkedList.next === null) {
    linkedList.next = new LinkedList(val);
  } else {
    addNode(linkedList.next, val);
  }
}

const getNumberFromLinkedList = (linkedList) => {
  let numbers = [];
  const recurse = (node) => {
    numbers.push(node.val.toString());
    if (node.next !== null) {
      recurse(node.next);
    }
  }
  recurse(linkedList);
  
  numbers = numbers.reverse().join('');
  return Number(numbers);
}