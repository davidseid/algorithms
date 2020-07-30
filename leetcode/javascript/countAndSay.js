const countAndSay = (n) => {
  let result = "1";
  if (n === 1) return result;
  let nextResult = ''; 
  
  for (let i = 1; i < n; i++) {
    nextResult = getNextIteration(result);
    result = nextResult;
  }
  return result;
}

const getNextIteration = (lastIteration) => {
  let num = lastIteration[0];
  let count = 0;
  let nextIteration = '';

  for (let j = 0; j < lastIteration.length; j++) {
    if (lastIteration[j] === num) {
      count++;
    } else {
      nextIteration += count;
      nextIteration += num;
      num = lastIteration[j];
      count = 1;
    }
  }

  if (count > 0) {
      nextIteration += count;
      nextIteration += num;
  }

  return nextIteration;
}