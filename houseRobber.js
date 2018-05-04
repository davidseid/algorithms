var rob = function(nums) {
  let evens = nums.filter((num) => {
    return num % 2 === 0;
  });
  console.log(evens);
  let odds = nums.filter((num) => {
    return num % 2 !== 0;
  })

  let evenSum = evens.reduce((a, b) => a + b);
  let oddSum = odds.reduce((a, b) => a + b);

  return Math.max(evenSum, oddSum);
};