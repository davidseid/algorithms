/*
  inputs: -100.0 < x < 100, n 32 bit int
  output: decimal number 
  constraints: none / handle decimals
  edge cases: negative numbers and positive numberes


  2.000, 10 -> 1024

  2.1


*/

const myPow = (x, n) => {
  let result = x;

  if (n === 0) return x;
  if (n === 1) return x;
  if (n === -1) return (1 / x);

  if (n < 0) {
    n = Math.abs(n);
    for (let i = 1; i < n; i++) {
      result = result * x;
    }
    return 1 / result;
  }

  for (let i = 1; i < n; i++) {
    result = result * x;
  }

  

  return result;
}

console.log(myPow(2.00000, 2147483648));