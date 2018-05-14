# Data Structures Notes

## Big O

Additional Problems Answers:
- VI.1: O(b) 
- VI.2: O(b)
- VI.3: O(1)
- VI.4: O(a/b) 
- VI.5: O(log n)
- VI.6: O(log n) !!! O(sqrt n)
- VI.7: O(n)
- VI.8: O(n)
- VI.9: O(n^2)
- VI.10: O(n/10) -> O(n) !!! O(log n) 
- VI.11: O(2^n) -> !!! O(kc^k)
- VI.12: O(n) !!! O(b log b + a log b)

## Technical Questions

Powers of 2 Table

Power of 2
7 -> 128
8 -> 256
10 -> 1024 -> 1k -> 1KB
16 -> 65,536 -> -> 64KB
20 -> 1,048,576 -> 1M -> 1 MB
30 -> 1,073,741,824 -> 1B -> 1 GB
32 -> 4,294,967,296 ->  -> 4 GB
40 -> 1,099,511,627,776 -> 1T -> 1TB

Purpose of this table is to quickly associate integers and bits with approx values and bytes in a machine. 

### Arrays and Strings Ch. 1

- ArrayLists are array-like and offer dynamic resizing
  - typical implementation is that when full, array doubles
  - in JavaScript, arrays are naturally very flexible
  - ASCII is 128 chars, and Unicode is a super set of 2^21 chars. UTF-32 or UTF-8 allow for byte conversions of unicode characters

- Finished Problems 1.1 - 1.3

### Linked Lists Ch. 2

- Single v doule
- Drawback -- no constant time access to particular index
- Benefit -- add and remove items from the beginning in constant time
- Also easy to remove a node without shifting things over 
- Be familiar with "runner" technique, two pointers, tortoise/hare, or weaving
- Recursive problems are common (take O(n) space)

- Finished Problems 2.1 - 2.3

### Stacks and Queues Ch. 3
- STACK
- pop, push, peek, isEmpty
- no constant time access to ith item, but constant time add and removes to the top
- can be implemented with an array or a linkedlist 
- can be useful for backtracking in recursion
- QUEUE
- add to end, remove first, peek at top, isEmpty
- queues can also be implemented with linkedlist, as long as add/remove are from opposite sides
- queues are good for BFS

- Finished Problems 3.1 - 3.3 

### Bit Manipulation
- Twos Complement for negatives 
  - Positive numbers are represented as themselves
  - Negative are the two's complement of its absolute value with 1 sign to indicate negative
  - Example: -3 as a 4 bit number 
    - abv value is 3, 4 bit comparison is 2^3 = 8
    - 8 - 3 = 5, which is 101, meaning the negative is 1101.
    - or invert bits and add 1 11 = 3 011 --> 100 + 1 = 101 -> 1101
- Arithmetic v Logical Shift
  - Arithmetic divides by two, logical shift is more visual
  - Logical = >>> 
  - Arithmetic = >> 
  - Logical shifts right and adds a 0, Arithmetic shifts right and adds the sign 1 or 0 depending
- Get Bit Method
  - given a num and i
  - 1 -> 0001000 -> mask that with AND on the number, that clears all the non 1 numbers, and the result comparison is either 0 or 1
  getBit(num, i) {
    return ((num & (1 << i)) != 0);
  }
  - 10101 = num , i = 4 
  1 << i = 10000
  - 10101 & 10000 = 10000 != 0, therefore true = 1, return true. 
