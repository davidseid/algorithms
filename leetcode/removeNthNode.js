const removeNthFromEnd = (head, n) => {
  let counter = 0;
  let pointer = head;

  while (pointer !== null) {
    pointer = pointer.next;
    counter++;
  }

  let target = n - counter;

  let newCounter = 0;
  while ()


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
head.next.next.next = new ListNode(5);

console.log(head);