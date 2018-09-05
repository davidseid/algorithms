/*
    Design LRU cache data structure
    - get and put methods

    get(key) -- gets the value, always positive, of the key if the key exists in the cache, otherwise -1
    put(key, value) -- sets the value if the key is not already present, when the cache reaches capacity, should invalidate the lru item before inserting new item

    Followup: both in O(1) time complexity
 */




class LRUCache {
    constructor(capacity) {
        this.capacity = capacity;
        this.dictionary = [];
        this.order = [];
        this.size = 0;
    }

    get(key) {

        let value = this.dictionary[key];



        if (value) {

            let indexOfItemInCache = this.order.indexOf(key);
            this.order.splice(indexOfItemInCache, 1);
            this.order.push(key);

            return value;
        } else {
            return -1;
        }

    }

    put(key, value) {

        let indexOfItemInCache = this.order.indexOf(key);

        if (indexOfItemInCache !== -1) {

            this.order.splice(indexOfItemInCache, 1);

        } else {

            if (this.size < this.capacity) {
                this.size++;

            } else {
                let keyToRemove = this.order[0];
                this.dictionary[keyToRemove] = -1;
                this.order.shift();
            }
        }

        this.dictionary[key] = value;
        this.order.push(key);

        // console.log(`ON PUT key ${key}, order is ${this.order}`);
        // console.log(`dict is ${this.dictionary}`);
    }
}

let output = [null,null,null,null,null,null,-1,null,19,17,null,-1,null,null,null,-1,null,-1,5,-1,12,null,null,3,5,5,null,null,1,null,-1,null,30,5,30,null,null,null,-1,null,-1,24,null,null,18,null,null,null,null,-1,null,null,18,null,null,-1,null,null,null,null,null,18,null,null,-1,null,4,29,30,null,12,-1,null,null,null,null,29,null,null,null,null,17,-1,18,null,null,null,-1,null,null,null,20,null,null,null,-1,18,18,null,null,null,null,20,null,null,null,null,null,null,null];
let expected = [null,null,null,null,null,null,-1,null,19,17,null,-1,null,null,null,-1,null,-1,5,-1,12,null,null,3,5,5,null,null,1,null,-1,null,30,5,30,null,null,null,-1,null,-1,24,null,null,18,null,null,null,null,-1,null,null,18,null,null,-1,null,null,null,null,null,18,null,null,-1,null,4,29,30,null,12,-1,null,null,null,null,29,null,null,null,null,17,22,18,null,null,null,-1,null,null,null,20,null,null,null,-1,18,18,null,null,null,null,20,null,null,null,null,null,null,null];

for (let i = 0; i < output.length; i++) {
    if (output[i] !== expected[i]) {
        console.log(i, output[i], expected[i]);
    }
}