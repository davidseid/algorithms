const removeNthFromEnd = (head, n) => {
  // go through nodes to count up all of them
  // go through again and stop at the one just before the nth from the end
  // point its next to the one just after the nth
  // return the head

  let count = 0;
  let node = head;
  while (node !== null) {
    node = node.next;
    count++;
  }
  console.log(count);
}

class ListNode {
  constructor(val) {
    this.val = val;
    this.next = null;
  }
}

let head = new ListNode(1);
head.next = new ListNode(2);
head.next.next = new ListNode(3);
head.next.next.next = new ListNode(4);
head.next.next.next.next = new ListNode(5);

removeNthFromEnd(head, 2);