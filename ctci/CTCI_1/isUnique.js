const hasUniqueChars = (str) => {
  for (let i = 0; i < str.length; i++) {
    for (let j = 0; j < str.length; j++) {
      if (str[i] === str[j] && i !== j) {
        return false;
      }
    }
  }
  return true;

}

