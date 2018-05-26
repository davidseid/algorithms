
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
  
  return head;
}


