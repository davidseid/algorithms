const romanToInt = (romanNum) => {
  let result = 0;

  let converter = {
    'I': 1,
    'V': 5,
    'X': 10,
    'L': 50,
    'C': 100,
    'M': 1000
  }

  for (let i = 0; i < romanNum.length; i++) {
    let currChar = romanNum[i];
    let nextChar = romanNum[i + 1];

    if (converter[nextChar] > converter[currChar]) {
      result -= converter[currChar];
    } else {
      result += converter[currChar];
    }
  }

  return result;
}

console.log(romanToInt('CIV'));