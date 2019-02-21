


const testString = "Tact Coa";

const isPalPerm = str => {

  // remove whitespace and capitals
  // create bit vector
  // 0 represents even
  // 1 represents odd

  // update bit vector each time we see a count, using the ascii code
  // use bitmath to determine if there is 1 or more than one bit that is 1

  let modStr = str.toLowerCase().replace(' ', '');

  let vector = 0;

  for (let i = 0; i < modStr.length; i++) {
    console.log(('a').charCodeAt(0));
    let code = modStr.charCodeAt(i) - 'a';

    let mask = 1 << code;

    if (!(vector & mask)) {
      vector |= mask;
    } else {
      vector &= ~mask;
    }
  }

  console.log(vector);

  if (vector === 0) {
    console.log('no odds');
    return true;
  }

  if (vector & (vector - 1) === 0) {
    console.log('one odd');
    return true;
  }

  console.log('multiple odds');
  return false;
};



isPalPerm(testString);