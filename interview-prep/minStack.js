class MinStack {
    constructor() {
        this.stack = [];
        this.min = null;
    }

    print() {
        console.log(this.stack);
        console.log(this.min);
    }

    push(x) {
        if (this.min === null) {
            this.min = x;
            this.stack.push(x);
        } else if (x >= this.min) {
            this.stack.push(x);
        } else if (x < this.min) {
            this.stack.push(2 * x - this.min);
            this.min = x;
        }
    }

    pop() {
        let topElement = this.stack.pop();
        if (topElement < this.min) {
            this.min = 2 * this.min - topElement;
        }
        if (this.stack.length === 0) {
            this.min = null;
        }
    }

    top() {
        let top = this.stack[this.stack.length - 1];
        if (top < this.min) {
            return this.min;
        } else {
            return top;
        }
    }

    // val = 2 * x - min
    // x = val + min / 2

    // top = -6442450943
    // min = -2147483648
    // original value = -2147483648


    getMin() {
        return this.min;
    }
}

const myStack = new MinStack();

/* Testing */
myStack.print();
myStack.push(2147483646);
myStack.push(2147483646);
myStack.push(2147483647);
console.log(myStack.top());
myStack.pop();
myStack.print();
myStack.pop();
myStack.print();
myStack.pop();
myStack.push(2147483647);
console.log(myStack.top());
myStack.print();
myStack.push(-2147483648);
console.log(myStack.top());
myStack.print();
myStack.pop();
myStack.print();
