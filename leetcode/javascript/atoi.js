var myAtoi = function(str) {

  let hasNumbers = false;
  let isNegative = false;
  let number = '';
  for (let i = 0; i < str.length; i++) {
    let char = str[i];
    if (char === '.' && hasNumbers) {
      break;
    }
    if (char !== ' ' && char !== '-' && char !== '+') {
      if (!hasNumbers && !(char >= '0' && char <= '9')) {
        return 0;
      }
    }
    if (hasNumbers && (char < '0' || char > '9')) {
      break;
    }
    if (char === '+' || char === '-') {
      if (str[i + 1] < '0' || str[i + 1] > '9') {
        return 0;
      }
    }
    if (char === '-' && str[i + 1] >= '0' && str[i + 1] <= '9') {
      isNegative = true;
    }
    if (char >= '0' && char <= '9') {
      number += char;
      hasNumbers = true;
    }
  }
  if (!hasNumbers) {
    return 0;
  }
  number = Number(number);
  if (isNegative) {
    number *= -1;
  }
  if (number > Math.pow(2, 31) - 1) {
    return Math.pow(2, 31) - 1;
  }
  if (number < Math.pow(-2, 31)) {
    return Math.pow(-2, 31);
  }
  return number;
};