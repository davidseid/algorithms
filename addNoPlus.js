const addNoPlus = (a, b) => {
  // if (b === 0) return a;
  // let sum = a ^ b;
  // let carry = (a & b) << 1;
  // return addNoPlus(sum, carry);


  while (b !== 0) {
    let sum = a ^ b;
    let carry = (a & b) << 1;
    a = sum;
    b = carry;
  }

  return a;
}

console.log(addNoPlus(5, 3));