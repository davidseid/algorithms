/*
  "(()"
  ")()())"
  "(()(())"

  ")"

*/

const isValid = (str) => {
  let counter = 0;

  for (let i = 0; i < str.length; i++) {
    if (str[i] === '(') {
      counter++;
    } else {
      counter--;
    }
    if (counter < 0) {
      break;
    }
  }
  if (counter === 0) return true;
  return false;
}

const longestValidParentheses = (s) => {
  
  let start = 0;
  let pointer = 0;
  let validityCounter = 0;
  let longest = 0;

  while (pointer < s.length) {
    start = pointer;

    if (s[pointer] === ')') {
      pointer++;
      continue;
    }

    while (validityCounter >= 0 && s[pointer] !== undefined) {
      if (s[pointer] === '(') {
        validityCounter++;
      } else if (s[pointer] === ')') {
        validityCounter--;
      }

      if (validityCounter === 0) {
        let length = (pointer - start) + 1;
        if (length > longest) longest = length;
      }
      pointer++;
    }

    if (s[pointer] === undefined && validityCounter > 0) {
      
      // make substring
      // while not valid and length > 0
      // make substring smaller
      // if substring valid
      let substring = s.slice(start, pointer);
      while (!isValid(substring) && substring.length > 0) {
        substring = substring.slice(1);
      }

      if (isValid(substring)) {
        if (substring.length > longest) longest = substring.length;
      }

    }

    validityCounter = 0; 


  }
  return longest;

}


let testInput = ')()())';

console.log(longestValidParentheses(testInput));