

const urlify = (str, size) => {
  let p = size - 1;
  let i = size - 1;

  let inWord = false;

  while (i >= 0) {
    console.log(i, p);
    console.log(str, str[i], str[p]);
    if (!inWord) {
      if (str[i] !== ' ') {
        inWord = true;
      }
    }

    if (inWord) {
      if (str[i] === ' ') {
        str[p] = '0';
        str[p - 1] = '2';
        str[p - 2] = '%';
        p -= 3;
        console.log(str);
      } else {
        str[p] = str[i];
        p--;
      }
    }
    i--;
  }
};

console.log(urlify('Mr John Smith    ', 13));
