

/*
    - write an lru cache -
    - should have methods to add key/value pair
    - should have method to remove key/value pair
    - should accept a certain size
    - plan is to hold two data structures, a map to hold references to the linked list
    - and a doubly linked list to manage the order
*/

/* 
    - TODO: 
        - write lruCache constructor
        - write linkedListNOde constructor 
*/


class LRUCache {
    constructor(capacity) {
        this.size = 0;
        this.capacity = capacity;
        this.map = {};
        this.head = null;
        this.tail = null;
    }

    get(key) {
        // get the node from the map
        // move it from its current position to the tail
        // return the node
    }

    put(key, value) {
        // create a node
        const node = new Node(value);
        // add it to the tail
        if (!this.tail) {
            this.tail = node;
            this.head = node;
        } else {
            this.tail.next = node;
            this.tail = node;
        }

        if (this.size < this.capacity) {
            this.map[key] = node;
            this.size++;
        } else {
            this.head = this.head.next;
            this.head.prev = null;
            
        }
        // if cache is not at capacity
            // add it to the map
            // increment the size
        // if the cache is at capacity
            // remove the head
            // remove the head node from map
    }
}

class Node {
    constructor(value) {
        this.prev = null;
        this.next = null;
        this.value = value;
    }
}