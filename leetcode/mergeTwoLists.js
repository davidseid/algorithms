
const mergeTwoLists = (l1, l2) => {
  let head;
  let p1 = l1;
  let p2 = l2;
  
  if (l1.val <= l2.val) {
    head = l1;
    p1 = l1.next;
  } else {
    head = l2;
    p2 = l2.next; 
  }

  let curr = head; 

  while (p1 !== null && p2 !== null) {
    if (p1.val <= p2.val) {
      curr.next = p1;
      curr = curr.next;
      p1 = p1.next;
    } else {
      curr.next = p2;
      curr = curr.next;
      p2 = p2.next;
    }
  }
  
  if (p1 === null) {
    while (p2 !== null) {
      curr.next = p2;
      curr = curr.next;
      p2 = p2.next;
    }
  } else {
    while (p1 !== null) {
      curr.next = p1;
      curr = curr.next;
      p1 = p1.next;
    }
  }

  return head;
}

class Node {
  constructor(val) {
    this.val = val;
    this.next = null;
  }
}

let test1 = new Node(1);
test1.next = new Node(2);
test1.next.next = new Node(4);

let test2 = new Node(1);
test2.next = new Node(3);
test2.next.next = new Node(4);

console.log(mergeTwoLists(test1, test2));