/*
  "(()"
  ")()())"
  "(()(())"

  ")"

*/

const longestValidParentheses = (s) => {
  let longest = 0;
  let subStart = 0;
  
  while (subStart < s.length) {
    let pointer = subStart;
    
    if (s[subStart] === '(') {
      let stackCount = 0;
      let subLength = 0;
      
      while (stackCount >= 0 && pointer < s.length) {

        if (s[pointer] === '(') {
          stackCount++;
        } else if (s[pointer] === ')') {
          stackCount--;
        }
        subLength++;
        pointer++;
      }

      if (pointer === s.length) {
        pointer--;
        while (stackCount !== 0) {
          if (s[pointer] === '(') {
            stackCount--;
          } else if (s[pointer] === ')') {
            stackCount++;
          }
          subLength--;
          pointer--;
        }
      }

      if (subLength > longest) longest = subLength;
    }
    subStart++;
  }

  return longest;

}

// let testInput = "()(()"
let testInput = ")()())()()(";

console.log(longestValidParentheses(testInput));


