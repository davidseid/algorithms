const countAndSay = (n) => {
  // initiate the first term n === 1 return 1; 
  // for each in n, update the result, at the end return the value;

  let result = "1";
  if (n === 1) return result;
  let nextResult = ''; 
  
  for (let i = 1; i <= n; i++) {
    nextResult = getNextIteration(result);
    result = nextResult;
  }
  return result;

}

const getNextIteration = (lastIteration) => {
  let num = lastIteration[0];
  let count = 0;
  let nextIteration = '';

  for (let i = 0; i < lastIteration.length; i++) {
    if (lastIteration[i] === num) {
      count++;
    } else {
      nextIteration += count;
      nextIteration += num;
      num = lastIteration[i];
      count = 0;
    }
  }

  return nextIteration;
}

console.log(countAndSay(4));