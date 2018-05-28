const recursiveMultiply = (a, b) => {
  
  let half;
  if (a === 1) {
    return b;
  } else {
    s = a >> 1;
    half = recursiveMultiply(s, b);
  }

  return half + half;
 
}

console.log(recursiveMultiply(4, 9));