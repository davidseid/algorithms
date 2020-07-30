/*
  inputs: -100.0 < x < 100, n 32 bit int
  output: decimal number 
  constraints: none / handle decimals
  edge cases: negative numbers and positive numberes


  2.000, 10 -> 1024

  2.1


*/

const pow = (x, n) => {
  if (n == 0) return 1;
  let y = pow(x, n / 2);

  if (n % 2 === 0) return y * y;
  
  return y * y * x;
}

const myPow = (x, n) => {
  if (n === 0) return 1;

  if (n > 0) {
    return pow(x, n);
  } else if (n < 0) {
    return 1 / pow(x, n);
  }


}

console.log(myPow(2, 10));