/*
input: two strings representing binary numbers (non empty)
output: string represennting binary number (sum of a and b)
constraints: none
edgecases: none

examples: 
11 1 -> 100

1010, 1011 -> 10101


*/
const makeEqualLength = (a, b) => {
  if (a.length > b.length) {
    while (b.length < a.length) {
      b = '0' + b;
    }
  } else {
    while (a.length < b.length) {
      a = '0' + a;
    }
  }

  return [a, b];
}

const addBinary = (a, b) => {
  let sum = '';
  let carry = 0;

  [a, b] = makeEqualLength(a, b);


  
  for (let i = a.length - 1; i >= 0; i--) {
    let digitA = Number(a[i]);
    let digitB = Number(b[i]);
    let digitSum = digitA + digitB + carry;

    if (digitSum < 2) {
      sum = digitSum + sum;
      carry = 0;
    } else if (digitSum === 2) {
      sum = '0' + sum;
      carry = 1;
    } else if (digitSum === 3) {
      sum = '1' + sum;
    }
  }

  if (carry === 1) {
    sum = '1' + sum;
  }

  return sum;
}

console.log(addBinary('1111', '1111'));