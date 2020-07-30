/*
This should be close to working! Keep at it
 */

class LRUNode {
    constructor(key, value) {
        this.key = key;
        this.value = value;
        this.prev = null;
        this.next = null;
    }
}

class LRUCache {
    constructor(capacity) {
        this.size = 0;
        this.capacity = capacity;
        this.map = {};
        this.head = null;
        this.tail = null;
    }

    keyInMap(key) {
        if (this.map[key]) return true;
        return false;
    }

    getNodeFromMap(key) {
        return this.map[key];
    }

    createNode(key, value) {
        return new LRUNode(key, value);
    }

    addNodeToCache(node) {
        this.size++;
        if (this.head === null) {
            this.head = node;
        }
        this.addNodeToTail(node);
        this.map[node.key] = node;
    }

    put(key, value) {
        if (this.keyInMap(key)) {
            this.map[key].value = value;
            const node = this.getNodeFromMap(key);
            this.moveToTail(node);
        } else {
            if (this.size >= this.capacity) {
                this.removeHead();
            }
            const node = this.createNode(key, value);
            this.addNodeToCache(node);
        }
    }

    addNodeToTail(node) {
        if (this.tail !== null) {
            this.tail.next = node;
            node.prev = this.tail;
        }
        this.tail = node;
        this.tail.next = null;
    }

    get(key) {
        if (this.keyInMap(key)) {
            const node = this.map[key];
            this.moveToTail(node);
            console.log(node.value);
            return node.value;
        } else {
            console.log(-1);
            return -1;
        }
    }

    moveToTail(node) {
        if (node.next !== null) {
            if (node.prev !== null) {
                node.prev.next = node.next;
            } else {
                this.head = node.next;
            }
            node.next.prev = node.prev;
            this.addNodeToTail(node);
        }
    }

    removeHead() {
        delete this.map[this.head.key];
        this.size--;
        this.head = this.head.next;
        if (this.head !== null) {
            this.head.prev = null;
        }
    }
}

let cache = new LRUCache(10); 
cache.put(10, 13);
cache.put(3, 17);
cache.put(6, 11);
cache.put(10, 5);
cache.put(9, 10);
cache.get(13);
cache.put(2, 19);
cache.get(2);
cache.get(3);
cache.put(5, 25);
cache.get(8);
cache.put(9, 22);
cache.put(5, 5);
cache.put(1, 30);
cache.get(11);
cache.put(9, 12);
cache.get(7);
cache.get(5);
cache.get(8);
cache.get(9);
cache.put(4, 30);
cache.put(9, 3);
cache.get(9);
cache.get(10);
cache.get(10);
cache.put(6, 14);
cache.put(3, 1);
cache.get(3);
cache.put(10, 11);
cache.get(8);
cache.put(2, 14);
cache.get(1);
cache.get(5);
cache.get(4);
cache.put(11, 4);
cache.put(12, 24);
cache.put(5, 18);
cache.get(13);
cache.put(7, 23);
cache.get(8);
cache.get(12);
cache.put(3, 27);
cache.put(2, 12);
cache.get(5);
cache.put(2, 9);
cache.put(13, 4);
cache.put(8, 18);
cache.put(1, 7);
cache.get(6);