- Set Bit
  - setBit(num, i) {
    return ((num | (1 << i));
  } 
  - num = 10101, i = 4 
  - return 10101 | 10000
  - little unclear, seems like that only creates a 1 in the bit place
- Clear Bit 
  - clearBit(num, i) {
    mask = ~(i << i)
    return mask & num;
  }
- UpdateBit 
  - updateBit(num, i, bitIs1) {
    value = bitIs1 ? 1 : 0;
    mask = ~(1 << 1);
    return num & mask || value << i  
  }
- In summary, getBit works by making a mask with a 1 in the spot, performing an &, and returning whether that number != 0. 
- setBit works by creating a mask and performing an OR to keep all the other bits the same but converting i to 1
- clearBit works by creating an inverse mask and performing an AND
- updateBit works by creating an inverse mask to clear the bit, and then change it to either 1 or 0 with a new mask based on the input

### Object-Oriented Design

Step 1: Handle ambiguity
  - Ask clarifying questions, don't make assumptions
  - Ask WHO is going to use it and HOW they are going to use it
  - Possibly who what when where why and how
  - This helps you develop the API / methods
  - OOD for a coffee maker. 
    - Who - someone who wants coffee
    - How - fill with water, fill with coffee, press button
    - What - coffee, water, maker
    - When - morning every day
    - Why - for energy and as a ritual
    - Where - in the kitchen
  - CoffeeMaker - fill with water, fill with coffee, start brew

Step 2: Define the core objects
  - For a restaurant, table, guest, party, order, meal, employee, server, host

Step 3: Analyze Relationships
  - Which objects are members of others? Do they inherit? Many to many or one to many?

Step 4: Investigate Actions
  - Party goes to restaurant, guest requests table from host, host looks up reservation, etc.

Design Patterns 
  - Singleton Class:
    - Class has only one instance, ensures access to the instance through the app
    - Useful where global object with exactly one instance
    - May be an antipattern because it interferes with unit testing

  - Factory Method:
    - A factory takes a parameter telling it which class to instantiate
    - Uses a more abstract creator class
    - Ex: CardGame, createCardGame(type of Game)

Deck of Cards:
  -  Ambiguities: 
    - Who: gameplayer
    - How: uses a full deck of cards to play blackjack
    - 52 cards, with two jokers
  - Core Objects:
    - Card, Suits, Numbers, Jokers, Players, Dealer, Deck, Game, Hand
  - Relationships:
    - Card object takes suit and number to build a card, joker has no number
    - Players have hands
    - Deck hass cards
    - Game has players, dealer, and decks
  - Actions:
    - Hit, Stay, Deal

### Recursion and Dynamic Programming

Bottom Up Approach
- Start with base case, then two, three, etc.
- build one case off the other

Top Down Approach
- How do we divide up this problem? 

Half and Half Approach
- Divide data set in half
- Binary search, merge sort

Recursive v Iterative
- Recursive algorithms can be space inefficient because of call stack
- All recursive problems can be implemented iteratively, but sometimes much more complex

Dynamic Programming
- Take a recursive algorithm and find the overlapping subproblems (repeated calls);- Cache those results
- Top Down Dynamic Programming is also called memoization, and bottom up is dynamic programming
- nth Fibonacci is a simple example 

Fibonacci Numbers
- fibonacci(n) {
  // 1 1 2 3 5 8 13 21
  // fib(1) = 1
  if n === 1
    return 1
  if n === 0
    return 0;
  else 
    return fib(n -1) + fib(n - 2);
}

- with dynamic programming

fibonnaci(n) {
  let memo = {};
  memo[0] = 0;
  memo[1] = 1;

  if (n === 0) {
    return 0;
  }
  if (n === 1) {
    return 1;
  }

  for (let i = 2; i < n; i++) {
    memo[i] = memo[i - 1] + memo[i - 2];
  }
  return memo[n - 1] + memo[n - 2];
}

- with dp but not storing unnecessary info

fibonacci(n) {
  let a = 0;
  let b = 1;

  for (let i = 2; i < n; i++) {
    let c = a + b;
    let a = b;
    let b = c;
  }
  return a + b;
}

### System Design and Scalability

General Tips:
- Communicate
- Broad first
- Whiteboard draw
- Acknowledge concerns
- Don't assume too much
- State any you make
- Estimate!
- Drive

Design Steps:

1. Scope
  - List features and use cases (6Ws)
2. Reasonable Assumptions
  - Out of date, etc
3. Major Components
  - Diagram
  - Servers, data, client, workers
  - Show flow (disregard scalability first)
4. Identify Key Issues
  - Bottlenecks/challenges
5. Redesign
  - Adjust on whiteboard to accommodate challenges 
  - Be aware of drawbacks

Algorithms that Scale Steps:

1. Ask questions
2. Make Believe, remove constraints first to get a general solution outline
3. Get Real 
  - think about how you would realistically divide the data
4. Solve Problems
  - Poke holes in your solution and try to solve it

Key Concepts

- Horizontal v Vertical Scaling
- Load Balancing 
- Database denormalization 
  - Denormalization avoids joins
  - NoSQL can also help with that
- Database partitioning (sharding)
  - vertical partitioning
    - partition by feature
    - drawback is might need to repartition
  - key-based partitioning
    - expensive to shift if num servers changes
  - directory based
    - maintain a lookup table
    - single point of failure
    - constantly accessing it hurts performance
    - allows for flexible servers
Caching
  - Key value pair 
  - can cache data or html
Asynchronous Processing
  - Slow -> async
  - Speed v correctness tradeoffs
Networking metrics
  - bandwidth 
    - max amount of data in unit of time bits per second or gB/s
  - throughput
    - actual amount in given time
  - latency
    - how long it takes for data transfer
MapReduce
  - program to process large amounts of data
  - map makes a kv pair
  - reduce takes a key and set of values and reduces them in some way, emitting a new key and value to be reduced
  - allows for processing in parallel
Considerations
  - Failures
  - Availability and Reliability
  - Read v write
  - Security
No Perfect System
  - lots of ways to solve a problem

### Sorting and Searching 

#### Bubble Sort

Runtime: O(n^2)
Memory: O(1)
Explanation: Repeatedly loop from beginning to end, swapping swapping pairs that are out of order. The smaller items bubble up to the beginning of the list. 

