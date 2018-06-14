const removeNthFromEnd = (head, n) => {
  if (head.next === null) return [];
  let lastIndex = 1; 
  let node = head;

  while (node.next !== null) {
    lastIndex++;
    node = node.next;
  }

  node = head;
    
  if (lastIndex === n) {
    return node.next;
  }

  while (lastIndex !== (n + 1)) {
    lastIndex--;
    node = node.next;
  }

  let startNode = node;
  if (n === 1) {
    startNode.next = null;
  } else {
    let nextNode = startNode.next.next;
    startNode.next = nextNode;
  }

  return head;
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

console.log(removeNthFromEnd(head, 2));