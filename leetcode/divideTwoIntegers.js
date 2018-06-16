/*
  input: two numbers (divisor never 0, 32 bit signed ints)
  output: number quotient
  constraints: no mod, division, or multiplication operators
  edge cases: numbers are the same, numbers are positive or negative, first number is 0

  use subtraction, recursion, or bitmath

  ex. 10, 3

  product = 0
  result = 0

  result = 1
  product = 3
  r = 2
  p = 6
  r = 3 
  p = 9
  

*/

const divide = (dividend, divisor) => {
  let negative = false;
  if (dividend < 0 || divisor < 0) negative = true;
  dividend = Math.abs(dividend);
  divisor = Math.abs(divisor);
  let result = 0;
  let product = 0;
  while ((dividend - product) > divisor) {
    result++;
    product += divisor;
  }
  if (negative) result = 0 - result;
  return result;
}

console.log(divide(7, -3));