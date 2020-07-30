const hasCycle = (head) => {
    if (head === null || head.next === null) {
        return false;
    } 

    let slow = head;
    let fast = head.next;

    while (slow !== fast) {
        if (fast === null || fast.next === null) {
            return false;
        }
        slow = slow.next;
        fast = fast.next.next;
    }
    return true;
};

/*
- To determine if a linked list has a cycle
- smoothest way is to send two runners
- 
*/