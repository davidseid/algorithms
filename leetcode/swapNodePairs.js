const swapPairs = (head) => {
  if (head.next === null) return head;
  if (head.next.next === null) {
    let newHead = head.next;
    newHead.next = head;
    head.next = null;
    return newHead;
  }
  if (head.next.next.next === null) {
    let newHead = head.next;
    let temp = head.next.next;
    newHead.next = head;
    head.next = temp;
    return newHead;
  }


}

class List {
  constructor(val) {
    this.val = val;
    this.next = null;
  }
}

let newHead = new List(1);
newHead.next = new List(2);
newHead.next.next = new List(3);
// newHead.next.next.next = new List(4);

console.log(swapPairs(newHead));