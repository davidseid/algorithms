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

// Second Attempt with Carry
const addTwoNumbers = function(l1, l2) {
  // traverse the linkedlists at the same time
  // each time add the sum together and build a new linked list with that sum
  // if there is anything about 9, track the carry (1 or 0, and add it to the next)
  // make sure to handle if one list is less than the other
  
  // make answer linkedList head
  // currentL1
  // currentL2
  // carry = 0
  // while currentL1 or currentL2 have next
    // if both have next
      // sum both values and add to to the answer
      // make a new node for the answer
      // move to the next nodes of l1 and l2
    // if currentl1 has next but not currentl2
      // add currentl1 to answer
      // make a new node for the answer
      // move to next node of current 1
    // if currentl2 has next but not currentl1
      // add currentl2 to answer
      // make new node for ansewr
      //m ove to next node of current 2
  // do one more version of it all for the last node 

  let sumLinkedListHead = new LinkedList(0);
  let sumNode = sumLinkedListHead;
  let nodel1 = l1;
  let nodel2 = l2;
  let carry = 0;

  while (nodel1.next !== null || nodel2.next !== null) {
    let valuel1 = nodel1.val;
    let valuel2 = nodel2.val;
    let rawSum = valuel1 + valuel2 + carry;
    carry = 0; 
    let answerVal;
    if (rawSum < 10) {
      answerVal = rawSum;
    } else {
      answerVal = rawSum % 10;
      carry = 1; 
    }
    sumNode.val = answerVal;
    addNode(sumNode, null);
    sumNode = sumNode.next;

    if (nodel1.next !== null && nodel2.next !== null) {
      nodel1 = nodel1.next;
      nodel2 = nodel2.next;
    } else if (nodel1.next !== null) {
      nodel1 = nodel1.next;
      nodel2.val = 0;
    } else if (nodel2.next !== null) {
      nodel2 = nodel2.next;
      nodel1.val = 0;
    }
  }
    
  let lastRaw = nodel1.val + nodel2.val + carry;
  if (lastRaw < 10) {
    sumNode.val = lastRaw;
  } else {
    sumNode.val = lastRaw % 10;
    addNode(sumNode, 1);
  }
  
  console.log(sumLinkedListHead);
  let answerArray = [];
  answerArray.push(sumLinkedListHead.val);
  sumNode = sumLinkedListHead;
  while (sumNode.next !== null) {
    answerArray.push(sumNode.next.val);
    sumNode = sumNode.next;
  }
  return answerArray;
}

var LinkedList = function(val) {
  this.val = val;
  this.next = null;
}

var addNode = (linkedList, val) => {
  if (linkedList.next === null) {
    linkedList.next = new LinkedList(val);
  } else {
    addNode(linkedList.next, val);
  }
}