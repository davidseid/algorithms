/*
  123
  456

      738
    6150
    49200
      56088

*/

const multiply = (num1, num2) => {


  let longerNum;
  let shorterNum;

  if (num1.length >= num2.length) {
    longerNum = num1;
    shorterNum = num2;
  } else {
    longerNum = num2;
    shorterNum = num1;
  }

  let numsToSum = [];
  let multiplier = '';

  for (let i = shorterNum.length - 1; i >= 0; i--) {
    let bottomDigit = shorterNum[i]; // 9
    let bottomValue = bottomDigit.charCodeAt(0) - ('0').charCodeAt(0);
    let num = '';
    let carry = 0;

    for (let j = longerNum.length - 1; j >= 0; j--) {
      let topDigit = longerNum[j]; // 9
      let topValue = topDigit.charCodeAt(0) - ('0').charCodeAt(0);
      
      let tempProduct = topValue * bottomValue; // 81
      tempProduct += carry; // 89
      let onesPlace = tempProduct % 10; // 9
      let tensPlace = Math.floor(tempProduct / 10); // 8

      num = onesPlace.toString() + num ; // 90
      carry = tensPlace; // 8
      
      if (j === 0) {
        if (carry !== 0) {
          num = carry.toString() + num;
        }
      }
    }

    num += multiplier; 
    multiplier += '0';

    num = num.toString();

    numsToSum.push(num);
  }

  let largestLength = numsToSum[numsToSum.length - 1].length;

  for (let i = 0; i < numsToSum.length; i++) {
    while (numsToSum[i].length < largestLength) {
      numsToSum[i] = '0' + numsToSum[i];
    }
  }

  let result = '';
  let finalCarry = 0;

  for (let i = largestLength - 1; i >= 0; i--) {
    let sum = 0;
    for (let j = 0; j < numsToSum.length; j++) {
      sum += numsToSum[j].charCodeAt(i) - ('0').charCodeAt(0);
    }
    sum += finalCarry;

    let ones = sum % 10;
    let tens = Math.floor(sum / 10);

    finalCarry = tens;
    result = ones + result;

    if (i === 0) {
      if (finalCarry !== 0) {
        result = finalCarry + result;
      }
    }
  }
  while (result.length > 1 && result[0] === '0') {
      result = result.slice(1);
  }

  return result;

}


console.log(multiply("999", "999"));

