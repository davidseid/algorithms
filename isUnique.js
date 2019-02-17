
const isUnique = (str) => {
  let bitVector = 0;

  for (let i = 0; i < str.length; i++) {
    let code = str.charCodeAt(i);
    console.log(code);

    if (bitVector & (1 << code)) {
      return false;
    } else {
      bitVector |= (1 << code);
    }
  }

  return true;
};


console.log(isUnique('halol'));