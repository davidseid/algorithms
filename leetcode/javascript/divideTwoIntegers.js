/*
  input: two numbers (divisor never 0, 32 bit signed ints)
  output: number quotient
  constraints: no mod, division, or multiplication operators
  edge cases: numbers are the same, numbers are positive or negative, first number is 0

  use subtraction, recursion, or bitmath

  ex. 10, 3

  // 1010
  // 0011
  
  
  

*/

const divide = (dividend, divisor) => {
  let negative = false;
  if (dividend === divisor) return 1;
  if (divisor === 1) return dividend;
  if (dividend < 0 ^ divisor < 0) negative = true;
  dividend = Math.abs(dividend);
  divisor = Math.abs(divisor);
  
  let product = 0;
  let result = 0;
  while ((dividend - product) >= divisor) {
    result++;
    product += divisor;
  }

  if (!negative && result > Math.pow(2, 31) - 1) return Math.pow(2, 31) - 1;
  if (negative && result > Math.pow(2, 30)) return 0 - Math.pow(2, 30);

  if (negative) return 0 - result;
  return result;
}

console.log(divide(12, -3));