/*
  input: non empty array of digits non negative
  output: array of digits representing number plus one
  constriants: none 
  edge cases: number could be zero


  [1, 2, 3]
  [1, 2, 4]

  [4, 3, 2, 1]
  [4, 3, 2, 2]

  [9, 9]
  [1, 0, 0]



  Explanation:
  starting from the back, if the number <= 8
    add one to that number
  if number === 9
    make it zero and repeat the process on the next number

  
*/


const plusOne = (digits) => {
  
  let i = digits.length - 1;

  while (i >= 0) {
    if (digits[i] < 9) {
      digits[i]++;
      return digits;
    } else if (digits[i] === 9) {
      digits[i] = 0;
    }
    i--;
  }
  digits.unshift(1);
  return digits;
}

let testInput = [9, 9];
console.log(plusOne(testInput));