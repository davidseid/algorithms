

const climbStairs = (n, cache = {}) => {
  
  cache[1] = 1;
  cache[2] = 2;

  if (cache[n]) {
    return cache[n];
  } else {
    cache[n] = climbStairs(n - 1, cache) + climbStairs(n - 2, cache)
    return cache[n];
  }
  

}

console.log(climbStairs(3));