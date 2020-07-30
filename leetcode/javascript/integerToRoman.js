const intToRoman = (num) => {
  let result = '';

  while (num >= 1000) {
    num -= 1000;
    result += 'M';
  }

  if (num >= 900) {
    num -= 900;
    result += 'CM';
  } else if (num >= 500) {
    num -= 500;
    result += 'D';
  } else if (num >= 400) {
    num -= 400;
    result += 'CD';
  }

  while (num >= 100) {
    num -= 100;
    result += 'C';
  }

  if (num >= 90) {
    num -= 90;
    result += 'XC';
  } else if (num >= 50) {
    num -= 50;
    result += 'L';
  } else if (num >= 40) {
    num -= 40;
    result += 'XL';
  }

  while (num >= 10) {
    num -= 10;
    result += 'X';
  }

  if (num >= 9) {
    num -= 9;
    result += 'IX';
  } else if (num >= 5) {
    num -= 5;
    result += 'V';
  } else if (num >= 4) {
    num -= 4;
    result += 'IV';
  }

  while (num > 0) {
    num -= 1;
    result += 'I';
  }

  return result;
}

console.log(intToRoman(1994));