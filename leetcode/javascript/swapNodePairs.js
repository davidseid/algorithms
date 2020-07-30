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
  
  let a = head; 
  let b = head.next;
  let c = head.next.next;

  let newHead = b;
  let d;

  while (c !== null) {
    b.next = a;
    d = c.next;
    b = c;
    c = d;
    a.next = c;
    a = b; 
    b = c;
    c = d.next;
  }

  b.next = a;
  a.next = null;

  return newHead;

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
newHead.next.next.next = new List(4);

console.log(swapPairs(newHead));