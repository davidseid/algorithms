/*
  input: two numbers (divisor never 0, 32 bit signed ints)
  output: number quotient
  constraints: no mod, division, or multiplication operators
  edge cases: numbers are the same, numbers are positive or negative, first number is 0

  use subtraction, recursion, or bitmath

  ex. 10, 3
  while 3 > 0
  

*/

const divide = (dividend, divisor, count = divisor) => {
  console.log(dividend, divisor, count);
  if (dividend === 0) return 0;

  if (count === 0) return dividend;

  dividend = divide(dividend - divisor, divisor, count - 1);

  return dividend;
}

console.log(divide(10, 3));