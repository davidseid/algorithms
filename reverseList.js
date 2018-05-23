const reverseList = (head) => {

  let prev = null;
  let curr = head;

  while (curr !== null) {
    let tempNext = curr.next;
    curr.next = prev;
    prev = curr;
    curr = tempNext;
  }
  return prev;
}