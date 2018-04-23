var isMatch = function(s, p) {
  let stringIndex = 0;
  for (let i = 0; i < p.length; i++) {
    let pChar = p[i];
    if (pChar >= 'a' && pChar <= 'z') {
      if (p[i + 1] !== '*') {
        if (s[stringIndex] === pChar) {
          stringIndex++;
        } else {
          return false;
        }
      }
      if (p[i + 1] === '*') {
        while (s[stringIndex] === pChar && s[stringIndex] !== undefined) {
          stringIndex++;
        }
      }
    }
    if (pChar === '.') {
      if (p[i + 1] === '*') {
        return true
      } else {
        stringIndex++;
      }
    }
  }
  if (stringIndex === s.length) {
    return true;
  } else {
    return false;
  }
};