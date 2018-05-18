const numMax = (a, b) => {
  let diff = a - b;
  console.log(diff);

  let signBit = (diff >> 31) & 1;
  console.log(signBit);
  
  let k = !signBit;

  return k*a + !k*b;

}

console.log(numMax(45, 100));

/*
// no if else or comparisons

// plan:
// get the diff 
// make k the opposite of the sign bit
// return k(a) + !k(b)


*/
