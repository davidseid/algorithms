/* Goal: reverse nodes k at a time
  - leave extra nodes as is
  - k is guaranteed to be a positive integer
  
  i: head of a linkedlist, k
  o: linked list
  c: O(1) memory
  e: k = 1, k = length of list, k is or is not factor of list length

  Diagram:

  if k = 2

  1 -> 2 -> 3 -> 4 -> 5

  store 1 2 3

  1 -> 2 === 2 -> 1 -> ?
    3 -> 4 === 4 -> 3 -> ?
      5 -> null === 5 -> null
  nodes a, b
  if one of the nodes is null...
  return the head
  otherwise, switch the head and return it 
    and recursively point the tail the result of the next call
  // return the head

  if k = 3

  1, 2, 3

  1 -> 2 -> 3 === 3 -> 2 -> 1 -> ?
    4 -> 5 -> null === 4 -> 5 -> null

  nodes a b c (get d)
  if one of them is null, return head
  otherwise make the last one the head, point to the second, point to the first

  make that point to calling it on d
  return head

*/

const reverseKGroup = (head, k) => {
  let count = k;
  let nodes = [];
  let node = head;
  while (count > 0) {
    if (node === null || node.next === null) return head;
    nodes.push(node);
    node = node.next;
    count--;
  }

  let a = head;
  let b = head.next;
  let c = head.next.next;

  head = b;
  b.next = a;
  a.next = reverseKGroup(c, k);

  return head;
}

class ListNode {
  constructor(val) {
    this.val = val;
    this.next = null;
  }
}

let myList = new ListNode(1);
myList.next = new ListNode(2);
myList.next.next = new ListNode(3);
myList.next.next.next = new ListNode(4);
myList.next.next.next.next = new ListNode(5);

console.log(reverseKGroup(myList, 2));