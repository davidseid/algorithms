const mergeKLists = (lists) => {
  if (lists.length === 0) return [];
  if (lists.length === 1) return lists[0];
  let values = [];

  for (let i = 0; i < lists.length; i++) {
    let node = lists[i];
    while (node !== null) {
      values.push(node.val);
      node = node.next;
    }
  }

  values.sort((a, b) => a - b);

  let head = new List(values[0]);
  let result = head;

  for (let i = 1; i < values.length; i++) {
    result.next = new List(values[i]);
    result = result.next;
  }
  if (head.val === undefined) return [];
  return head;
}

// 1 4 5
// 1 3 4
// 2 6



class List {
  constructor(val) {
    this.val = val;
    this.next = null;
  }
}

let first = new List(1);
first.next = new List(4);
first.next.next = new List(5);

let second = new List(1);
second.next = new List(3);
second.next.next = new List(4);

let third = new List(2);
third.next = new List(6);

console.log(mergeKLists([first, second, third]));