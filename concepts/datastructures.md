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
