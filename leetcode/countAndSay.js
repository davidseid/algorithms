const countAndSay = (n) => {
  
  let result = '1';
  for (let j = 2; j <= n; j++) {
    let next = '';
    let count = 1;
    let num = result[0];
    for (let i = 1; i <= result.length; i++) {
      if (num === result[i]) count++;
      if (num !== result[i]) {
        next += count + num;
        num = result[i];
        count = 0;
      }
    }
    result = next;
  }
  return result;
}

console.log(countAndSay(4));