


const deleteDuplicates = (head) => {

  let node = head;
  while (node.next !== null) {
    if (node.next.val === node.val) {
      node.next = node.next.next;
    }
    node = node.next;
  }
  return head;
}

class ListNode {
  constructor(val) {
    this.val = val;
    this.next = null;
  }
}

let myHead = new ListNode(1);
myHead.next = new ListNode(1);
myHead.next.next = new ListNode(2);

console.log(myHead);
console.log(deleteDuplicates(myHead));